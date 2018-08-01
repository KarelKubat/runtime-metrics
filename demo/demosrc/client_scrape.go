package demosrc

import (
	"fmt"
	"time"

	"github.com/KarelKubat/runtime-metrics/reporter"
)

// ClientScrapeDemo runs a reporter client that fetches server metrics in an
// endless loop (or until the server exits). The metrics are displayed on
// stdout.
func ClientScrapeDemo() {

	// Instantiate client.
	c, err := reporter.NewClient(":1234")
	CheckErr(err)
	defer c.Close()

	for {
		fmt.Printf("\n")

		fval, n, err := c.Average("my_average")
		CheckErr(err)
		fmt.Printf("metric my_average: %v over %v values\n", fval, n)

		fval, n, until, err := c.AveragePerDuration("my_average_per_sec")
		CheckErr(err)
		fmt.Printf("metric my_average_per_sec: %v over %v values, sampled until %v\n", fval, n, until)

		ival, err := c.Count("my_counter")
		CheckErr(err)
		fmt.Printf("metric my_counter: %v\n", ival)

		ival, until, err = c.CountPerDuration("my_counter_per_5_sec")
		CheckErr(err)
		fmt.Printf("metric my_counter_per_5_sec: %v, sampled until %v\n", ival, until)

		fval, n, err = c.Sum("my_sum")
		CheckErr(err)
		fmt.Printf("metric my_sum: %v over %v values\n", fval, n)

		fval, n, until, err = c.SumPerDuration("my_sum_per_3_sec")
		CheckErr(err)
		fmt.Printf("metric my_sum_per_3_sec: %v over %v values, sampled until %v\n", fval, n, until)

		time.Sleep(time.Second)
	}
}
