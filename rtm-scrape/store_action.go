package main

import (
	"database/sql"
	"fmt"

	"github.com/KarelKubat/runtime-metrics/reporter"
)

// StoreAction handles storing full metrics dumps.
type StoreAction struct {
	DB     *sql.DB
	Driver *string
}

// Init optionally creates tables.
func (d *StoreAction) Init() error {
	_, err := d.DB.Exec(
		fmt.Sprintf("CREATE TABLE IF NOT EXISTS metric_name (id %s, name text)", d.serialName()))
	if err != nil {
		return fmt.Errorf("failed to initialize: %v", err)
	}

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

func (d *StoreAction) upsertName(n string) (int, error) {
	var ret int

	r := d.DB.QueryRow("SELECT id FROM metric_name WHERE name=?", n)
	if r.Next() {
		if err := r.Scan(&ret); err != nil {
			return 0, err
		}
		return ret, nil
	}

	r := db.DB.QueryRow("INSERT INTO metric_name (name) VALUES (?) RETURNING id", n)
	if !r.Next() {
		return 0, fmt.Errorf("failed to insert metric name %q (no ID)", n)
	}
	if err := r.Scan(&ret); err != nil {
		return 0, err
	}
	return ret, nil
}

func (d *StoreAction) serialName() string {
	switch *d.Driver {
	case "sqlite3":
		return "AUTO_INCREMENT"
	case "postgres":
		return "SERIAL"
	}
	return ""
}
