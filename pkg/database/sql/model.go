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

	v := reflect.ValueOf(dest)
	kind := v.Kind()
	//desc should be ptr
	if kind != reflect.Ptr {
		err = ErrNoPtr
		return
	}

	a := v.Elem()
	if reflect.Struct != a.Kind() {
		err = errors.New("the type of dest is not allowed")
		return
	}

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

	if reflect.Struct == a.Kind() {
		// convertFiled(a, list[0])
	}
	fmt.Println(list)
	return
}

type setOption struct {
	name       string
	timeFormat string //
}

func convertFiled(v reflect.Value, val interface{}, set setOption) {
	switch v.Kind(){
	case reflect.Int:
		return setIntField(val, 0, value)
	case reflect.Int8:
		return setIntField(val, 8, value)
	case reflect.Int16:
		return setIntField(val, 16, value)
	case reflect.Int32:
		return setIntField(val, 32, value)
	case reflect.Int64:
		switch value.Interface().(type) {
		case time.Duration:
			return setTimeDuration(val, value, field)
		}
		return setIntField(val, 64, value)
	case reflect.Uint:
		return setUintField(val, 0, value)
	case reflect.Uint8:
		return setUintField(val, 8, value)
	case reflect.Uint16:
		return setUintField(val, 16, value)
	case reflect.Uint32:
		return setUintField(val, 32, value)
	case reflect.Uint64:
		return setUintField(val, 64, value)
	case reflect.Bool:
		return setBoolField(val, value)
	case reflect.Float32:
		return setFloatField(val, 32, value)
	case reflect.Float64:
		return setFloatField(val, 64, value)
	case reflect.String:
		value.SetString(val)
	case reflect.Struct:
	case reflect.Map:
		return json.Unmarshal(bytesconv.StringToBytes(val), value.Addr().Interface())
	default:
		return errUnknownType
	}
}

func setIntField(val interface{}, bitSize int, field reflect.Value) error {
	v, ok := val.(int)
	if ok{
		field.SetInt(v)
	}

	s := val.(string)
	intVal, err := strconv.ParseInt(s, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
	return err
}

func setUintField(val interface{}, bitSize int, field reflect.Value) error {
	v, ok := val.(uint)
	if ok{
		field.SetUint(v)
	}

	s := val.(string)
	uintVal, err := strconv.ParseUint(s, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}
	return err
}

func setBoolField(val string, field reflect.Value) error {
	v, ok := val.(bool)
	if ok{
		field.SetBool(v)
	}

	s := val.(string)
	boolVal, err := strconv.ParseBool(s)
	if err == nil {
		field.SetBool(boolVal)
	}
	return err
}

func setFloatField(val string, bitSize int, field reflect.Value) error {
	v, ok := val.(uint)
	if ok{
		field.SetUint(v)
	}

	s := val.(string)
	floatVal, err := strconv.ParseFloat(val, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
	return err
}

func getSetOption(v reflect.Value, m map[string]interface{}) map[string]setOption {
	tValue := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := tValue.Field(i)
		setOpt := setOption{}

		if k, ok := field.Tag.Lookup("orm"); ok {
			setOpt.name = k
		} else {
			setOpt.name = field.Name
		}
		m[field.Name] = setOpt
		// t := sf.Type
		// switch t. {

		// }
	}
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
