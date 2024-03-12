package main

import (
	"context"
	"log"
	"runtime"

	"github.com/ChizhovVadim/CounterGo/pkg/common"
	"github.com/ChizhovVadim/counterutils/internal/domain"
	counter "github.com/ChizhovVadim/counterutils/internal/eval/counter"
	"github.com/ChizhovVadim/counterutils/internal/tuner"
)

type IFeatureProvider interface {
	ComputeFeatures(pos *common.Position) domain.TuneEntry
	Size() int
}

func featureExtractor() IFeatureProvider {
	return counter.NewEvaluationService()
}

type Sample struct {
	domain.TuneEntry
	Target float32
}

func main() {
	var err = tunerHandler()
	if err != nil {
		log.Println(err)
	}
}

func tunerHandler() error {
	var (
		gamesFolder = mapPath("~/chess/Dataset2023/stockfish")
		//validationPath = mapPath("~/chess/tuner/quiet-labeled.epd")
		searchRatio = 0.5
		mirrorPos   = false
		//mergeRepeats   = false
		maxDatasetSize = 3_000_000
		sigmoidScale   = calcSigmoidScale()
		epochs         = 10
		concurrency    = runtime.NumCPU()
	)

	var ctx = context.Background()
	var inputSize = featureExtractor().Size()
	log.Println("input size", inputSize)

	var training, validation []Sample
	var err error

	training, err = loadDataset(ctx, featureExtractor, sigmoidScale, gamesFolder, searchRatio, mirrorPos, maxDatasetSize, concurrency)
	if err != nil {
		return err
	}
	log.Println("Loaded dataset",
		"size", len(training))

	var model = NewModelHCE(inputSize)
	err = train(ctx, training, validation, epochs, model, concurrency)
	if err != nil {
		return err
	}
	model.Print()
	return err
}

func calcSigmoidScale() float64 {
	//https://www.chessprogramming.org/Pawn_Advantage,_Win_Percentage,_and_Elo
	const (
		PawnValue                   = 100
		PawnAdvantageWinProbability = 2.0 / 3
	)
	return tuner.ReverseSigmoid(PawnAdvantageWinProbability) / PawnValue
}
