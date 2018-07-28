package named

import (
	"github.com/KarelKubat/runtime-metrics/threadsafe"
)

type Average struct {
	handler *threadsafe.Average
	name    string
}

func NewAverage(n string) (*Average, error) {
	metric := &Average{
		handler: threadsafe.NewAverage(),
		name:    n,
	}
	return metric, RegisterAverage(metric)
}

func (a *Average) Mark(val float64) {
	a.handler.Mark(val)
}

func (a *Average) Report() (float64, int64) {
	return a.handler.Report()
}

func (a *Average) Reset() {
	a.handler.Reset()
}

func (a *Average) Name() string {
	return a.name
}
