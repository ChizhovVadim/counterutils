package main

import (
	"context"
	"log"
	"os"

	"github.com/ChizhovVadim/CounterGo/pkg/common"
	"github.com/ChizhovVadim/counterutils/internal/engine"
	"github.com/ChizhovVadim/counterutils/internal/evalbuilder"
)

var cliArgs = NewCommandArgs(os.Args)
var tacticTestsPath = mapPath("~/chess/tests/tests.epd")

func main() {
	var err = run()
	if err != nil {
		log.Println(err)
	}
}

func run() error {
	var cli = NewCommandHandler()
	cli.Add("tactic", tacticHandler)
	cli.Add("arena", arenaHandler)
	return cli.Execute(cliArgs.CommandName())
}

type UciEngine interface {
	Search(ctx context.Context, searchParams common.SearchParams) common.SearchInfo
}

func newEngine(evalName string) *engine.Engine {
	var options = engine.NewMainOptions(evalbuilder.Get(evalName))
	options.Hash = 128
	var eng = engine.NewEngine(options)
	return eng
}
