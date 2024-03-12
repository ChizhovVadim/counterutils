package main

import (
	"fmt"
	"math"

	"github.com/ChizhovVadim/counterutils/internal/tuner"
)

// Начальные веса не задаем (линейная функция)
// смещения не используем (симметричная функция с точки зрения белых-черных)
type ModelHCE struct {
	activationFn tuner.IActivationFn
	weights      tuner.Matrix
	wGradients   tuner.Gradients
	cost         tuner.IModelCost
}

func NewModelHCE(
	inputSize int,
) *ModelHCE {
	return &ModelHCE{
		activationFn: &tuner.SigmoidActivation{},
		weights:      tuner.NewMatrix(2, inputSize),
		wGradients:   tuner.NewGradients(2, inputSize),
		cost:         &tuner.MSECost{},
	}
}

func (m *ModelHCE) ApplyGradients() {
	m.wGradients.Apply(&m.weights)
}

func (m *ModelHCE) CalcCost(sample *Sample) float64 {
	var cost float64
	m.work(sample, false, &cost)
	return cost
}

func (m *ModelHCE) Train(sample *Sample) {
	var cost float64
	m.work(sample, true, &cost)
}

func (m *ModelHCE) work(sample *Sample, train bool, cost *float64) {
	const (
		Opening = 0
		Endgame = 1
	)
	var mg, eg float64
	for _, input := range sample.Features {
		var inputIndex = int(input.Index)
		var inputValue = float64(input.Value)
		mg += m.weights.Get(Opening, inputIndex) * inputValue
		eg += m.weights.Get(Endgame, inputIndex) * inputValue
	}
	var phase = float64(sample.MgPhase)
	var mix = phase*mg + (1-phase)*eg
	var strongSideScale float64
	if mix > 0 {
		strongSideScale = float64(sample.WhiteStrongScale)
	} else {
		strongSideScale = float64(sample.BlackStrongScale)
	}
	var x = mix * strongSideScale
	var predicted = m.activationFn.Sigma(x)
	if !train {
		*cost = m.cost.Cost(predicted, float64(sample.Target))
		return
	}
	// back propagation
	var outputGradient = m.cost.CostPrime(predicted, float64(sample.Target)) *
		m.activationFn.SigmaPrime(x) *
		strongSideScale
	for _, input := range sample.Features {
		var inputIndex = int(input.Index)
		var inputValue = float64(input.Value)
		m.wGradients.Add(Opening, inputIndex, inputValue*phase*outputGradient)
		m.wGradients.Add(Endgame, inputIndex, inputValue*(1-phase)*outputGradient)
	}
}

func (m *ModelHCE) Print() {
	const (
		ScaleEval = 100
	)
	var evalSale = ScaleEval / calcSigmoidScale()

	var weights = m.weights.Data
	var wInt = make([]int, len(weights))
	for i := range wInt {
		wInt[i] = int(math.Round(evalSale * weights[i]))
	}
	fmt.Printf("var w = %#v\n", wInt)
}

func (m *ModelHCE) ThreadCopy() *ModelHCE {
	return &ModelHCE{
		activationFn: m.activationFn,
		weights:      m.weights,
		wGradients:   tuner.NewGradients(m.wGradients.Rows, m.weights.Cols),
		cost:         m.cost,
	}
}

func (m *ModelHCE) AddGradients(mainModel *ModelHCE) {
	if m == mainModel {
		return
	}
	m.wGradients.AddTo(&mainModel.wGradients)
}
