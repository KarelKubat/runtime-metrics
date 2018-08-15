package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/KarelKubat/runtime-metrics/reporter"
	log "github.com/sirupsen/logrus"
)

// CompressPolicy defines after which duration, datapoints should be compressed or dropped.
type CompressPolicy struct {
	to1Min    time.Duration
	to5Mins   time.Duration
	to15Mins  time.Duration
	to30Mins  time.Duration
	to1Hour   time.Duration
	dropAfter time.Duration
}

// NewCompressPolicy returns an initialized CompressPolicy.
func NewCompressPolicy(to1Min, to5Mins, to15Mins, to30Mins, to1Hour,
	dropAfter time.Duration) *CompressPolicy {

	return &CompressPolicy{
		to1Min:    to1Min,
		to5Mins:   to5Mins,
		to15Mins:  to15Mins,
		to30Mins:  to30Mins,
		to1Hour:   to1Hour,
		dropAfter: dropAfter,
	}
}

// StoreAction handles storing full metrics dumps.
type StoreAction struct {
	db             *sql.DB
	driver         *string
	compressPolicy *CompressPolicy
	nameCache      map[string]int
}

// NewStoreAction returns an initialized StoreAction.
func NewStoreAction(db *sql.DB, driver *string,
	compressPolicy *CompressPolicy) (*StoreAction, error) {

	ret := &StoreAction{
		db:             db,
		driver:         driver,
		compressPolicy: compressPolicy,
		nameCache:      map[string]int{},
	}
	if err := ret.createTables(); err != nil {
		return nil, err
	}
	go ret.compress()
	return ret, nil
}

// logSQL logs a SQL query and bind variables.
func logSQL(qry string, bindVariables ...interface{}) {
	loggableSQL := qry
	for _, sub := range []string{"\n", "\t", "  "} {
		for strings.Contains(loggableSQL, sub) {
			loggableSQL = strings.Replace(loggableSQL, sub, " ", -1)
		}
	}
	loggableVars := ""
	for i, b := range bindVariables {
		if loggableVars != "" {
			loggableVars += ", "
		}
		loggableVars += fmt.Sprintf("$%d=%v", i+1, b)
	}
	log.WithFields(log.Fields{
		"sql":  loggableSQL,
		"vars": loggableVars,
	}).Debug("db query")
}

// queryRow wraps sql.QueryRow for logging and error recovery.
func (d *StoreAction) queryRow(qry string, bindVariables ...interface{}) *sql.Row {
	logSQL(qry, bindVariables...)
	return d.db.QueryRow(qry, bindVariables...)
}

// exec wraps sql.Exec for logging and error recovery. Exec's Result is not returned.
func (d *StoreAction) exec(qry string, bindVariables ...interface{}) error {
	var err error

	logSQL(qry, bindVariables...)
	for i := 0; i < 5; i++ {
		_, err = d.db.Exec(qry, bindVariables...)
		if err != nil && strings.Index(err.Error(), "database is locked") != -1 {
			time.Sleep(time.Millisecond * 100)
			continue
		}
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Warn("db error")
		}
		return err
	}
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("db error")
	}
	return err
}

// query wraps sql.Query for logging and error recovery.
func (d *StoreAction) query(qry string, bindVariables ...interface{}) (*sql.Rows, error) {
	var err error
	var rows *sql.Rows

	logSQL(qry, bindVariables...)
	for i := 0; i < 5; i++ {
		rows, err = d.db.Query(qry, bindVariables...)
		if err != nil && strings.Index(err.Error(), "database is locked") != -1 {
			time.Sleep(time.Millisecond * 100)
			continue
		}
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Warn("db error")
		}
		return rows, err
	}
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("db error")
	}
	return rows, err
}

