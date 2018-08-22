package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/KarelKubat/runtime-metrics/reporter"
	log "github.com/sirupsen/logrus"
)

// StoreAction handles storing full metrics dumps.
type StoreAction struct {
	db              *sql.DB
	compressionInfo []CompressionInfo
	cleanerInterval time.Duration
	nameCache       map[string]int
}

// NewStoreAction returns an initialized StoreAction.
func NewStoreAction(db *sql.DB, compressionInfo []CompressionInfo,
	cleanerInterval time.Duration) (*StoreAction, error) {

	ret := &StoreAction{
		db:              db,
		compressionInfo: compressionInfo,
		cleanerInterval: cleanerInterval,
		nameCache:       map[string]int{},
	}
	if err := ret.createTables(); err != nil {
		return nil, err
	}
	go ret.runCleanup()
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
	if loggableVars != "" {
		log.WithFields(log.Fields{
			"sql":  loggableSQL,
			"vars": loggableVars,
		}).Debug("db query")
	} else {
		log.WithFields(log.Fields{
			"sql": loggableSQL,
		}).Debug("db query")
	}
}

// queryRow wraps sql.QueryRow for logging and error recovery.
func (d *StoreAction) queryRow(qry string, bindVariables ...interface{}) *sql.Row {
	logSQL(qry, bindVariables...)
	return d.db.QueryRow(qry, bindVariables...)
}

// exec wraps sql.Exec for logging and error recovery. The Result is discarded.
func (d *StoreAction) exec(qry string, bindVariables ...interface{}) error {
	var err error

	logSQL(qry, bindVariables...)
	if _, err = d.db.Exec(qry, bindVariables...); err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("db error")
	}
	return err
}

// query wraps sql.Query for logging and error recovery.
func (d *StoreAction) query(qry string, bindVariables ...interface{}) (*sql.Rows, error) {
	var err error
	var rows *sql.Rows

	logSQL(qry, bindVariables...)
	rows, err = d.db.Query(qry, bindVariables...)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("db error")
	}
	return rows, err
}

