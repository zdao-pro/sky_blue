package sql

import (
	"context"
	"database/sql"
	"fmt"
)

//Model ..
type Model struct {
	DB *DB
	Tx *Tx
}

//NewModel ..
func NewModel(db *DB) (md Model) {
	md = Model{
		DB: db,
	}
	return
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (m *Model) Exec(c context.Context, query string, args ...interface{}) (res sql.Result, err error) {
	if m.Tx != nil {
		fmt.Println("33333")
		return m.Tx.Exec(query, args...)
	}
	return m.DB.Exec(c, query, args...)
}

// Query executes a query that returns rows, typically a SELECT. The args are
// for any placeholder parameters in the query.
func (m *Model) Query(c context.Context, query string, args ...interface{}) (*Rows, error) {
	if m.Tx != nil {
		return m.Tx.Query(query, args...)
	}
	return m.DB.Query(c, query, args...)
}

// QueryRow executes a query that is expected to return at most one row.
// QueryRow always returns a non-nil value. Errors are deferred until Row's
// Scan method is called.
func (m *Model) QueryRow(c context.Context, query string, args ...interface{}) *Row {
	if m.Tx != nil {
		return m.Tx.QueryRow(query, args...)
	}
	return m.DB.QueryRow(c, query, args...)
}

// Begin starts a transaction. The isolation level is dependent on the driver.
func (m *Model) Begin(c context.Context) (err error) {
	m.Tx, err = m.DB.Begin(c)
	return err
}

// Rollback aborts the transaction.
func (m *Model) Rollback() (err error) {
	err = m.Tx.Rollback()
	return
}

// Commit commits the transaction.
func (m *Model) Commit() (err error) {
	// fmt.Println(m.Tx)
	err = m.Tx.Commit()
	return
}
