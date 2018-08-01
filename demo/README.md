# demo
--
The files in demo/ and demosrc/ implement a server and a client.

The server instantiates some metrics and starts a reporter. Then, in a loop, it
manipulates the metrics, and stops after 20 seconds. To run the server, type

    go run main.go server &

so that the server is started in the background. You have 20 seconds to try out
the clients that are listed below.

Client client_allnames.go discovers the names of metrics. To try it, run

    go run main.go allnames

Client client_scrape.go scrapes predefined metrics ("my_average", "my_sum" etc.,
which are fortunately published by the server) and prints their values on
stdout. This is endlessly repeated until the server shuts down and the client
fails. This client could be combined with the name discovery from
client_allnames.go to not scrape predefined, but available names -- but the
source is kept short for demo purposes. To try it, type

    go run main.go scraper

Client client_fulldump.go requests a full metric dump from the server. All known
metrics and their values are shown. To try it, type

    go run main.go fulldump

For reference, (some version of) the server and clients are shown below. Refer
to the source files under demosrc/for the most recent versions.

### The Server

     package demosrc

     import (
    	"time"

    	"github.com/KarelKubat/runtime-metrics/base"
    	"github.com/KarelKubat/runtime-metrics/registry"
    	"github.com/KarelKubat/runtime-metrics/reporter"
     )

     func CheckErr(err error) {
    	if err != nil {
    		panic(err)
    	}
     }

     func ServerDemo() {

    	// Start up the reporting server on port 1234, all network addresses.
    	go func() {
    		CheckErr(reporter.StartReporter(":1234"))
    	}()

    	// Create some metrics and register them.
    	avg := base.NewAverage()
    	CheckErr(registry.AddAverage("my_average", avg))

    	avgPerSec := base.NewAveragePerDuration(time.Duration(time.Second))
    	CheckErr(registry.AddAveragePerDuration("my_average_per_sec", avgPerSec))

    	cntr := base.NewCount()
    	CheckErr(registry.AddCount("my_counter", cntr))

    	cntrPer5Sec := base.NewCountPerDuration(time.Duration(5 * time.Second))
    	CheckErr(registry.AddCountPerDuration("my_counter_per_5_sec", cntrPer5Sec))

    	sum := base.NewSum()
    	CheckErr(registry.AddSum("my_sum", sum))

    	sumPer3Sec := base.NewSumPerDuration(time.Duration(3 * time.Second))
    	CheckErr(registry.AddSumPerDuration("my_sum_per_3_sec", sumPer3Sec))

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

### The All Names Client

     package demosrc

     import (
    	"fmt"
    	"strings"

    	"github.com/KarelKubat/runtime-metrics/reporter"
     )

     func ClientAllNamesDemo() {

    	// Instantiate client.
    	c, err := reporter.NewClient(":1234")
    	CheckErr(err)
    	defer c.Close()

    	// See what metrics we have.
    	allNames, err := c.AllNames()
    	CheckErr(err)

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

### The Scraper Client

     package demosrc

     import (
    	"fmt"
    	"time"

    	"github.com/KarelKubat/runtime-metrics/reporter"
     )

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

### The Full Dump Client

     package demosrc

     import (
    	"fmt"

    	"github.com/KarelKubat/runtime-metrics/reporter"
     )

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


### A Self Publishing Program

If you don't want a server publishing metrics, and a client scraping them; but
instead want a program that pushes metrics as they change, then have a look at
demosrc/publishing_program.go.

The approach is as follows: the below shown function publishMetrics() asks the
registry for names, and using each name, gets the appropriate metric (sum,
average, etc.). The values of the metric are returned by its Report() function.
This is repeated ad infinitum with some delay between runs. The function would
typically be started as a go-routine.

     // somewhere in the code...
     go pushMetrics()

     func pushMetrics() {
    	// This could push the metrics onto some remote monitoring system. It's the alternative to
    	// having a server that just publishes its metrics and a client that scrapes them and processes
    	// them further.
        //
        // It would be started as a go-routine to run in a separate thread.
    	//
    	// In each loop we re-query registery.*Names() incase new metrics were created since the
    	// previous run.

    	for {
    		for _, name := range registry.AverageNames() {
    			avg, err := registry.GetAverage(name)
    			if err == nil {
    				val, n := avg.Report()
    				fmt.Printf("%q: %v (n=%v)\n", name, val, n)
    				// some network call would be here
    			} else {
    				fmt.Printf("Problem with %q: %v\n", name, err)
    			}
    		}
    		for _, name := range registry.AveragePerDurationNames() {
    			avgPD, err := registry.GetAveragePerDuration(name)
    			if err == nil {
    				val, n, _ := avgPD.Report()
    				fmt.Printf("%q: %v (n=%v)\n", name, val, n)
    				// some network call would be here
    			} else {
    				fmt.Printf("Problem with %q: %v\n", name, err)
    			}
    		}
    		for _, name := range registry.CountNames() {
    			cnt, err := registry.GetCount(name)
    			if err == nil {
    				val := cnt.Report()
    				fmt.Printf("%q: %v\n", name, val)
    				// some network call would be here
    			} else {
    				fmt.Printf("Problem with %q: %v\n", name, err)
    			}
    		}
    		for _, name := range registry.CountPerDurationNames() {
    			cntPD, err := registry.GetCountPerDuration(name)
    			if err == nil {
    				val, _ := cntPD.Report()
    				fmt.Printf("%q: %v\n", name, val)
    				// some network call would be here
    			} else {
    				fmt.Printf("Problem with %q: %v\n", name, err)
    			}
    		}
    		for _, name := range registry.SumNames() {
    			avg, err := registry.GetSum(name)
    			if err == nil {
    				val, n := avg.Report()
    				fmt.Printf("%q: %v (n=%v)\n", name, val, n)
    				// some network call would be here
    			} else {
    				fmt.Printf("Problem with %q: %v\n", name, err)
    			}
    		}
    		for _, name := range registry.SumPerDurationNames() {
    			avgPD, err := registry.GetSumPerDuration(name)
    			if err == nil {
    				val, n, _ := avgPD.Report()
    				fmt.Printf("%q: %v (n=%v)\n", name, val, n)
    				// some network call would be here
    			} else {
    				fmt.Printf("Problem with %q: %v\n", name, err)
    			}
    		}

    		// Pause for 5 seconds before re-fetching all metrics.
    		time.Sleep(5 * time.Second)
    	}
     }
