/*

The files in demo implement a server and a client.

The server instantiates some metrics and starts a reporter. Then, in
an endless loop, it manipulates the metrics. To run the server, type

  go run server.go

The server will run for approx. 20 seconds and then shut down.


Client client_allnames.go discovers the names of metrics. To try it,
run

  go client_allnames.go

Client client_scrape.go scrapes predefined metrics ("my_average",
"my_sum" etc., which are fortunately published by the server) and
prints their values on stdout. This is endlessly repeated until the
server shuts down and the client fails. This client could be
combined with the name discovery from client_allnames.go to not
scrape predefined, but available names -- but the source is kept
short for demo purposes.

Client client_fulldump.go requests a full metric dump from the server.
All known metrics and their values are shown.

For reference, (some version of) the server and clients are shown below.
Refer to the source files for the most recent versions.

The Server

 package main

 import (
	"time"

	"github.com/KarelKubat/runtime-metrics/base"
	"github.com/KarelKubat/runtime-metrics/registry"
	"github.com/KarelKubat/runtime-metrics/reporter"
 )

 func checkErr(err error) {
	if err != nil {
		panic(err)
	}
 }

 func main() {

	// Start up the reporting server on port 1234, all network addresses.
	go func() {
		checkErr(reporter.StartReporter(":1234"))
	}()

	// Create some metrics and register them.
	avg := base.NewAverage()
	checkErr(registry.AddAverage("my_average", avg))

	avgPerSec := base.NewAveragePerDuration(time.Duration(time.Second))
	checkErr(registry.AddAveragePerDuration("my_average_per_sec", avgPerSec))

	cntr := base.NewCount()
	checkErr(registry.AddCount("my_counter", cntr))

	cntrPer5Sec := base.NewCountPerDuration(time.Duration(5 * time.Second))
	checkErr(registry.AddCountPerDuration("my_counter_per_5_sec", cntrPer5Sec))

	sum := base.NewSum()
	checkErr(registry.AddSum("my_sum", sum))

	sumPer3Sec := base.NewSumPerDuration(time.Duration(3 * time.Second))
	checkErr(registry.AddSumPerDuration("my_sum_per_3_sec", sumPer3Sec))

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

The All Names Client

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

The Scraper Client

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

The Full Dump Client

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
*/
package main
