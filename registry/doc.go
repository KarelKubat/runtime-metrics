/*
registry implements singletons of named sets.

The purpose of this is to uniquely register named metrics using
singletons of named sets. The metrics are then not only unique per
named set, but also per binary. This way, the reporter can pinpoint a
metric by its name and be sure that the metric is unique.

The typical usage is to create base metrics, register them using a
unique name, and start a reporting server. This is also shown in
demo/server.go:

  err := registry.AddAverage("my_average", base.NewAverage())
  if err != nil { ... } // name collision

  err := registry.AddSum("my_sum", base.NewSum())
  if err != nil { ... } // name collision

There are methods to retrieve metrics by their name, which return an
error when the metric isn't in the registry. This is used by the
reporter:

  avg, err := registry.GetAverage("my_average")
  if err != nil { ... } // not found

  sum, err := registry.GetSum("my_sum")
  if err != nil { ... } // not found

A sorted list of the names of all registered averages is retrieved using
AverageNames(), of all counts using CountNames(), etc.:

  for n := range registry.AverageNames() {
    fmt.Printf("An Average metric has the name %s\n", n)
    av, err := registry.GetAverage(n)
    if err != nil { ... } // internal error, should not happen
    val, n, until := av.Report()
    fmt.Printf("This average reports: average value %v, computed over %v values, until time %v\n",
      val, n, until)
  }

*/
package registry
