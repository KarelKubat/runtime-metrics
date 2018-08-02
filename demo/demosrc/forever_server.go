package demosrc

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/KarelKubat/runtime-metrics/base"
	"github.com/KarelKubat/runtime-metrics/registry"
	"github.com/KarelKubat/runtime-metrics/reporter"
)

// ForeverServerDemo simulates what a real-life server might do to obtain and
// to publish its metrics.
func ForeverServerDemo() {

	// Start up the reporting server on port 1234, all network addresses.
	go func() {
		CheckErr(reporter.StartReporter(":1234"))
	}()

	// Create and register metrics.
	errorRatioPer10Sec := base.NewAveragePerDuration(10 * time.Second)
	latencyPerSec := base.NewAveragePerDuration(time.Second)
	latencyPer10Sec := base.NewAveragePerDuration(10 * time.Second)
	callsPerSec := base.NewCountPerDuration(time.Second)
	callsPer10Sec := base.NewCountPerDuration(time.Minute)

	CheckErr(registry.AddAveragePerDuration("Error ratio (over 10s)", errorRatioPer10Sec))
	CheckErr(registry.AddAveragePerDuration("Latency (ms, over 1s)", latencyPerSec))
	CheckErr(registry.AddAveragePerDuration("Latency (ms, over 10s)", latencyPer10Sec))
	CheckErr(registry.AddCountPerDuration("Calls/s", callsPerSec))
	CheckErr(registry.AddCountPerDuration("Calls/10s", callsPer10Sec))

	// Simulate forever.
	for {
		start := time.Now()
		err := dummyCall()
		end := time.Now()

		if err != nil {
			errorRatioPer10Sec.Mark(1.0)
		} else {
			errorRatioPer10Sec.Mark(0.0)
		}

		latencyPerSec.Mark(math.Floor(float64(end.Sub(start) / time.Millisecond)))
		latencyPer10Sec.Mark(math.Floor(float64(end.Sub(start) / time.Millisecond)))

		callsPerSec.Mark()
		callsPer10Sec.Mark()
	}
}

// This might be an RPC to someplace.
func dummyCall() error {
	// Wait for some time, let's assume latency is 100-600ms
	time.Sleep(time.Duration(100+rand.Intn(500)) * time.Millisecond)
	// Simulate an error in approx. 10% of the cases
	if rand.Intn(100) <= 10 {
		return fmt.Errorf("dummyCall returns an error")
	}
	return nil
}
