package sql

import (
	"fmt"
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
	fmt.Println("src:  ", src, " type:", f.FieldType)
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
	// switch v.Kind(){
	// case reflect.Int:
	// 	return setIntField(val, 0, value)
	// case reflect.Int8:
	// 	return setIntField(val, 8, value)
	// case reflect.Int16:
	// 	return setIntField(val, 16, value)
	// case reflect.Int32:
	// 	return setIntField(val, 32, value)
	// case reflect.Int64:
	// 	switch value.Interface().(type) {
	// 	case time.Duration:
	// 		return setTimeDuration(val, value, field)
	// 	}
	// 	return setIntField(val, 64, value)
	// case reflect.Uint:
	// 	return setUintField(val, 0, value)
	// case reflect.Uint8:
	// 	return setUintField(val, 8, value)
	// case reflect.Uint16:
	// 	return setUintField(val, 16, value)
	// case reflect.Uint32:
	// 	return setUintField(val, 32, value)
	// case reflect.Uint64:
	// 	return setUintField(val, 64, value)
	// case reflect.Bool:
	// 	return setBoolField(val, value)
	// case reflect.Float32:
	// 	return setFloatField(val, 32, value)
	// case reflect.Float64:
	// 	return setFloatField(val, 64, value)
	// case reflect.String:
	// 	value.SetString(val)
	// case reflect.Struct:
	// case reflect.Map:
	// 	return json.Unmarshal(bytesconv.StringToBytes(val), value.Addr().Interface())
	// default:
	// 	return errUnknownType
	// }
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

// func setUintField(val interface{}, bitSize int, field reflect.Value) error {
// 	v, ok := val.(uint)
// 	if ok{
// 		field.SetUint(v)
// 	}

// 	s := val.(string)
// 	uintVal, err := strconv.ParseUint(s, 10, bitSize)
// 	if err == nil {
// 		field.SetUint(uintVal)
// 	}
// 	return err
// }

// func setBoolField(val string, field reflect.Value) error {
// 	v, ok := val.(bool)
// 	if ok{
// 		field.SetBool(v)
// 	}

// 	s := val.(string)
// 	boolVal, err := strconv.ParseBool(s)
// 	if err == nil {
// 		field.SetBool(boolVal)
// 	}
// 	return err
// }

// func setFloatField(val string, bitSize int, field reflect.Value) error {
// 	v, ok := val.(uint)
// 	if ok{
// 		field.SetUint(v)
// 	}

// 	s := val.(string)
// 	floatVal, err := strconv.ParseFloat(val, bitSize)
// 	if err == nil {
// 		field.SetFloat(floatVal)
// 	}
// 	return err
// }

// func setTimeDuration(val interface{}, value reflect.Value, field reflect.StructField) error {
// 	d, err := time.ParseDuration(val)
// 	if err != nil {
// 		return err
// 	}
// 	value.Set(reflect.ValueOf(d))
// 	return nil
// }
