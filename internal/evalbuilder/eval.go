package evalbuilder

import (
	"fmt"

	counterold "github.com/ChizhovVadim/CounterGo/pkg/eval/counter"
	nnueold "github.com/ChizhovVadim/CounterGo/pkg/eval/nnue"
	counter "github.com/ChizhovVadim/counterutils/internal/eval/counter"
	nnue "github.com/ChizhovVadim/counterutils/internal/eval/nnue"
	weiss "github.com/ChizhovVadim/counterutils/internal/eval/weiss"
)

func Get(key string) func() interface{} {
	return func() interface{} {
		switch key {
		case "weiss":
			return weiss.NewEvaluationService()
		case "counter":
			return counter.NewEvaluationService()
		case "nnue":
			return nnue.NewDefaultEvaluationService()
		case "counterold":
			return counterold.NewEvaluationService()
		case "nnueold":
			return nnueold.NewDefaultEvaluationService()
		}
		panic(fmt.Errorf("bad eval %v", key))
	}
}
