# foo
--
    import "github.com/KarelKubat/runtime-metrics"

runtime-metrics implements the collection of in-program metrics, a server to
publish these, and a client to collect them.


What it is and what not

The design principle of runtime-metrics is that it's light-weight enough to be
embedded into a server, and that it won't hinder performance. It can collect and
report simple, one-dimensional metrics: counters, sums and averages - either
over the lifetime of the metric, or over a given period (e.g., counts per
minute).

It cannot handle multi-dimensional data, such as lists of numbers or strings. It
doesn't retain series; it's designed so that the impact of a metric is known
a-priori. If you need to retain historical counters, sums or averages, scrape
them with an external client and retain there.


### Basic In-Program Usage

To collect metrics inside a program and to act on changes; i.e., without
publishing the metrics using a server and without scraping them using a client,
the following can be used:

    import "github.com/KarelKubat/runtime-metrics/base"
    ...

    // Create some metrics. We have:
    // - NewCount() for incremental counting
    // - NewSum() for totalling float64 values
    // - NewAverage() for averaging float64 values
    // and metrics per given a period:
    // - NewCountPerDuration(d time.Duration)
    // - NewSumPerDuration(d time.Duration)
    // - NewAveragePerDuration(d time.Duration)
    totals = base.NewCount()
    failures = base.NewCount()

    // Check failures vs. totals and do something when there is >= 1% failures.
    // Poll every 10 seconds.
    go func() {
      tot := totals.Report()
      fail := failures.Report()
      if tot > 0 {
        ratio = float(fail) / float(tot)
        if ratio > 1.0
          log.Printf("WARNING %v percent of lookups is failing", ratio / 100.0)
        }
      }
      time.Sleep(time.Second * 10)
    }()

    // Loop and track totals and failures
    for {
      err := lookupSomething() // hypothetical function
      totals.Mark()
      if err != nil {
        failures.Mark()
      }
    }

The metric types all have a somewhat similar API: New*() instantiates a metric,
Mark() registers an event, and Report() returns some result. In the case of a
counter, as above, neither Mark() nor Report() have arguments, and Report()
returns an int64. Besides counters, there are metrics for storing sums and
averages.


Thread-safe or not

The base metrics are not thread-safe. When this is needed, then the same types
can be used from package threadsafe:

    import "github.com/KarelKubat/runtime-metrics/threadsafe"
    ...
    totals = threadsafe.NewCount()


### Publishing Metrics

When publishing metrics (to be scraped by a client), metrics are instantiated
with a unique name:

    import "github.com/KarelKubat/runtime-metrics/named"
    ...
    totals = named.Count("total-lookups")
    failures = named.Count("failed-lookups")

Marking events is identical to the non-client-server example above. In order to
publish these metrics, they are added to a registry, and a server is started:

    import "github.com/KarelKubat/runtime-metrics/registry"
    import "github.com/KarelKubat/runtime-metrics/reporter"
    ...
    err := registry.AddCount(totals)
    if err != nil { ... }   // collision of name "total-lookups"
    err = registry.AddCount(failures)
    if err != nil { ... {   // collision of name "total-failures"
    go func() {
      err := reporter.StartReporter(":1234")
      if err != nil { ... } // probably port 1234 is already in use
    }()


### Scraping Metrics

Published metrics can be scraped by a client:

    import "github.com/KarelKubat/runtime-metrics/reporter"

    c, err := reporter.NewClient(":1234")
    if err != nil { ... }  // connection error

    tot, err := c.Count("total-lookups")
    if err != nil { ... }  // counter doesn't exist
    fail, err := c.count("failed-lookups")
    if err != nil { ... }

    if tot > 0 {
      percentage = float(fail) / float(tot) / 100
      if percentage > 0.01 {
        log.Printf("WARNING %v percent of lookups is failing", average)
      }
    }

There is also discovery: a client can get a list of all the names of counts,
sums, and averages, and query these. See demo/client.go for an example.

## Usage
