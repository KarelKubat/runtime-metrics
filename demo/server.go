package main

import (
	"fmt"
	"time"

	"github.com/KarelKubat/runtime-metrics/named"
	"github.com/KarelKubat/runtime-metrics/registry"
	"github.com/KarelKubat/runtime-metrics/reporter"
)

func main() {

	// Start up the reporting server on port 1234, all network addresses.
	go func() {
		if err := reporter.StartReporter(":1234"); err != nil {
			panic(fmt.Sprintf("failed to start server: %v", err))
		}
	}()

	// Create some metrics and register them.
	avg := named.NewAverage("some_average")
	if err := registry.AddAverage(avg); err != nil {
		panic(err)
	}

	avgPerSec := named.NewAveragePerDuration("some_average_per_sec",
		time.Duration(time.Second))
	if err := registry.AddAveragePerDuration(avgPerSec); err != nil {
		panic(err)
	}

	cntr := named.NewCounter("some_counter")
	if err := registry.AddCounter(cntr); err != nil {
		panic(err)
	}

	cntrPer5Sec := named.NewCounterPerDuration("some_counter_per_5_sec",
		time.Duration(5*time.Second))
	if err := registry.AddCounterPerDuration(cntrPer5Sec); err != nil {
		panic(err)
	}

	sum := named.NewSum("some_sum")
	if err := registry.AddSum(sum); err != nil {
		panic(err)
	}

	sumPer30Sec := named.NewSumPerDuration("some_sum_per_30_sec",
		time.Duration(30*time.Second))
	if err := registry.AddSumPerDuration(sumPer30Sec); err != nil {
		panic(err)
	}

	// Do stuff to the metrics so that the server may report them.
	for i := 0; ; i++ {
		avg.Mark(float64(i % 10))
		avgPerSec.Mark(float64(i % 10))
		cntr.Mark()
		cntrPer5Sec.Mark()
		sum.Mark(float64(i))
		sumPer30Sec.Mark(float64(i))
	}
}
