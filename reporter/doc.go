/*
reporter  implements a metrics reporting server and client.

The Server

The server is started using reporter.StartReporter(addr), where addr
defines the IP and port to listen, separated by :. The IP can be left
out. Example:

  err := reporter.StartReporter(":9000")    // bind to TCP port 9000
  if err != nil { ... }                     // probably port already taken

Typically this will be wrapped in a go function, so that the reporting
server runs in its own thread:

  go func() {
    if err := reporter.StartReporter(":9000"); err != nil {
      ... // probably port 9000 is already taken
    }
  }()

Any metrics that are registered are available for the server to be
published:

  func checkErr(err) {
    if err != nil { ... } // name collision
  }
  ...
  // Create and register metrics
  myAverage := base.NewAverage()
  checkErr(registry.AddAverage("average-metric", myAverage))
  mySumPD := base.NewSumPerDuration(time.Minute)
  checkErr(registry.AddSumPerDuration("sum-per-minute", mySumPD))
  ...
  // As metrics are updated, the reporting server will publish them.
  myAverage.Mark(3.14)
  mySumPD.Mark(2.71)

The Client

The client is instantiated using reporter.NewClient(addr), where again
addr is a string defining the IP address and TCP port:

  c, err := reporter.NewClient(":9000")
  checkErr(err)
  defer c.Close()

The client has a number of handy methods that contact the client to
discover the names of metrics or to fetch metric values:

  allNames, err := c.AllNames()
  // allNames.Averages           is an array of strings (names) of all Average-type metrics
  // allNames.AveragesByDuration is an array of strings (names) of all AveragePerDuration-type metrics
  // allNames.Counts             is an array of strings (names) of all Count-type metrics
  // allNames.CountsPerDuration  is an array of strings (names) of all CountPerDuration-type metrics
  // allNames.Sums               is an array of strings (names) of all Sum-type metrics
  // allNames.SumsPerDuration    is an array of strings (names) of all SumPerDuration-type metrics

In order to fetch the values of a metric, the client calls c.Average(name), c.Sum(name) etc. The
returned values are always what the base type returns, and an error:

  avg, n, err := c.Average("my-average")
  // val is the average
  // n   is the number of observations
  // err is nil or an error

A full dump of all server-known metrics can be obtained using FullDump().
See also demo/client.go for an example.

  dump, err := c.FullDump()
  if err != nil { ... }
  for _, av := range dump.Averages {
    // av.Name is the name, av.Value is the average, av.N is the number of cases
  }
  for _, avPD := range dump.AveragesPerDuration {
    // avPD.Name is the name, avPD.Value is the average, avPD.N is the number of cases,
    // av.Until is the up-to timestamp
  }
  for _, c := range dump.Counts {
    // c.Name is the name, c.Value is the count
  }
  for _, cPD := range dump.CountsPerDuration {
    // cPD.Name is the name, cPD.Value is the count, cPD.Until is the up-to timestamp
  }
  for _, s := range dump.Sums {
    // s.Name is the name, s.Value is the sum, s.N is the number of cases
  }
  for _, sPD := range dump.SumsPerDuration {
    // sPD.Name is the name, sPD.Value is the sum, sPD.N is the number of cases,
    // sPD.Until is the up-to timestamp
  }
*/
package reporter
