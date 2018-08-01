package demosrc

import (
	"fmt"
	"strings"

	"github.com/KarelKubat/runtime-metrics/reporter"
)

// ClientAllNamesDemo runs a reporter client, fetches all metric names,
// and displays them.
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
