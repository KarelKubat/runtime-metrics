package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/KarelKubat/runtime-metrics/reporter"
)

// StoreAction handles storing full metrics dumps.
type StoreAction struct {
	DB        *sql.DB
	Driver    *string
	nameCache map[string]int
}

// Init optionally creates tables.
func (d *StoreAction) Init() error {
	serialSQL, err := d.serialName()
	if err != nil {
		return err
	}
	_, err = d.DB.Exec(
		fmt.Sprintf("CREATE TABLE IF NOT EXISTS metric_name (id %s, name text)", serialSQL))
	if err != nil {
		return fmt.Errorf("failed to initialize: %v", err)
	}
	d.nameCache = make(map[string]int)
	return nil
}

// HandleFullDump implements ActionHandler to store a full metrics dump.
func (d *StoreAction) HandleFullDump(dump *reporter.FullDumpReturn) error {
	for _, m := range dump.Averages {
		nameID, err := d.upsertName(m.Name)
		if err != nil {
			return err
		}
		fmt.Printf("upserted name %v as %v\n", m.Name, nameID)
	}

	for _, m := range dump.AveragesPerDuration {
		nameID, err := d.upsertName(m.Name)
		if err != nil {
			return err
		}
		fmt.Printf("upserted name %v as %v\n", m.Name, nameID)
	}

	return nil
}

// upsertName makes sure that a metric name is stored.
func (d *StoreAction) upsertName(n string) (int, error) {
	var ret int

	ret, ok := d.nameCache[n]
	if ok {
		return ret, nil
	}

	r := d.DB.QueryRow("SELECT id FROM metric_name WHERE name=?", n)
	err := r.Scan(&ret)
	if err == nil {
		d.nameCache[n] = ret
		return ret, nil
	}

	if strings.Index(err.Error(), "sql: no rows") == -1 {
		return 0, err
	}

	res, err := d.DB.Exec("INSERT INTO metric_name (name) VALUES (?)", n)
	if err != nil {
		return 0, err
	}
	ret64, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	d.nameCache[n] = ret
	return int(ret64), nil
}

// serialName is the driver-dependent SQL DDL to define an auto-incrementing ID.
func (d *StoreAction) serialName() (string, error) {
	switch *d.Driver {
	case "sqlite3":
		return "INTEGER PRIMARY KEY AUTOINCREMENT", nil
	case "postgres":
		return "SERIAL", nil
	}
	return "", fmt.Errorf("driver %v not implemented", d.Driver)
}