// createTables optionally creates tables.
func (d *StoreAction) createTables() error {
	var err error

	// Metric names
	if err = d.exec(fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS metric_name (
		   id SERIAL PRIMARY KEY,
		   name TEXT UNIQUE
		 )`)); err != nil {
		return fmt.Errorf("failed to initialize: %v", err)
	}

	// Tables for scraped data. We have UNIQUE constraints for all timestamps, which are
	// the times that we received the data. Furthermore we have UNIQUE constraints on
	// values and `until` of all _per_duration data to avoid duplicates. The scraper may
	// run in parallel, or may sample with a higher frequency than server-metrics change.

	// Averages and -per duration
	if err = d.exec(fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS average (
		   id SERIAL PRIMARY KEY,
		   name_id INTEGER NOT NULL REFERENCES metric_name(id),
		   value REAL NOT NULL,
		   n REAL NOT NULL,
		   timestamp TEXT NOT NULL UNIQUE
		 )`)); err != nil {
		return fmt.Errorf("failed to initialize: %v", err)
	}
	if err = d.exec(fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS average_per_duration (
		   id SERIAL PRIMARY KEY,
		   name_id INTEGER NOT NULL REFERENCES metric_name(id),
		   value REAL NOT NULL,
		   n REAL NOT NULL,
		   until TEXT NOT NULL,
		   timestamp TEXT NOT NULL UNIQUE,
		   UNIQUE (name_id, value, n, until)
		 )`)); err != nil {
		return fmt.Errorf("failed to initialize: %v", err)
	}

	// Counts and -per duration
	if err = d.exec(fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS count (
		   id SERIAL PRIMARY KEY,
		   name_id INTEGER NOT NULL REFERENCES metric_name(id),
		   value INTEGER NOT NULL,
		   timestamp TEXT NOT NULL UNIQUE
		 )`)); err != nil {
		return fmt.Errorf("failed to initialize: %v", err)
	}
	if err = d.exec(fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS count_per_duration (
		   id SERIAL PRIMARY KEY,
		   name_id INTEGER NOT NULL REFERENCES metric_name(id),
		   value INTEGER NOT NULL,
		   until TEXT NOT NULL,
		   timestamp TEXT NOT NULL,
		   UNIQUE (name_id, value, until)
		 )`)); err != nil {
		return fmt.Errorf("failed to initialize: %v", err)
	}

	// Sums and -per duration
	if err = d.exec(fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS sum (
		   id SERIAL PRIMARY KEY,
		   name_id INTEGER NOT NULL REFERENCES metric_name(id),
		   value REAL NOT NULL,
		   n REAL NOT NULL,
		   timestamp TEXT NOT NULL UNIQUE
		 )`)); err != nil {
		return fmt.Errorf("failed to initialize: %v", err)
	}
	if err = d.exec(fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS sum_per_duration (
		   id SERIAL PRIMARY KEY,
		   name_id INTEGER NOT NULL REFERENCES metric_name(id),
		   value REAL NOT NULL,
		   n REAL NOT NULL,
		   until TEXT NOT NULL,
		   timestamp TEXT UNIQUE,
		   UNIQUE (name_id, value, n, until)
		 )`)); err != nil {
		return fmt.Errorf("failed to initialize: %v", err)
	}

	return nil
}

// HandleFullDump implements ActionHandler to store a full metrics dump.
func (d *StoreAction) HandleFullDump(dump *reporter.FullDumpReturn) error {
	var nameID int
	var err error

	// All data are inserted with `ON CONFLICT DO NOTHING` because the scraper may sample at
	// a higher frequency than server-metrics are changing.

	// Averages and -per duration
	for _, m := range dump.Averages {
		if nameID, err = d.upsertName(m.Name); err != nil {
			return err
		}
		if err = d.exec(
			`INSERT INTO average (name_id, value, n, timestamp)
			 VALUES ($1,$2,$3,$4)
			 ON CONFLICT DO NOTHING`, nameID, m.Value, m.N, time.Now().UTC()); err != nil {
			return err
		}
	}
	for _, m := range dump.AveragesPerDuration {
		if nameID, err = d.upsertName(m.Name); err != nil {
			return err
		}
		if err = d.exec(
			`INSERT INTO average_per_duration (name_id, value, n, until, timestamp)
			 VALUES ($1,$2,$3,$4,$5)
			 ON CONFLICT DO NOTHING`, nameID, m.Value, m.N, m.Until, time.Now().UTC()); err != nil {
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
			 VALUES ($1,$2,$3)
			 ON CONFLICT DO NOTHING`, nameID, m.Value, time.Now().UTC()); err != nil {
			return err
		}
	}
	for _, m := range dump.CountsPerDuration {
		if nameID, err = d.upsertName(m.Name); err != nil {
			return err
		}
		if err = d.exec(
			`INSERT INTO count_per_duration (name_id, value, until, timestamp)
			 VALUES ($1,$2,$3,$4)
			 ON CONFLICT DO NOTHING`, nameID, m.Value, m.Until, time.Now().UTC()); err != nil {
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
			 VALUES ($1,$2,$3,$4)
			 ON CONFLICT DO NOTHING`, nameID, m.Value, m.N, time.Now().UTC()); err != nil {
			return err
		}
	}
	for _, m := range dump.SumsPerDuration {
		if nameID, err = d.upsertName(m.Name); err != nil {
			return err
		}
		if err = d.exec(
			`INSERT INTO sum_per_duration (name_id, value, n, until, timestamp)
			 VALUES ($1,$2,$3,$4,$5)
			 ON CONFLICT DO NOTHING`, nameID, m.Value, m.N, m.Until, time.Now().UTC()); err != nil {
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
	if err = d.queryRow("SELECT id FROM metric_name WHERE name=$1", n).Scan(&ret); err == nil {
		d.nameCache[n] = ret
		return ret, nil
	}
	if strings.Index(err.Error(), "sql: no rows") == -1 {
		return 0, err
	}

	// Need to add to table metric_name. Unfortunately, LastInsertId() may not work well so we
	// have to requery.
	if err = d.exec("INSERT INTO metric_name (name) VALUES ($1)", n); err != nil {
		return 0, err
	}
	if err = d.queryRow("SELECT id FROM metric_name WHERE name=$1", n).Scan(&ret); err != nil {
		return 0, err
	}
	d.nameCache[n] = ret
	return ret, nil
}

// compress takes care of compressing the datapoints.
func (d *StoreAction) runCleanup() {
	for {
		time.Sleep(d.cleanerInterval)
		for i := 0; i < len(d.compressionInfo); i++ {
			// If the period isn't stated, then the intention is to drop.
			// Otherwise, compress entries between after and until into period-intervals.
			if d.compressionInfo[i].period == 0 {
				d.dropAfter(d.compressionInfo[i].after)
			} else {
				d.compress(compressionInfo[i].after, compressionInfo[i].until,
					compressionInfo[i].period)
			}
		}
	}
}

// compress recalculates datapoints between its first two duration arguments into
// periods given by the third argument.
func (d *StoreAction) compress(after, until, period time.Duration) {
	afterTime := time.Now().UTC().Add(-after)
	untilTime := time.Now().UTC().Add(-until)
	log.WithFields(log.Fields{
		"after":  afterTime,
		"until":  untilTime,
		"period": period,
	}).Debug("compress-to cleanup")

}

// dropAfter drops stale entries, i.e., older than the policies dropAfter duration.
func (d *StoreAction) dropAfter(after time.Duration) {
	var cutoff time.Time

	// Drop entries that we no longer need
	cutoff = time.Now().UTC().Add(-after)
	log.WithFields(log.Fields{"cutoff": cutoff}).Debug("drop-after cleanup")

	for _, table := range []string{
		"average",
		"average_per_duration",
		"count",
		"count_per_duration",
		"sum",
		"sum_per_duration",
	} {
		d.exec(fmt.Sprintf("DELETE FROM %s WHERE timestamp < $1", table), cutoff)
	}
}
