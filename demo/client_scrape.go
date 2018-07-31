package main

import (
	"fmt"
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

	for {
		fmt.Printf("\n")

		fval, n, err := c.Average("my_average")
		checkErr(err)
		fmt.Printf("metric my_average: %v over %v values\n", fval, n)

		fval, n, until, err := c.AveragePerDuration("my_average_per_sec")
		checkErr(err)
		fmt.Printf("metric my_average_per_sec: %v over %v values, sampled until %v\n", fval, n, until)

		ival, err := c.Count("my_counter")
		checkErr(err)
		fmt.Printf("metric my_counter: %v\n", ival)

		ival, until, err = c.CountPerDuration("my_counter_per_5_sec")
		checkErr(err)
		fmt.Printf("metric my_counter_per_5_sec: %v, sampled until %v\n", ival, until)

		fval, n, err = c.Sum("my_sum")
		checkErr(err)
		fmt.Printf("metric my_sum: %v over %v values\n", fval, n)

		fval, n, until, err = c.SumPerDuration("my_sum_per_3_sec")
		checkErr(err)
		fmt.Printf("metric my_sum_per_3_sec: %v over %v values, sampled until %v\n", fval, n, until)

		time.Sleep(time.Second)
	}
}
