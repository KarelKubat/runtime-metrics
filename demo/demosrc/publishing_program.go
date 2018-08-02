package demosrc

import (
	"fmt"
	"time"

	"github.com/KarelKubat/runtime-metrics/base"
	"github.com/KarelKubat/runtime-metrics/registry"
)

// PublishingProgramDemo shows how a program can inspect its own metrics, so
// that they can be e.g. pushed to a monitoring service. This is a different
// approach than having the metrics in a server, and having a client scrape
// them and process them further.
//
// Advantage: All-in-one, no need for TCP ports that can already be take
// (because the server/client force you to pick a por);
//
// Disadvantage: Forces your program to be midful of what could go wrong when
// sending metrics to a remote service, how to recover when out of quota,
// etc..

func PublishingProgramDemo() {

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

	// Start the metrics watcher.
	go pushMetrics()

	// Do stuff to the metrics so that server may find them in the above go function.
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

func pushMetrics() {
	// This could push the metrics onto some remote monitoring system. It's the alternative to
	// having a server that just publishes its metrics and a client that scrapes them and processes
	// them further.
	//
	// In each loop we re-query registery.*Names() incase new metrics were created since the
	// previous run.

	for {
		for _, name := range registry.AverageNames() {
			avg, err := registry.AverageBy(name)
			if err == nil {
				val, n := avg.Report()
				fmt.Printf("%q: %v (n=%v)\n", name, val, n)
				// some network call would be here
			} else {
				fmt.Printf("Problem with %q: %v\n", name, err)
			}
		}
		for _, name := range registry.AveragePerDurationNames() {
			avgPD, err := registry.AveragePerDurationBy(name)
			if err == nil {
				val, n, _ := avgPD.Report()
				fmt.Printf("%q: %v (n=%v)\n", name, val, n)
				// some network call would be here
			} else {
				fmt.Printf("Problem with %q: %v\n", name, err)
			}
		}
		for _, name := range registry.CountNames() {
			cnt, err := registry.CountBy(name)
			if err == nil {
				val := cnt.Report()
				fmt.Printf("%q: %v\n", name, val)
				// some network call would be here
			} else {
				fmt.Printf("Problem with %q: %v\n", name, err)
			}
		}
		for _, name := range registry.CountPerDurationNames() {
			cntPD, err := registry.CountPerDurationBy(name)
			if err == nil {
				val, _ := cntPD.Report()
				fmt.Printf("%q: %v\n", name, val)
				// some network call would be here
			} else {
				fmt.Printf("Problem with %q: %v\n", name, err)
			}
		}
		for _, name := range registry.SumNames() {
			avg, err := registry.SumBy(name)
			if err == nil {
				val, n := avg.Report()
				fmt.Printf("%q: %v (n=%v)\n", name, val, n)
				// some network call would be here
			} else {
				fmt.Printf("Problem with %q: %v\n", name, err)
			}
		}
		for _, name := range registry.SumPerDurationNames() {
			avgPD, err := registry.SumPerDurationBy(name)
			if err == nil {
				val, n, _ := avgPD.Report()
				fmt.Printf("%q: %v (n=%v)\n", name, val, n)
				// some network call would be here
			} else {
				fmt.Printf("Problem with %q: %v\n", name, err)
			}
		}

		// Pause for 5 seconds before re-fetching all metrics.
		time.Sleep(5 * time.Second)
	}
}
