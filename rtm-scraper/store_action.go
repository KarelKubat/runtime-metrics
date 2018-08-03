package main

import (
	"database/sql"

	"github.com/KarelKubat/runtime-metrics/reporter"
)

// StoreAction handles storing full metrics dumps.
type StoreAction struct {
	DB *sql.DB
}

// HandleFullDump implements ActionHandler to store a full metrics dump.
func (d *StoreAction) HandleFullDump(dump *reporter.FullDumpReturn) {
	panic("not implemented")
}
