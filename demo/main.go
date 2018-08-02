package main

import (
	"fmt"
	"github.com/KarelKubat/runtime-metrics/demo/demosrc"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr,
		`Usage:
  demo server           -  runs a 20 seconds server that publishes metrics
  demo allnames         -  runs a client that fetches the names of all metrics from
                           the server
  demo scraper          -  runs a client that repetitively fetches metrics and shows
                           them as they change
  demo fulldump         -  runs a client that gets one full dump of the server metrics
                           and shows them
  demo publishingprogram - how to detect when metrics change inside your own program
                           so that you could push them somewhere
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
	default:
		usage()
	}
}
