package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/KarelKubat/runtime-metrics/reporter"
)

func main() {
	// Supported flags
	remoteAddress := flag.String("remote-address", ":1234", "address:port to scrape")
	runs := flag.Int("runs", 0, "times to run, 0 means forever")
	interval := flag.Duration("interval", time.Second*5, "delay between runs")
	action := flag.String("action", "display", "action to perform, one of display or store")
	datasource := flag.String("datasource", "", "datasource string to open a storage database")
	driver := flag.String("driver", "",
		"db driver to access --datasource, available: sqlite3 or postgres")
	h := flag.Bool("h", false, "show help and exit")

	// Check command line
	flag.Parse()

	if *h {
		usage()
	}
	if *remoteAddress == "" {
		fmt.Fprintf(os.Stderr, "Flag --remote-address flag is required\n")
		os.Exit(1)
	}
	if *action == "store" {
		if *datasource == "" || *driver == "" {
			fmt.Fprintf(os.Stderr, "Flags --datasource and --driver are required when --action=store")
			os.Exit(1)
		}
		if *driver != "sqlite3" && *driver != "postgres" {
			fmt.Fprintf(os.Stderr, "Only implemented drivers are --driver=sqlite3 or --driver=postgres")
			os.Exit(1)
		}
	}

	// Ensure that we can access the storage.
	var handler ActionHandler
	switch *action {
	case "display":
		handler = &DisplayAction{}
	case "store":
		db, err := sql.Open(*driver, *datasource)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to connect to data source: %v\n", err)
		}
		handler = &StoreAction{
			DB:     db,
			Driver: driver,
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown --action=%v\n", *action)
		os.Exit(1)
	}
	if err := handler.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize handler: %v\n", err)
		os.Exit(1)
	}

	// Create the client.
	c, err := reporter.NewClient(*remoteAddress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize client: %v\n", err)
		os.Exit(1)
	}

	// Scrape and process.
	consecutiveErrors := 0
	for i := -1; i < *runs; {
		dump, err := c.FullDump()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to fetch full dump: %v\n", err)
			consecutiveErrors++
			if consecutiveErrors > 10 {
				fmt.Fprintf(os.Stderr, "Consecutive errors exceeds 10, stopping\n")
				os.Exit(1)
			}
		} else {
			consecutiveErrors = 0
			if err := handler.HandleFullDump(dump); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to handle dump: %v\n", err)
				os.Exit(1)
			}
		}
		if *runs > 0 {
			i++
		}
		time.Sleep(*interval)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Supported flag:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nAt a minimum, --remote-address must be given.\n")
	fmt.Fprintf(os.Stderr, "When --action=sql, --driver and --datasource must be given.\n")

	os.Exit(1)
}
