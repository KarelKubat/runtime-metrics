# foo
--
    import "github.com/KarelKubat/runtime-metrics"

runtime-metrics implements the collection of in-program metrics, a server to
publish these, and a client to collect them.


What is it

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
the following can be used. As an example, we wish to track the ratio of
failures, and log something when this ratio exceeds 1% over a period of 30
seconds.

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
    errorAverage = base.NewAveragePerDuration(30 * time.Second)

    // Check failures vs. totals and do something when there is >= 1% failures.
    // Poll every 10 seconds.
    go func() {
      average, n, until := errorAverage.Report()
      if average >= 0.01 {
        log.Printf("WARNING %v percent of lookups is failing " +
          "over a period of 30 seconds until %v, %v cases ",
          ratio * 100.0, until, n)
        }
      }
      time.Sleep(time.Second * 30)
    }()

    // Loop and track totals and failures
    for {
      err := lookupSomething() // hypothetical function
      if err != nil {
        errorAverage.Mark(1.0)
      } else {
        errorAverage.Mark(0.0)
    }

It should be noted here that there are different ways to solve this. One could
also use two counters, one for the total loops and one for the failures. In this
case, it's good to limit the collection of metrics and their reporting to a
given duration; otherwise, a long run of successes might mask suddenly occurring
errors until it's too late.

The metric types all have a somewhat similar API: New*() instantiates a metric,
Mark() registers an event, and Report() returns some result. In the case of an
average, Mark() expects one float64 argument, and returns three values: average,
number of cases, and a timestamp.


Threadsafe or not

The base metrics are not thread-safe. When this is needed, then the same types
can be used from package threadsafe:

    import "github.com/KarelKubat/runtime-metrics/threadsafe"
    ...
    totals = threadsafe.NewAveragePerDuration(30 * time.Second)


### Publishing Metrics

When publishing metrics (to be scraped by a client), metrics are instantiated
with a unique name using package named:

    import "github.com/KarelKubat/runtime-metrics/named"
    ...
    errorAverage := named.AveragePerDuration(30 * time.Second)

Marking events is identical to the non-client-server example above. In order to
publish these metrics, they are added to a registry, and a server is started:

    import "github.com/KarelKubat/runtime-metrics/registry"
    import "github.com/KarelKubat/runtime-metrics/reporter"
    ...

    errorAverage := named.NewAveragePerDuration(
      "lookup-error-ratio-per-30-sec",    // uniquely identifying name
      30 * time.Second)				    // standard argument

    err := registry.AddCount(errorAverage)
    if err != nil { ... }                 // collision of name

    go func() {
      err := reporter.StartReporter(":1234")
      if err != nil { ... } // probably port 1234 is already in use
    }()


### Scraping Metrics

Published metrics can be scraped by a client:

    import "github.com/KarelKubat/runtime-metrics/reporter"

    c, err := reporter.NewClient(":1234")
    if err != nil { ... }  // connection error

    av, n, until, err := c.AveragePerDuration("lookup-error-ratio-per-30-sec")
    if err != nil { ... }  // metric doesn't exist

    if av > 0.01 {
      log.Printf("WARNING %v percent of lookups is failing", av * 100)
    }

There is also discovery: a client can get a list of all the names of counts,
sums, and averages, and query these. See demo/client.go for an example.

## Usage
