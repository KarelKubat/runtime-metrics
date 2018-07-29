package main

import (
	"fmt"
	"strings"
	"time"

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

	fmt.Printf("Metrics for 'average' %s\n",
		strings.Join(allNames.Averages, ","))
	fmt.Printf("Metrics for 'average per duration' %s\n",
		strings.Join(allNames.AveragesPerDuration, ","))
	fmt.Printf("Metrics for 'counter' %s\n",
		strings.Join(allNames.Counts, ","))
	fmt.Printf("Metrics for 'counter per duration' %s\n",
		strings.Join(allNames.CountsPerDuration, ","))
	fmt.Printf("Metrics for 'sum' %s\n",
		strings.Join(allNames.Sums, ","))
	fmt.Printf("Metrics for 'sum per duration' %s\n",
		strings.Join(allNames.SumsPerDuration, ","))

	for {
		fmt.Printf("\n")

		for _, name := range allNames.Averages {
			val, n, err := c.Average(name)
			checkErr(err)
			fmt.Printf("%q: %v over %v values\n", name, val, n)
		}

		for _, name := range allNames.AveragesPerDuration {
			val, n, until, err := c.AveragePerDuration(name)
			checkErr(err)
			fmt.Printf("%q: %v over %v values, sampled until %v\n", name, val, n, until)
		}
		for _, name := range allNames.Counts {
			val, err := c.Count(name)
			checkErr(err)
			fmt.Printf("%q: %v\n", name, val)
		}

		for _, name := range allNames.CountsPerDuration {
			val, until, err := c.CountPerDuration(name)
			checkErr(err)
			fmt.Printf("%q: %v, sampled until %v\n", name, val, until)
		}
		for _, name := range allNames.Sums {
			val, n, err := c.Sum(name)
			checkErr(err)
			fmt.Printf("%q: %v over %v values\n", name, val, n)
		}

		for _, name := range allNames.SumsPerDuration {
			val, n, until, err := c.SumPerDuration(name)
			checkErr(err)
			fmt.Printf("%q: %v over %v values, sampled until %v\n", name, val, n, until)
		}

		time.Sleep(time.Second)
	}
}
