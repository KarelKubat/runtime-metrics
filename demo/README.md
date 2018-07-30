# demo
--
The files in demo implement a server and a client.

The server instantiates some metrics and starts a reporter. Then, in an endless
loop, it manipulates the metrics. The client discovers dials the server,
discovers metrics, and prints their values on stdout.

To run the examples, in one terminal start

    go run server.go

In a different terminal, start

    go run client.go

To stop the test, hit ^C in both the terminals. The client and server are
hardwired to use TCP port 1234. If that port is already taken on your system,
then you'll have to edit the sources.

For reference, (some version of) the server and client are shown below. Refer to
the source files for the most recent version.

### The Server

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

    	sumPer30Sec := base.NewSumPerDuration(time.Duration(30 * time.Second))
    	checkErr(registry.AddSumPerDuration("my_sum_per_30_sec", sumPer30Sec))

    	// Do stuff to the metrics so that server may report and the client may scrape them.
    	for i := 0; ; i++ {
    		avg.Mark(float64(i % 10))
    		avgPerSec.Mark(float64(i % 10))
    		cntr.Mark()
    		cntrPer5Sec.Mark()
    		sum.Mark(float64(i))
    		sumPer30Sec.Mark(float64(i))
    		time.Sleep(500 * time.Millisecond)
    	}
      }

### The Client

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
      			fmt.Printf("%q: %v over %v values, sampled until %v\n", name, val, n, until                      )
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