// createTables optionally creates tables.
func (d *StoreAction) createTables() error {
	var err error
	var serialSQL string

	if serialSQL, err = d.serialName(); err != nil {
		return err
	}

	// Metric names
	if err = d.exec(fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS metric_name (
		   id %s,
		   name TEXT
		 )`, serialSQL)); err != nil {
		return fmt.Errorf("failed to initialize: %v", err)
	}

	// Averages and -per duration
	if err = d.exec(fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS average (
		   id %s,
		   name_id INTEGER NOT NULL REFERENCES metric_name(id),
		   value REAL NOT NULL,
		   n REAL NOT NULL,
		   timestamp TEXT NOT NULL
		 )`, serialSQL)); err != nil {
		return fmt.Errorf("failed to initialize: %v", err)
	}
	if err = d.exec(fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS average_per_duration (
		   id %s,
		   name_id INTEGER NOT NULL REFERENCES metric_name(id),
		   value REAL NOT NULL,
		   n REAL NOT NULL,
		   until TEXT NOT NULL,
		   timestamp TEXT NOT NULL
		 )`, serialSQL)); err != nil {
		return fmt.Errorf("failed to initialize: %v", err)
	}

	// Counts and -per duration
	if err = d.exec(fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS count (
		   id %s,
		   name_id INTEGER NOT NULL REFERENCES metric_name(id),
		   value INTEGER NOT NULL,
		   timestamp TEXT NOT NULL
		 )`, serialSQL)); err != nil {
		return fmt.Errorf("failed to initialize: %v", err)
	}
	if err = d.exec(fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS count_per_duration (
		   id %s,
		   name_id INTEGER NOT NULL REFERENCES metric_name(id),
		   value INTEGER NOT NULL,
		   until TEXT NOT NULL,
		   timestamp TEXT NOT NULL
		 )`, serialSQL)); err != nil {
		return fmt.Errorf("failed to initialize: %v", err)
	}

	// Sums and -per duration
	if err = d.exec(fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS sum (
		   id %s,
		   name_id INTEGER NOT NULL REFERENCES metric_name(id),
		   value REAL NOT NULL,
		   n REAL NOT NULL,
		   timestamp TEXT NOT NULL
		 )`, serialSQL)); err != nil {
		return fmt.Errorf("failed to initialize: %v", err)
	}
	if err = d.exec(fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS sum_per_duration (
		   id %s,
		   name_id INTEGER NOT NULL REFERENCES metric_name(id),
		   value REAL NOT NULL,
		   n REAL NOT NULL,
		   until TEXT NOT NULL,
		   timestamp TEXT
		 )`, serialSQL)); err != nil {
		return fmt.Errorf("failed to initialize: %v", err)
	}

	return nil
}

// HandleFullDump implements ActionHandler to store a full metrics dump.
func (d *StoreAction) HandleFullDump(dump *reporter.FullDumpReturn) error {
	var nameID int
	var err error

	// Averages and -per duration
	for _, m := range dump.Averages {
		if nameID, err = d.upsertName(m.Name); err != nil {
			return err
		}
		if err = d.exec(
			`INSERT INTO average (name_id, value, n, timestamp)
			 VALUES (?,?,?,?)`, nameID, m.Value, m.N, time.Now()); err != nil {
			return err
		}
	}
	for _, m := range dump.AveragesPerDuration {
		if nameID, err = d.upsertName(m.Name); err != nil {
			return err
		}
		if err = d.exec(
			`INSERT INTO average_per_duration (name_id, value, n, until, timestamp)
			 VALUES (?,?,?,?,?)`, nameID, m.Value, m.N, m.Until, time.Now()); err != nil {
			return err
		}
	}

	// Counts and -per duration
	for _, m := range dump.Counts {
		if nameID, err = d.upsertName(m.Name); err != nil {
			return err
		}
		if err = d.exec(
			`INSERT INTO count (name_id, value, timestamp)
			 VALUES (?,?,?)`, nameID, m.Value, time.Now()); err != nil {
			return err
		}
	}
	for _, m := range dump.CountsPerDuration {
		if nameID, err = d.upsertName(m.Name); err != nil {
			return err
		}
		if err = d.exec(
			`INSERT INTO count_per_duration (name_id, value, until, timestamp)
			 VALUES (?,?,?,?)`, nameID, m.Value, m.Until, time.Now()); err != nil {
			return err
		}
	}

	// Sums and -per duration
	for _, m := range dump.Sums {
		if nameID, err = d.upsertName(m.Name); err != nil {
			return err
		}
		if err = d.exec(
			`INSERT INTO sum (name_id, value, n, timestamp)
			 VALUES (?,?,?,?)`, nameID, m.Value, m.N, time.Now()); err != nil {
			return err
		}
	}
	for _, m := range dump.SumsPerDuration {
		if nameID, err = d.upsertName(m.Name); err != nil {
			return err
		}
		if err = d.exec(
			`INSERT INTO sum_per_duration (name_id, value, n, until, timestamp)
			 VALUES (?,?,?,?,?)`, nameID, m.Value, m.N, m.Until, time.Now()); err != nil {
			return err
		}
	}

	return nil
}

// upsertName makes sure that a metric name is stored.
func (d *StoreAction) upsertName(n string) (int, error) {
	var ret int
	var err error

	// If cached, return the ID immediately.
	ret, ok := d.nameCache[n]
	if ok {
		return ret, nil
	}

	// If already known, cache and return.
	r := d.queryRow("SELECT id FROM metric_name WHERE name=?", n)
	if err = r.Scan(&ret); err == nil {
		d.nameCache[n] = ret
		return ret, nil
	}
	if strings.Index(err.Error(), "sql: no rows") == -1 {
		return 0, err
	}

	// Need to add to table metric_name.
	// Unfortunately, sqlite3 doesn't respect LastInsertId() of the return
	// value. We have to insert and then requery.
	if err = d.exec("INSERT INTO metric_name (name) VALUES (?)", n); err != nil {
		return 0, err
	}
	r = d.queryRow("SELECT id FROM metric_name WHERE name=?", n)
	if err = r.Scan(&ret); err != nil {
		return 0, err
	}
	d.nameCache[n] = ret
	return ret, nil
}

// serialName is the driver-dependent SQL DDL to define an auto-incrementing ID.
func (d *StoreAction) serialName() (string, error) {
	switch *d.driver {
	case "sqlite3":
		return "INTEGER PRIMARY KEY AUTOINCREMENT", nil
	case "postgres":
		return "SERIAL", nil
	}
	return "", fmt.Errorf("driver %v not implemented", d.driver)
}

// compress takes care of compressing the datapoints.
func (d *StoreAction) compress() {
	time.Sleep(time.Minute)
	d.dropAfter()
}

// dropAfter drops stale entries, i.e., older than the policies dropAfter duration.
func (d *StoreAction) dropAfter() {
	var cutoff time.Time

	// Drop entries that we no longer need
	cutoff = time.Now().Add(-d.compressPolicy.dropAfter)
	for _, table := range []string{
		"average",
		"average_per_duration",
		"count",
		"count_per_duration",
		"sum",
		"sum_per_duration",
	} {
		rows, err := d.query(fmt.Sprintf(
			"SELECT id FROM %s WHERE timestamp < ?", table), cutoff)
		if err != nil {
			continue
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			if err := rows.Scan(&id); err != nil {
				log.WithFields(log.Fields{"err": err}).Warn("db error")
				continue
			}
			d.exec(fmt.Sprintf("DELETE FROM %s WHERE id=?", table), cutoff)
		}
	}
}
