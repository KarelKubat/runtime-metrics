package demosrc

import (
	"fmt"

	"github.com/KarelKubat/runtime-metrics/reporter"
)

// ClientFullDumpDemo runs a reporter client, obtains a full metrics dump from
// the server, and displays all metrics.
func ClientFullDumpDemo() {

	// Instantiate client.
	c, err := reporter.NewClient(":1234")
	CheckErr(err)
	defer c.Close()

	dump, err := c.FullDump()
	CheckErr(err)

	for _, av := range dump.Averages {
		fmt.Printf("Average metric %q: average=%v over %v cases\n",
			av.Name, av.Value, av.N)
	}

	for _, av := range dump.AveragesPerDuration {
		fmt.Printf("Average per duration metric %q: average=%v over %v cases, measured until %v cases\n",
			av.Name, av.Value, av.N, av.Until)
	}

	for _, av := range dump.Counts {
		fmt.Printf("Count metric %q: count=%v\n",
			av.Name, av.Value)
	}

	for _, av := range dump.CountsPerDuration {
		fmt.Printf("Count per duration metric %q: count=%v, measured until %v cases\n",
			av.Name, av.Value, av.Until)
	}

	for _, av := range dump.Sums {
		fmt.Printf("Sum metric %q: sum=%v over %v cases\n",
			av.Name, av.Value, av.N)
	}

	for _, av := range dump.SumsPerDuration {
		fmt.Printf("Sum per duration metric %q: sum=%v over %v cases, measured until %v cases\n",
			av.Name, av.Value, av.N, av.Until)
	}
}
