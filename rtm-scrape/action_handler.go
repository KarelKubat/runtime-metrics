package main

import (
	"github.com/KarelKubat/runtime-metrics/reporter"
)

// ActionHandler is the generic interface for actions.
type ActionHandler interface {
	HandleFullDump(*reporter.FullDumpReturn) error
}
