package main

import (
	"math/rand"

	"github.com/ChizhovVadim/counterutils/internal/domain"
	"github.com/ChizhovVadim/counterutils/internal/tuner"
)

type Neuron struct {
	Activation float64
	Error      float64
	Prime      float64
}

type Layer struct {
	activationFn tuner.IActivationFn
	outputs      []Neuron
	weights      tuner.Matrix
	biases       tuner.Matrix
	wGradients   tuner.Gradients
	bGradients   tuner.Gradients
}

func (l *Layer) ThreadCopy() *Layer {
	return &Layer{
		activationFn: l.activationFn,
		outputs:      make([]Neuron, len(l.outputs)),
		weights:      l.weights,
		biases:       l.biases,
		wGradients:   tuner.NewGradients(l.wGradients.Rows, l.wGradients.Cols),
		bGradients:   tuner.NewGradients(l.bGradients.Rows, l.bGradients.Cols),
	}
}

func NewLayer(
	inputSize int,
	outputs []Neuron,
	activationFn tuner.IActivationFn,
) *Layer {
	var outputSize = len(outputs)
	return &Layer{
		outputs:      outputs,
		activationFn: activationFn,
		weights:      tuner.NewMatrix(outputSize, inputSize),
		biases:       tuner.NewMatrix(outputSize, 1),
		wGradients:   tuner.NewGradients(outputSize, inputSize),
		bGradients:   tuner.NewGradients(outputSize, 1),
	}
}

func (layer *Layer) InitWeightsSigmoid(rnd *rand.Rand) *Layer {
	var outputSize = layer.weights.Rows
	var inputSize = layer.weights.Cols
	var variance = 2.0 / float64(inputSize+outputSize)
	tuner.InitUniform(rnd, layer.weights.Data, variance)
	return layer
}

func (layer *Layer) InitWeightsReLU(rnd *rand.Rand) *Layer {
	var inputSize = layer.weights.Cols
	var variance = 2.0 / float64(inputSize)
	tuner.InitUniform(rnd, layer.weights.Data, variance)
	return layer
}

func (layer *Layer) InitWeightsCount(rnd *rand.Rand, count float64) *Layer {
	tuner.InitUniform(rnd, layer.weights.Data, 1.0/count)
	return layer
}

func (layer *Layer) Forward(input1 []Neuron, input2 []domain.FeatureInfo) {
	for outputIndex := range layer.outputs {
		var x = layer.biases.Data[outputIndex]
		for inputIndex := range input1 {
			var inputValue = input1[inputIndex].Activation
			x += layer.weights.Get(outputIndex, inputIndex) * inputValue
		}
		var offsetIndex = len(input1)
		for _, input := range input2 {
			var inputIndex = offsetIndex + int(input.Index)
			var inputValue = float64(input.Value)
			x += layer.weights.Get(outputIndex, inputIndex) * inputValue
		}
		var n = &layer.outputs[outputIndex]
		n.Activation = layer.activationFn.Sigma(x)
		n.Prime = layer.activationFn.SigmaPrime(x)
	}
}

func (layer *Layer) Backward(input1 []Neuron, input2 []domain.FeatureInfo) {
	for inputIndex := range input1 {
		input1[inputIndex].Error = 0
	}
	for outputIndex := range layer.outputs {
		var n = &layer.outputs[outputIndex]
		var x = n.Error * n.Prime
		for inputIndex := range input1 {
			input1[inputIndex].Error += layer.weights.Get(outputIndex, inputIndex) * x
		}
	}

	for outputIndex := range layer.outputs {
		var n = &layer.outputs[outputIndex]
		var x = n.Error * n.Prime
		layer.bGradients.Add(outputIndex, 0, x*1)

		for inputIndex := range input1 {
			var inputValue = input1[inputIndex].Activation
			layer.wGradients.Add(outputIndex, inputIndex, x*inputValue)
		}

		var offsetIndex = len(input1)
		for _, input := range input2 {
			var inputIndex = offsetIndex + int(input.Index)
			var inputValue = float64(input.Value)
			layer.wGradients.Add(outputIndex, inputIndex, x*inputValue)
		}
	}
}

func (layer *Layer) AddGradients(main *Layer) {
	layer.wGradients.AddTo(&main.wGradients)
	layer.bGradients.AddTo(&main.bGradients)
}

func (layer *Layer) ApplyGradients() {
	layer.wGradients.Apply(&layer.weights)
	layer.bGradients.Apply(&layer.biases)
}
