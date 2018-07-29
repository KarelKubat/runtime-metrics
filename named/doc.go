/*
named wraps the thread-safe metrics and assigns a name to them.

The constructors all take a string argument, and return a named metric:

 avg := named.NewAverage("my-average")
 avg_per_s := named.NewAveragePerDuration("my-average-per-sec", Time.Second)

 cnt := named.NewCount("my-count")
 cnt_per_s := named.NewcountPerDuration("my-count-per-sec", Time.Second)

 sum := named.NewSum("my-sum")
 sum_per_s := named.NewSumPerDuration("my-sum-per-sec", Time.Second)

The name is returned using method Name():

  fmt.Printf("The metric for averages is named %s\n", avg.Name())

The remainder of the API is identical to the underlying thread-safe
metric (which in its turn is identical to the underlying base metric);
there are methods Mark(), Report() and Reset(). Refer to the base
package for a description.

*/
package named
