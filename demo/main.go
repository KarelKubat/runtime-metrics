package main

import (
	"fmt"
	"github.com/KarelKubat/runtime-metrics/demo/demosrc"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr,
		`Usage: One argument is needed to specify what to demo. This may be:
   server           -  runs a 20 seconds server that publishes metrics
   allnames         -  runs a client that fetches the names of all metrics from
                       the server
   scraper          -  runs a client that repetitively fetches metrics and shows
                       them as they change
   fulldump         -  runs a client that gets one full dump of the server metrics
                       and shows them
   publishingprogram - how to detect when metrics change inside your own program
                       so that you could push them somewhere
   foreverserver     - like server, but runs forever (for benchmarking etc.)
`)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}
	switch os.Args[1] {
	case "server":
		demosrc.ServerDemo()
	case "allnames":
		demosrc.ClientAllNamesDemo()
	case "scraper":
		demosrc.ClientScrapeDemo()
	case "fulldump":
		demosrc.ClientFullDumpDemo()
	case "publishingprogram":
		demosrc.PublishingProgramDemo()
	case "foreverserver":
		demosrc.ForeverServerDemo()
	default:
		usage()
	}
}
