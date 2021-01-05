package sql

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"
)

func convertStruct(a reflect.Value, vType reflect.Type, m map[string]interface{}, fieldMap map[string]Field) (err error) {
	for i := 0; i < a.NumField(); i++ {
		v := a.Field(i)
		t := vType.Field(i)
		name := t.Name
		f, ok := fieldMap[name]
		if !ok {
			continue
		}
		vMap, ok := m[f.Name]
		if !ok {
			continue
		}
		convertFiled(f, vMap, v)
	}
	return
}

func convertFiled(f Field, src interface{}, value reflect.Value) (err error) {
	// fmt.Println("src:  ", src, " type:", f.FieldType)
	switch src.(type) {
	case string:
		switch f.DataType {
		case String:
			value.SetString(src.(string))
		case Bytes:
			value.SetBytes([]byte(src.(string)))
		case Int:
			setIntField(src.(string), 64, value)
		case Uint:
			setUintField(src.(string), 64, value)
		case Float:
			setFloatField(src.(string), 64, value)
		case Struct:
			setStructField([]byte(src.(string)), value)
		}
	case *string:
		switch f.DataType {
		case String:
			value.SetString(*src.(*string))
		case Bytes:
			value.SetBytes([]byte(*src.(*string)))
		case Int:
			setIntField(*src.(*string), 64, value)
		case Uint:
			setUintField(*src.(*string), 64, value)
		case Float:
			setFloatField(*src.(*string), 64, value)
		case Struct:
			setStructField([]byte(*src.(*string)), value)
		}
	case []byte:
		switch f.DataType {
		case String:
			value.SetString(string(src.([]byte)))
		case Bytes:
			value.SetBytes(src.([]byte))
		case Int:
			setIntField(string(src.([]byte)), 64, value)
		case Uint:
			setUintField(string(src.([]byte)), 64, value)
		case Float:
			setFloatField(string(src.([]byte)), 64, value)
		case Struct:
			setStructField(src.([]byte), value)
		}
	case time.Time:
		switch f.DataType {
		case Time:
			value.Set(reflect.ValueOf(src))
		}
	case int, int8, int16, int32, int64:
		switch f.DataType {
		case Int:
			value.SetInt(src.(int64))
		case Time:
			setTimeField(src, value, f.TimeFormat)
		}
	case uint, uint8, uint16, uint32, uint64:
		switch f.DataType {
		case Uint:
			value.SetUint(src.(uint64))
		case Time:
			setTimeField(src, value, f.TimeFormat)
		}
	case float32, float64:
		switch f.DataType {
		case Float:
			value.SetFloat(src.(float64))
		}
	case bool:
		switch f.DataType {
		case Bool:
			value.SetBool(src.(bool))
		}
	}
	return
}

func setStructField(body []byte, field reflect.Value) {
	// fmt.Println(string(body))
	json.Unmarshal(body, field.Addr().Interface())
}
func setTimeField(val interface{}, field reflect.Value, timeFormat string) {
	var v int64
	i, ok := val.(int64)
	if ok {
		v = i
	}
	if v != 0 {
		if "unix" == timeFormat {
			t := time.Unix(v, v%int64(time.Duration(1)))
			field.Set(reflect.ValueOf(t))
		}
		if "unixnano" == timeFormat {
			t := time.Unix(v/int64(1000000), v%1000000)
			field.Set(reflect.ValueOf(t))
		}
		if "unixmilli" == timeFormat {
			t := time.Unix(v/1000, (v%1000)*int64(time.Millisecond))
			field.Set(reflect.ValueOf(t))
		}
	}
}

func setUintField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	uintVal, err := strconv.ParseUint(val, 10, bitSize)
	if err == nil {
		field.SetUint(uintVal)
	}
	return err
}

func setIntField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0"
	}
	intVal, err := strconv.ParseInt(val, 10, bitSize)
	if err == nil {
		field.SetInt(intVal)
	}
	return err
}

func setFloatField(val string, bitSize int, field reflect.Value) error {
	if val == "" {
		val = "0.0"
	}
	floatVal, err := strconv.ParseFloat(val, bitSize)
	if err == nil {
		field.SetFloat(floatVal)
	}
	return err
}
