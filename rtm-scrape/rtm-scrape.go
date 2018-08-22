package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/KarelKubat/runtime-metrics/reporter"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// Flag variables
var flagRemoteAddress string
var flagRuns int
var flagInterval time.Duration
var flagAction string
var flagDataSource string
var flagHelp bool
var flagCleanerInterval time.Duration
var flagCompressToMinute time.Duration
var flagCompressTo5Minutes time.Duration
var flagCompressTo15Minutes time.Duration
var flagCompressTo30Minutes time.Duration
var flagCompressTo1Hour time.Duration
var flagDropAfter time.Duration
var flagDebug bool

// CompressionInfo states what flag compresses to what period, after
// what delay it kicks in, and until what delay it should run.
type CompressionInfo struct {
	period   time.Duration
	after    time.Duration
	until    time.Duration
	flagName string
}

var compressionInfo []CompressionInfo

func init() {
	// Bind global flag variables
	flag.StringVar(&flagRemoteAddress, "remote-address", ":1234", "address:port to scrape")
	flag.IntVar(&flagRuns, "runs", 0, "times to run, 0 means forever")
	flag.DurationVar(&flagInterval, "interval", time.Second*5, "delay between runs")
	flag.StringVar(&flagAction, "action", "display",
		"action to perform, one of 'display' or 'store'")
	flag.StringVar(&flagDataSource, "datasource", "",
		"datasource string to open a storage database; something like "+
			"'user=myuser database=mydb'")
	flag.BoolVar(&flagHelp, "h", false, "show help and exit")
	flag.DurationVar(&flagCleanerInterval, "cleaner-interval", 5*time.Minute,
		"delay interval between cleaner/compressor runs")
	flag.DurationVar(&flagCompressToMinute, "compress-to-1-minute", 12*time.Hour,
		"compress datapoints into 1-minute-points when they are older than this duration")
	flag.DurationVar(&flagCompressTo5Minutes, "compress-to-5-minutes", 24*time.Hour,
		"compress datapoints into 5-minute-points when they are older than this duration")
	flag.DurationVar(&flagCompressTo15Minutes, "compress-to-15-minutes", 7*24*time.Hour,
		"compress datapoints into 15-minute-points when they are older than this duration")
	flag.DurationVar(&flagCompressTo30Minutes, "compress-to-30-minutes", 90*24*time.Hour,
		"compress datapoints into 30-minute-points when they are older than this duration")
	flag.DurationVar(&flagCompressTo1Hour, "compress-to-1-hour", 365*24*time.Hour,
		"compress datapoints into 1-hour-points when they are older than this duration")
	flag.DurationVar(&flagDropAfter, "drop-after", 2*365*24*time.Hour,
		"drop datapoints when they are older than this duration")
	flag.BoolVar(&flagDebug, "debug", false, "turn debug logging on")
}

func main() {
	// Check command line
	flag.Parse()
	checkFlags()

	// Ensure that we can access the storage.
	var handler ActionHandler
	switch flagAction {
	case "display":
		handler = &DisplayAction{}
	case "store":
		db, err := sql.Open("postgres", flagDataSource)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to connect to data source: %v\n", err)
		}
		handler, err = NewStoreAction(db, compressionInfo, flagCleanerInterval)
		if err != nil {
			errorAndDie("Failed to initialize storage: %v\n", err)
		}
	}

	// Create the client.
	c, err := reporter.NewClient(flagRemoteAddress)
	if err != nil {
		errorAndDie("Failed to initialize client: %v\n", err)
	}

	// Scrape and process.
	consecutiveErrors := 0
	for i := -1; i < flagRuns; {
		dump, err := c.FullDump()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to fetch full dump: %v\n", err)
			consecutiveErrors++
			if consecutiveErrors > 10 {
				errorAndDie("Consecutive errors exceeds 10, stopping\n")
			}
		} else {
			consecutiveErrors = 0
			if err := handler.HandleFullDump(dump); err != nil {
				errorAndDie("Failed to handle dump: %v\n", err)
			}
		}
		if flagRuns > 0 {
			i++
		}
		time.Sleep(flagInterval)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `
This is rtm-scrape, the real-time metrics scraper.
Supported flags:

`)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, `
At a minimum, --remote-address should be given.
When --action=store, --datasource must be given.
All durations (--compress-to-..., --drop-after, --interval, etc.) are expressed in hours,
minutes and seconds, e.g. '12h30m10s'. Zeroed values can be left out, e.g. '10m'. To turn off
compressing or dropping, set the respective flags to a zero-duration, such as '0s'.
`)
	os.Exit(1)
}

