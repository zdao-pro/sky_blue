package sql

import (
	"reflect"
	"time"
)

//DataType ..
type DataType string

//TimeType ..
type TimeType int64

//TimeReflectType ..
var TimeReflectType = reflect.TypeOf(time.Time{})

const (
	//UnixSecond ..
	UnixSecond TimeType = 1
	//UnixMillisecond ..
	UnixMillisecond TimeType = 2
	//UnixNanosecond ..
	UnixNanosecond TimeType = 3
)

const (
	//Bool ..
	Bool DataType = "bool"
	//Int ..
	Int DataType = "int"
	//Uint ..
	Uint DataType = "uint"
	//Float ..
	Float DataType = "float"
	//String ..
	String DataType = "string"
	//Time ..
	Time DataType = "time"
	//Bytes ..
	Bytes DataType = "bytes"
	//Struct ..
	Struct DataType = "struct"
)

//Field ...
type Field struct {
	Name       string
	FieldType  reflect.Type
	TimeFormat string //
	DataType   DataType
}

func parseField(v reflect.Type) map[string]Field {
	fileldMap := make(map[string]Field)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if "Model" == field.Name {
			continue
		}

		f := Field{}

		if k, ok := field.Tag.Lookup("json"); ok {
			f.Name = k
		} else {
			f.Name = field.Name
		}
		f.FieldType = field.Type

		switch field.Type.Kind() {
		case reflect.Bool:
			f.DataType = Bool
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			f.DataType = Int
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			f.DataType = Uint
		case reflect.Float32, reflect.Float64:
			f.DataType = Float
		case reflect.String:
			f.DataType = String
		case reflect.Array, reflect.Slice:
			f.DataType = Bytes
		case reflect.Struct:
			if f.FieldType.String() == "time.Time" {
				f.DataType = Time
				timeFormat := field.Tag.Get("time_format")
				if timeFormat != "" {
					f.TimeFormat = timeFormat
				}
			} else {
				f.DataType = Struct
			}
			// if _, ok := v.Interface().(time.Time); ok {
			// 	f.DataType = Time
			// } else if v.Type().ConvertibleTo(TimeReflectType) {
			// 	f.DataType = Time
			// } else if v.Type().ConvertibleTo(reflect.TypeOf(&time.Time{})) {
			// 	f.DataType = Time
			// } else {
			// 	f.DataType = Struct
			// }

		}

		fileldMap[field.Name] = f
	}
	return fileldMap
}
