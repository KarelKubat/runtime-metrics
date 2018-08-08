package main

import (
	"fmt"

	"github.com/KarelKubat/runtime-metrics/reporter"
)

// DisplayAction handles displaying full metrics dumps.
type DisplayAction struct {
}

// Init doesn't perform any tasks with DisplayAction.
func (d *DisplayAction) Init() error {
	return nil
}

// HandleFullDump implements ActionHandler to display a full metrics dump.
func (d *DisplayAction) HandleFullDump(dump *reporter.FullDumpReturn) error {
	maxLen := longestMetricName(dump)
	fmt.Printf("%-*s  %10s %6s %s\n", maxLen, "NAME", "VALUE", "N", "UNTIL")
	for _, m := range dump.Averages {
		fmt.Printf("%-*s  %10.2f %6d\n", maxLen, m.Name, m.Value, m.N)
	}
	for _, m := range dump.AveragesPerDuration {
		fmt.Printf("%-*s  %10.2f %6d %v\n", maxLen, m.Name, m.Value, m.N, m.Until)
	}
	for _, m := range dump.Counts {
		fmt.Printf("%-*s  %10d\n", maxLen, m.Name, m.Value)
	}
	for _, m := range dump.CountsPerDuration {
		fmt.Printf("%-*s  %10d %6s %v\n", maxLen, m.Name, m.Value, "", m.Until)
	}
	for _, m := range dump.Sums {
		fmt.Printf("%-*s  %10.2f %6d\n", maxLen, m.Name, m.Value, m.N)
	}
	for _, m := range dump.SumsPerDuration {
		fmt.Printf("%-*s  %10.2f %6d %v\n", maxLen, m.Name, m.Value, m.N, m.Until)
	}
	fmt.Printf("\n")
	return nil
}

func longestMetricName(dump *reporter.FullDumpReturn) int {
	ret := 0

	for _, m := range dump.Averages {
		if ret < len(m.Name) {
			ret = len(m.Name)
		}
	}
	for _, m := range dump.AveragesPerDuration {
		if ret < len(m.Name) {
			ret = len(m.Name)
		}
	}
	for _, m := range dump.Counts {
		if ret < len(m.Name) {
			ret = len(m.Name)
		}
	}
	for _, m := range dump.CountsPerDuration {
		if ret < len(m.Name) {
			ret = len(m.Name)
		}
	}
	for _, m := range dump.Sums {
		if ret < len(m.Name) {
			ret = len(m.Name)
		}
	}
	for _, m := range dump.SumsPerDuration {
		if ret < len(m.Name) {
			ret = len(m.Name)
		}
	}

	return ret
}
