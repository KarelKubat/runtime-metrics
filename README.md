# runtimemetrics
--
    import "github.com/KarelKubat/runtime-metrics"

runtime-metrics implements the collection of in-program metrics, a server to
publish these, and a client to collect them.


### What Is It

The design principle of runtime-metrics is that it's light-weight enough to be
embedded into a server, and that it won't hinder performance. It can collect and
report simple, one-dimensional metrics: counters, sums and averages - either
over the lifetime of the metric, or over a given period (e.g., counts per
minute).

It cannot handle multi-dimensional data, such as lists of numbers or strings. It
doesn't retain series; it's designed so that the impact of a metric is known
a-priori. If you need to retain lists of counters, sums or averages (e.g., to
analyze trends), then scrape them with an external client and retain there.


### Basic In-Program Usage

To collect metrics inside a program and to act on changes; i.e., without
publishing the metrics using a server and without scraping them using a client,
the following can be used. As an example, we wish to track the ratio of
failures, and log something when this ratio exceeds 1% over a period of 30
seconds.

    import "github.com/KarelKubat/runtime-metrics/base"
    ...

    // Create the metrics. We have:
    // - NewCount() for incremental counting
    // - NewSum() for totalling float64 values
    // - NewAverage() for averaging float64 values
    // and metrics per given a period:
    // - NewCountPerDuration(d time.Duration)
    // - NewSumPerDuration(d time.Duration)
    // - NewAveragePerDuration(d time.Duration)

    errorRatio = base.NewAveragePerDuration(30 * time.Second)

    // Check failures vs. totals and do something when there is >= 1% failures.
    // Poll every 30 seconds, a shorter period won't help because the average
    // cannot change any quicker.
    go func() {
      // average is the recorded value, n is the number of cases,
      // until is the up-to timestamp of the calculation
      average, n, until := errorRatio.Report()
      if average >= 0.01 {
        log.Printf("WARNING %v percent of lookups is failing " +
          "over a period of 30 seconds until %v, %v cases ",
          average * 100.0, until, n)
        }
      }
      time.Sleep(time.Second * 30)
    }()

    // Loop and track totals and failures
    for {
      err := lookupSomething()  // hypothetical function
      if err != nil {
        errorRatio.Mark(1.0)    // mark error (and increase #-cases)
      } else {
        errorRatio.Mark(0.0)    // mark success (only increase #-cases)
    }

It should be noted here that there are different ways to solve this. One could
also use two counters, one for the total loops and one for the failures, and
divide them to get a ratio.

In this case, it's also good to limit the collection of metrics and their
reporting to a given duration; otherwise, a long run of successes might mask
suddenly occurring errors until it's too late.

The metric types all have a somewhat similar API: New*() instantiates a metric,
Mark() registers an event, and Report() returns some result. In the case of an
average, Mark() expects one float64 argument, and returns three values: average,
number of cases, and a timestamp.


### Publishing Metrics

In order to publish metrics, they are added to a registry, and a server is
started:

    import "github.com/KarelKubat/runtime-metrics/base"
    import "github.com/KarelKubat/runtime-metrics/registry"
    import "github.com/KarelKubat/runtime-metrics/reporter"
    ...

    errorRatio := base.NewAveragePerDuration(30 * time.Second)
    err := registry.AddAveragePerDuration("lookup-error-ratio-per-30-sec", errorRatio)
    if err != nil { ... }                 // collision of name

    go func() {
      err := reporter.StartReporter(":1234")
      if err != nil { ... }               // probably port 1234 is already in use
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


### Further Reading

Be sure to read the docs of the base/ package to understand what the metrics do.
Particularly make sure you understand how *PerDuration metrics work, they always
report results that "lag" a duration.

For information how to publish metrics, read the server section in package
reporter/. For scraping, read the client section. An example is provided in
demo/.

Packages tools/, namedset/ and api/ are for internal usage.

## Usage
