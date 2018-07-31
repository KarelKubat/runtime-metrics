package main

import (
	"fmt"

	"github.com/KarelKubat/runtime-metrics/reporter"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// Instantiate client.
	c, err := reporter.NewClient(":1234")
	checkErr(err)
	defer c.Close()

	dump, err := c.FullDump()
	checkErr(err)

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
