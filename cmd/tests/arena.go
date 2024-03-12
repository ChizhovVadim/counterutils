package main

import (
	"context"
	"runtime"

	"github.com/ChizhovVadim/counterutils/internal/arena"
	"github.com/ChizhovVadim/counterutils/internal/engine"
	"github.com/ChizhovVadim/counterutils/internal/evalbuilder"
)

func arenaHandler() error {
	var tc = arena.TimeControl{
		//FixedTime:  1 * time.Second,
		FixedNodes: 2_000_000,
	}

	var gameConcurrency int
	if tc.FixedNodes != 0 {
		gameConcurrency = runtime.NumCPU()
	} else {
		gameConcurrency = runtime.NumCPU() / 2
	}

	return arena.Run(context.Background(), gameConcurrency, tc, newArenaEngine)
}

func newArenaEngine(experiment bool) arena.IEngine {
	return evalBattle(experiment)
	//return engineBattle(experiment)
}

func evalBattle(experiment bool) arena.IEngine {
	var evalName string
	if experiment {
		evalName = "counter"
	} else {
		evalName = "counterold"
	}
	var options = engine.NewMainOptions(evalbuilder.Get(evalName))
	options.Hash = 128
	var eng = engine.NewEngine(options)
	eng.Prepare()
	return eng
}

func engineBattle(experiment bool) arena.IEngine {
	var evalName = "counter"
	var options = engine.NewMainOptions(evalbuilder.Get(evalName))
	options.Hash = 128
	options.ExperimentSettings = experiment
	var eng = engine.NewEngine(options)
	eng.Prepare()
	return eng
}
