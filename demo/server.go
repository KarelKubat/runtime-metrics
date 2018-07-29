package main

import (
	"time"

	"github.com/KarelKubat/runtime-metrics/named"
	"github.com/KarelKubat/runtime-metrics/registry"
	"github.com/KarelKubat/runtime-metrics/reporter"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// Start up the reporting server on port 1234, all network addresses.
	go func() {
		checkErr(reporter.StartReporter(":1234"))
	}()

	// Create some metrics and register them.
	avg := named.NewAverage("my_average")
	checkErr(registry.AddAverage(avg))

	avgPerSec := named.NewAveragePerDuration("my_average_per_sec",
		time.Duration(time.Second))
	checkErr(registry.AddAveragePerDuration(avgPerSec))

	cntr := named.NewCount("my_counter")
	checkErr(registry.AddCount(cntr))

	cntrPer5Sec := named.NewCountPerDuration("my_counter_per_5_sec",
		time.Duration(5*time.Second))
	checkErr(registry.AddCountPerDuration(cntrPer5Sec))

	sum := named.NewSum("my_sum")
	checkErr(registry.AddSum(sum))

	sumPer30Sec := named.NewSumPerDuration("my_sum_per_30_sec",
		time.Duration(30*time.Second))
	checkErr(registry.AddSumPerDuration(sumPer30Sec))

	// Do stuff to the metrics so that server may report and the client may scrape them.
	for i := 0; ; i++ {
		avg.Mark(float64(i % 10))
		avgPerSec.Mark(float64(i % 10))
		cntr.Mark()
		cntrPer5Sec.Mark()
		sum.Mark(float64(i))
		sumPer30Sec.Mark(float64(i))
		time.Sleep(500 * time.Millisecond)
	}
}
