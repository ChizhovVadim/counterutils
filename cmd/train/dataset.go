package main

import (
	"context"
	"log"
	"sync"

	"github.com/ChizhovVadim/CounterGo/pkg/common"
	"github.com/ChizhovVadim/counterutils/internal/dataset"
	"github.com/ChizhovVadim/counterutils/internal/pgn"
	"golang.org/x/sync/errgroup"
)

func loadDataset(
	ctx context.Context,
	featureProviderBuilder func() IFeatureProvider,
	sigmoidScale float64,
	gamesFolder string,
	searchRatio float64,
	mirrorPos bool,
	maxSize int,
	concurrency int,
) ([]Sample, error) {
	var result []Sample
	var games = make(chan pgn.GameRaw, 16)
	var samples = make(chan []Sample, 128)
	var datasetReady = make(chan struct{})
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		defer close(games)
		return dataset.LoadGames(ctx, gamesFolder, games, datasetReady)
	})
	var wg = &sync.WaitGroup{}
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		g.Go(func() error {
			defer wg.Done()
			return extractFeaturesNN(ctx, sigmoidScale, searchRatio, featureProviderBuilder(), mirrorPos, games, samples)
		})
	}
	g.Go(func() error {
		wg.Wait()
		close(samples)
		return nil
	})
	g.Go(func() error {
		var err error
		result, err = collectSamplesNN(ctx, samples, false, maxSize, datasetReady)
		return err
	})
	return result, g.Wait()
}

func extractFeaturesNN(
	ctx context.Context,
	sigmoidScale float64,
	searchRatio float64,
	featureExtractor IFeatureProvider,
	mirrorPos bool,
	games <-chan pgn.GameRaw,
	samples chan<- []Sample,
) error {
	for gameRaw := range games {
		var chunk []Sample
		var err = dataset.AnalyzeGame(sigmoidScale, searchRatio, gameRaw, func(di dataset.DatasetItem2) error {
			var features = featureExtractor.ComputeFeatures(di.Pos)
			chunk = append(chunk, Sample{
				FeatureSet: features,
				Target:     float32(di.Target),
			})

			if mirrorPos {
				var mirror = common.MirrorPosition(di.Pos)
				var features = featureExtractor.ComputeFeatures(&mirror)
				chunk = append(chunk, Sample{
					FeatureSet: features,
					Target:     float32(1 - di.Target),
				})
			}

			return nil
		})
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case samples <- chunk:
		}
	}
	return nil
}

func collectSamplesNN(
	ctx context.Context,
	samples <-chan []Sample,
	mergeRepeats bool,
	maxSize int,
	datasetReady chan<- struct{},
) ([]Sample, error) {
	var result []Sample

	for sample := range samples {
		result = append(result, sample...)
		if maxSize != 0 && len(result) >= maxSize {
			if datasetReady != nil {
				log.Println("skip rest positions")
				close(datasetReady)
				datasetReady = nil
			}
		}
	}

	return result, nil
}
