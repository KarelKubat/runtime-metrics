package demosrc

import (
	"time"

	"github.com/KarelKubat/runtime-metrics/base"
	"github.com/KarelKubat/runtime-metrics/registry"
	"github.com/KarelKubat/runtime-metrics/reporter"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ServerDemo() {

	// Start up the reporting server on port 1234, all network addresses.
	go func() {
		CheckErr(reporter.StartReporter(":1234"))
	}()

	// Create some metrics and register them.
	avg := base.NewAverage()
	CheckErr(registry.AddAverage("my_average", avg))

	avgPerSec := base.NewAveragePerDuration(time.Duration(time.Second))
	CheckErr(registry.AddAveragePerDuration("my_average_per_sec", avgPerSec))

	cntr := base.NewCount()
	CheckErr(registry.AddCount("my_counter", cntr))

	cntrPer5Sec := base.NewCountPerDuration(time.Duration(5 * time.Second))
	CheckErr(registry.AddCountPerDuration("my_counter_per_5_sec", cntrPer5Sec))

	sum := base.NewSum()
	CheckErr(registry.AddSum("my_sum", sum))

	sumPer3Sec := base.NewSumPerDuration(time.Duration(3 * time.Second))
	CheckErr(registry.AddSumPerDuration("my_sum_per_3_sec", sumPer3Sec))

	// Do stuff to the metrics so that server may report and the client may scrape them.
	for i := 0; i < 40; i++ {
		avg.Mark(float64(i % 10))
		avgPerSec.Mark(float64(i % 10))
		cntr.Mark()
		cntrPer5Sec.Mark()
		sum.Mark(float64(i))
		sumPer3Sec.Mark(float64(i))
		time.Sleep(500 * time.Millisecond)
	}
}
