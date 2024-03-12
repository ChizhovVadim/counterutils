package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
)

func train(
	ctx context.Context,
	training, validation []Sample,
	epochs int,
	mainModel *ModelHCE,
	concurrency int,
) error {
	if len(validation) == 0 {
		var validationSize = min(500_000, len(training)/5)
		validation = training[:validationSize]
		training = training[validationSize:]
	}

	log.Println("Train started")
	defer log.Println("Train finished")

	const BatchSize = 16384
	var models = make([]*ModelHCE, concurrency)
	models[0] = mainModel
	for i := 1; i < len(models); i++ {
		models[i] = mainModel.ThreadCopy()
	}

	var rnd = rand.New(rand.NewSource(0))
	for epoch := 1; epoch <= epochs; epoch++ {
		shuffle(rnd, training)
		for i := 0; i+BatchSize <= len(training); i += BatchSize {
			var batch = training[i : i+BatchSize]
			trainBatch(batch, models)
			applyGradients(models)
		}
		log.Printf("Finished Epoch %v\n", epoch)
		validationCost := calcAverageCost(validation, models)
		log.Printf("Current validation cost is: %f\n", validationCost)
	}

	return nil
}

func shuffle(rnd *rand.Rand, training []Sample) {
	rnd.Shuffle(len(training), func(i, j int) {
		training[i], training[j] = training[j], training[i]
	})
}

func trainBatch(samples []Sample, models []*ModelHCE) {
	var index int32 = -1
	var wg = &sync.WaitGroup{}
	for i := range models {
		wg.Add(1)
		go func(m *ModelHCE) {
			defer wg.Done()
			for {
				var i = int(atomic.AddInt32(&index, 1))
				if i >= len(samples) {
					break
				}
				sample := &samples[i]
				m.Train(sample)
			}
		}(models[i])
	}
	wg.Wait()
}

func applyGradients(models []*ModelHCE) {
	for i := 1; i < len(models); i++ {
		models[i].AddGradients(models[0])
	}
	models[0].ApplyGradients()
}

func calcAverageCost(samples []Sample, models []*ModelHCE) float64 {
	var index int32 = -1
	var wg = &sync.WaitGroup{}
	var totalCost float64
	var mu = &sync.Mutex{}
	for i := range models {
		wg.Add(1)
		go func(m *ModelHCE) {
			defer wg.Done()
			var localCost float64
			for {
				var i = int(atomic.AddInt32(&index, 1))
				if i >= len(samples) {
					break
				}
				localCost += m.CalcCost(&samples[i])
			}
			mu.Lock()
			totalCost += localCost
			mu.Unlock()
		}(models[i])
	}
	wg.Wait()
	averageCost := totalCost / float64(len(samples))
	return averageCost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
