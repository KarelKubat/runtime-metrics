package main

import (
	"fmt"
	"strings"

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

	// See what metrics we have.
	allNames, err := c.AllNames()
	checkErr(err)

	fmt.Printf("Metrics for 'average': %s\n",
		strings.Join(allNames.Averages, ","))
	fmt.Printf("Metrics for 'average per duration': %s\n",
		strings.Join(allNames.AveragesPerDuration, ","))
	fmt.Printf("Metrics for 'counter': %s\n",
		strings.Join(allNames.Counts, ","))
	fmt.Printf("Metrics for 'counter per duration': %s\n",
		strings.Join(allNames.CountsPerDuration, ","))
	fmt.Printf("Metrics for 'sum': %s\n",
		strings.Join(allNames.Sums, ","))
	fmt.Printf("Metrics for 'sum per duration': %s\n",
		strings.Join(allNames.SumsPerDuration, ","))
}
