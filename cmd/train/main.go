package main

import (
	"context"
	"log"
	"math/rand"
	"runtime"

	"github.com/ChizhovVadim/CounterGo/pkg/common"
	"github.com/ChizhovVadim/counterutils/internal/domain"
	"github.com/ChizhovVadim/counterutils/internal/tuner"
)

type IFeatureProvider interface {
	ComputeFeatures(pos *common.Position) FeatureSet
	Size() int
}

type FeatureSet struct {
	Input []domain.FeatureInfo
}

type Sample struct {
	FeatureSet
	Target float32
}

func main() {
	var err = trainHandler()
	if err != nil {
		log.Println(err)
	}
}

func trainHandler() error {
	var (
		gamesFolder = mapPath("~/chess/Dataset2023/stockfish")
		//validationPath = mapPath("~/chess/tuner/quiet-labeled.epd")
		netFolderPath = mapPath("~/chess/net")
		startingNet   = "" //mapPath("~/chess/n-30-5268.nn")

		searchRatio      = 0.75
		mirrorPos        = true
		maxDatasetSize   = 3_000_000
		sigmoidScale     = calcSigmoidScale()
		featureExtractor = func() IFeatureProvider { return &Feature768Provider{} }
		epochs           = 15
		concurrency      = runtime.NumCPU()
	)

	var ctx = context.Background()

	var training, validation []Sample
	var err error

	training, err = loadDataset(ctx, featureExtractor, sigmoidScale,
		gamesFolder, searchRatio, mirrorPos, maxDatasetSize, concurrency)
	if err != nil {
		return err
	}
	log.Println("Loaded dataset",
		"size", len(training))

	var rnd = rand.New(rand.NewSource(0))
	var model = NewModelNN(rnd)
	if startingNet != "" {
		log.Println("Load net", "path", startingNet)
		err = model.Load(startingNet)
		if err != nil {
			return err
		}
	}

	return train(ctx, training, validation, epochs, model, concurrency, netFolderPath)
}

func calcSigmoidScale() float64 {
	//https://www.chessprogramming.org/Pawn_Advantage,_Win_Percentage,_and_Elo
	const (
		PawnValue                   = 100
		PawnAdvantageWinProbability = 2.0 / 3
	)
	return tuner.ReverseSigmoid(PawnAdvantageWinProbability) / PawnValue
}