func errorAndDie(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

func checkFlags() {
	if flagHelp {
		usage()
	}

	// Aggressive logging when --debug is given.
	if flagDebug {
		log.SetLevel(log.DebugLevel)
	}

	// Remote address must be given.
	if flagRemoteAddress == "" {
		errorAndDie("Flag --remote-address flag is required\n")
	}

	// Actions store and display are allowed.
	if flagAction != "store" && flagAction != "display" {
		errorAndDie("Unknown --action=%v (allowed: 'display' or 'store')\n", flagAction)
	}

	// When using --action=store, --datasource must be given.
	if flagAction == "store" && flagDataSource == "" {
		errorAndDie("Flag --datasource is required when --action=store\n")
	}

	// Set increasing periods for the compression info
	if flagCompressToMinute != 0 {
		compressionInfo = append(compressionInfo, CompressionInfo{
			period:   time.Minute,
			after:    flagCompressToMinute,
			until:    0,
			flagName: "--compress-to-1-minute",
		})
	}
	if flagCompressTo5Minutes != 0 {
		compressionInfo = append(compressionInfo, CompressionInfo{
			period:   5 * time.Minute,
			after:    flagCompressTo5Minutes,
			until:    0,
			flagName: "--compress-to-5-minutes",
		})
		if len(compressionInfo) > 1 {
			compressionInfo[len(compressionInfo)-2].until = flagCompressTo5Minutes
		}
	}
	if flagCompressTo15Minutes != 0 {
		compressionInfo = append(compressionInfo, CompressionInfo{
			period:   15 * time.Minute,
			after:    flagCompressTo15Minutes,
			until:    0,
			flagName: "--compress-to-15-minutes",
		})
		if len(compressionInfo) > 1 {
			compressionInfo[len(compressionInfo)-2].until = flagCompressTo15Minutes
		}
	}
	if flagCompressTo30Minutes != 0 {
		compressionInfo = append(compressionInfo, CompressionInfo{
			period:   30 * time.Minute,
			after:    flagCompressTo30Minutes,
			until:    0,
			flagName: "--compress-to-30-minutes",
		})
		if len(compressionInfo) > 1 {
			compressionInfo[len(compressionInfo)-2].until = flagCompressTo30Minutes
		}
	}
	if flagCompressTo1Hour != 0 {
		compressionInfo = append(compressionInfo, CompressionInfo{
			period:   time.Hour,
			after:    flagCompressTo1Hour,
			until:    0,
			flagName: "--compress-to-30-minutes",
		})
		if len(compressionInfo) > 1 {
			compressionInfo[len(compressionInfo)-2].until = flagCompressTo1Hour
		}
	}
	if flagDropAfter != 0 {
		compressionInfo = append(compressionInfo, CompressionInfo{
			period:   0,
			after:    flagDropAfter,
			until:    0,
			flagName: "--drop-after",
		})
		if len(compressionInfo) > 1 {
			compressionInfo[len(compressionInfo)-2].until = flagDropAfter
		}
	}

	for i := range compressionInfo {
		log.WithFields(log.Fields{
			"flag":   compressionInfo[i].flagName,
			"period": compressionInfo[i].period,
			"until":  compressionInfo[i].until,
			"after":  compressionInfo[i].after,
		}).Debug("compression")
		if compressionInfo[i].after < 0 {
			errorAndDie("%s may not specify a negative duration\n", compressionInfo[i].flagName)
		}
		if compressionInfo[i].period > 0 && compressionInfo[i].after > compressionInfo[i].until {
			errorAndDie("%s duration conflicts with the next compression\n",
				compressionInfo[i].flagName)
		}
	}
}
