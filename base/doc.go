/*
base defines the most basic metric types.

The types are:

 base.Count        // records ticks
 base.Sum          // adds float64 values
 base.Average      // just a complimentary type very similar to Sum

Besides these, *PerDuration exist, to record over a given period, which gives
the additional:

 base.CountPerDuration
 base.SumPerDuration
 base.AveragePerDuration

New metric instances are created using New*, such as NewCount(),
NewSum(), and so on. Marking an event in a metric is done using the
Mark() method, which takes different arguments, depending on the
metric. Resetting a metric is done using Reset(). Seeing what the
state of the metric is, is done using Report(), which returns
different values, depeinding on the metric.

Example:

 av := base.NewAverage()
 av.Mark(1.0)
 av.Mark(2.0)
 average, n := av.Report()
 fmt.Printf("average is %v, over %v values\n", average, n) // 1.5 over 2 values

 av.Reset()
 av.Mark(4.0)
 av.Mark(5.0)
 average, n = av.Report()
 fmt.Printf("average is %v, over %v values\n", average, n) // 4.5 over 2 values

All *PerDuration metrics are created with a time.Duration specifier,
which is the "width" of the observation window. NOTE THAT when
Report() is called on these metrics, then values are returned that
were obtained somewhere in the past. This is best illustrated in a
simple diagram:

                      av := base.NewAveragePerDuration(time.Second)
  T=0s
   |                  av.Mark(1.0)
   |                  av.Mark(2.0)
   |                  val0, n0, until0 := av.Report()   // state until T=0
  T=1s
   |                  av.Mark(20.0)
   |                  av.Mark(10.0)
   |                  val1, n1, until1 := av.Report()   // state between T=0 and T=1
  T=2s
   |                  av.Mark(200.0)
   |                  av.Mark(100.0)
   |                  val2, n2, until2 := av.Report()   // state between T=1 and T=2
  T=3
   |                  val3, n3, until3 := av.Report     // state between T=2 and T=3
   |                                                    // val3 == 150.0

Here, val0 and n0 will both be zero, because the very first period is
not finished yet when Report() is called. Once one second has expired
(T=1), val1 will report 1.5, and the marked values 10.0 and 20.0 are
again not taken into account, as they fall into the currently updating
period. Only after T=2 has passed, Report() will return 15.0 as the average.

For all *PerDuration metrics, Report() also returns the timestamp until which
the marked values are taken into account.

*/
package base
