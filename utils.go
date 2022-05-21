package flagx

import (
	"encoding"
	"fmt"
	"reflect"
)

type textValue struct {
	val encoding.TextUnmarshaler
}

func newTextValue(val encoding.TextMarshaler, p encoding.TextUnmarshaler) textValue {
	ptrVal := reflect.ValueOf(p)
	if ptrVal.Kind() != reflect.Ptr {
		panic("flagx: variable value type must be a pointer")
	}
	defVal := reflect.ValueOf(val)
	if defVal.Kind() == reflect.Ptr {
		defVal = defVal.Elem()
	}
	if defVal.Type() != ptrVal.Type().Elem() {
		panic(fmt.Sprintf("flagx: default type does not match variable type: %v != %v", defVal.Type(), ptrVal.Type().Elem()))
	}
	ptrVal.Elem().Set(defVal)
	return textValue{p}
}

func (v textValue) Set(s string) error {
	return v.val.UnmarshalText([]byte(s))
}

func (v textValue) Get() interface{} {
	return v.val
}

func (v textValue) String() string {
	if m, ok := v.val.(encoding.TextMarshaler); ok {
		if b, err := m.MarshalText(); err == nil {
			return string(b)
		}
	}
	return ""
}
