// Code generated by gotemplate. DO NOT EDIT.

package optional

import (
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"

	"time"

	"github.com/imiskolee/optional/optional_scanner"
)

var _Int8 = time.Time{}
var __Int8 = optional_scanner.ScanBool

// template type Optional(T,scan)

//swagger:type int8
type Int8 optionalInt8

type optionalInt8 []int8

const (
	valueKeyInt8 = iota
)

func scanValueInt8(input string) (val int8, err error) {
	v, err := optional_scanner.ScanInt(input)
	return int8(v), err
}

func maybeBlankInt8() bool {
	var emptyVal int8
	switch reflect.ValueOf(emptyVal).Interface().(type) {
	case string, []byte, bool:
		return true
	default:
		return false
	}
}

// Of wraps the value in an optional.
func OfInt8(value int8) Int8 {
	return Int8{valueKeyInt8: value}
}

func OfInt8Ptr(ptr *int8) Int8 {
	if ptr == nil {
		return EmptyInt8()
	} else {
		return OfInt8(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyInt8() Int8 {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Int8) Get() (value int8, ok bool) {
	o.If(func(v int8) {
		value = v
		ok = true
	})
	return
}

func (o Int8) IsNil() bool {
	return o == nil
}

func (o Int8) IsPresent() bool {
	return !o.IsBlank()
}

func (o Int8) IsBlank() bool {
	if o.IsNil() {
		return true
	}
	if !maybeBlankInt8() {
		return false
	}
	var emptyVal int8
	if o.V() == emptyVal {
		return true
	}
	return false
}

// If calls the function if there is a value wrapped by this optional.
func (o Int8) If(f func(value int8)) {
	if !o.IsNil() {
		f(o[valueKeyInt8])
	}
}

func (o Int8) ElseFunc(f func() int8) (value int8) {
	if !o.IsNil() {
		o.If(func(v int8) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Int8) Else(elseValue int8) (value int8) {
	return o.ElseFunc(func() int8 { return elseValue })
}

// V returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Int8) V() (value int8) {
	var zero int8
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Int8) String() string {
	if !o.IsNil() {
		return fmt.Sprintf("%v", o.V())
	}
	return fmt.Sprintf("%v", nil)
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Int8) MarshalJSON() (data []byte, err error) {
	if !o.IsNil() {
		return json.Marshal(o[valueKeyInt8])
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Int8) UnmarshalJSON(data []byte) error {
	//nothing todo if null value
	if string(data) == "null" {
		return nil
	}
	var v int8
	err := json.Unmarshal(data, &v)

	//Try unmarshal string numbers with quote
	if err != nil && len(data) > 2 {
		if data[0] == '"' && data[len(data)-1] == '"' {
			data = data[1 : len(data)-1]
		}
		err = json.Unmarshal(data, &v)
	}
	if err != nil {
		return err
	}
	*o = OfInt8(v)
	return nil
}

func (o *Int8) UnmarshalText(data []byte) error {
	return o.Scan(string(data))
}

func (o *Int8) MarshalText() ([]byte, error) {
	if o == nil {
		return []byte(""), nil
	}
	return []byte(o.String()), nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Int8) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.V(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Int8) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v int8
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfInt8(v)
	return nil
}

func (c Int8) Value() (driver.Value, error) {
	v, ok := c.Get()
	if ok {
		return driver.DefaultParameterConverter.ConvertValue(v)
	}
	return driver.DefaultParameterConverter.ConvertValue(nil)
}

func (c *Int8) Scan(input interface{}) (err error) {
	var vv string
	var isvalid = true
	if reflect.ValueOf(input).IsNil() {
		isvalid = false
	}

	if isvalid {
		switch value := input.(type) {
		case string:
			vv = value
		case []byte:
			if value != nil {
				vv = string(value)
			} else {
				isvalid = false
			}
		default:
			vv = fmt.Sprint(input)
		}
	}

	if isvalid {
		val, err := scanValueInt8(vv)
		if err != nil {
			return err
		}
		*c = OfInt8(val)
	}
	return
}
