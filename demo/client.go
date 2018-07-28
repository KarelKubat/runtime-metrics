package main

import (
	"fmt"
	"strings"

	"github.com/KarelKubat/runtime-metrics/reporter"
)

func main() {

	// Instantiate client.
	c, err := reporter.NewClient(":1234")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// See what metrics we have.
	allNames, err := c.AllNames()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Metrics for 'average' %s\n",
		strings.Join(allNames.Averages, ","))
	fmt.Printf("Metrics for 'average per duration' %s\n",
		strings.Join(allNames.AveragesPerDuration, ","))
	fmt.Printf("Metrics for 'counter' %s\n",
		strings.Join(allNames.Counters, ","))
	fmt.Printf("Metrics for 'counter per duration' %s\n",
		strings.Join(allNames.CountersPerDuration, ","))
	fmt.Printf("Metrics for 'sum' %s\n",
		strings.Join(sums, ","))
	fmt.Printf("Metrics for 'sum per duration' %s\n",
		strings.Join(allNames.SumsPerDuration, ","))
}
