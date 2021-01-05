package sql

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

var (
	//ErrNoPtr ..
	ErrNoPtr = errors.New("noptr")
	//ErrNoResult ..
	ErrNoResult = errors.New("noResult")
)

//Model ..
type Model struct {
	DB *DB
	Tx *Tx
}

//NewModel ..
func NewModel(db *DB) (md *Model) {
	md = &Model{
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

//Select ..
func (m *Model) Select(c context.Context, dest interface{}, query string, args ...interface{}) (err error) {
	var rs *Rows
	if m.Tx != nil {
		rs, err = m.Tx.Query(query, args...)
	} else {
		rs, err = m.DB.Query(c, query, args...)
	}

	if nil != err {
		return
	}

	rt := reflect.ValueOf(dest)
	kind := rt.Kind()
	//desc should be ptr
	if kind != reflect.Ptr {
		err = ErrNoPtr
		return
	}

	a := rt.Elem()
	// if reflect.Struct != a.Kind() {
	// 	err = errors.New("the type of dest is not allowed")
	// 	return
	// }

	// convert the query result to the list of map
	columns, _ := rs.Columns()
	columnLength := len(columns)
	cache := make([]interface{}, columnLength)
	for index := range cache {
		var a interface{}
		cache[index] = &a
	}

	var list []map[string]interface{}
	for rs.Next() {
		_ = rs.Scan(cache...)

		item := make(map[string]interface{})
		for i, data := range cache {
			item[columns[i]] = *data.(*interface{}) //get the real kind
		}
		list = append(list, item)
	}

	if len(list) <= 0 {
		err = ErrNoResult
		return
	}

	// d := list[0]
	// for k, v := range d {
	// 	t := fmt.Sprintf("%T", v)
	// 	fmt.Println("key:", k, "v:", v, "type:", t)
	// }

	if reflect.Struct == a.Kind() {
		vType := a.Type()
		fieldMap := parseField(vType)
		fmt.Println(fieldMap)
		convertStruct(a, vType, list[0], fieldMap)
	} else if reflect.Slice == a.Kind() {
		vType := a.Type().Elem()
		fieldMap := parseField(vType)
		for _, data := range list {
			v := reflect.New(vType).Elem()
			convertStruct(v, vType, data, fieldMap)
			a = reflect.Append(a, v)
			// fmt.Println(a)
		}

		// rt := reflect.ValueOf(dest)
		rt.Elem().Set(a)
	}

	return
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
	if m.Tx != nil {
		fmt.Println("444444", m.Tx)
		err = m.Tx.Commit()
	}
	return
}
