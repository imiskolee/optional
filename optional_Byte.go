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

var _Byte = time.Time{}
var __Byte = optional_scanner.ScanBool

// template type Optional(T,scan)

//swagger:type byte
type Byte optionalByte

type optionalByte []byte

const (
	valueKeyByte = iota
)

func scanValueByte(input string) (val byte, err error) {
	v, err := optional_scanner.ScanByte(input)
	return byte(v), err
}

func maybeBlankByte() bool {
	var emptyVal byte
	switch reflect.ValueOf(emptyVal).Interface().(type) {
	case string, []byte, bool:
		return true
	default:
		return false
	}
}

// Of wraps the value in an optional.
func OfByte(value byte) Byte {
	return Byte{valueKeyByte: value}
}

func OfBytePtr(ptr *byte) Byte {
	if ptr == nil {
		return EmptyByte()
	} else {
		return OfByte(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyByte() Byte {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Byte) Get() (value byte, ok bool) {
	o.If(func(v byte) {
		value = v
		ok = true
	})
	return
}

func (o Byte) IsNil() bool {
	return o == nil
}

func (o Byte) IsPresent() bool {
	return !o.IsBlank()
}

func (o Byte) IsBlank() bool {
	if o.IsNil() {
		return true
	}
	if !maybeBlankByte() {
		return false
	}
	var emptyVal byte
	if o.V() == emptyVal {
		return true
	}
	return false
}

// If calls the function if there is a value wrapped by this optional.
func (o Byte) If(f func(value byte)) {
	if !o.IsNil() {
		f(o[valueKeyByte])
	}
}

func (o Byte) ElseFunc(f func() byte) (value byte) {
	if !o.IsNil() {
		o.If(func(v byte) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Byte) Else(elseValue byte) (value byte) {
	return o.ElseFunc(func() byte { return elseValue })
}

// V returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Byte) V() (value byte) {
	var zero byte
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Byte) String() string {
	if !o.IsNil() {
		return fmt.Sprintf("%v", o.V())
	}
	return fmt.Sprintf("%v", nil)
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Byte) MarshalJSON() (data []byte, err error) {
	if !o.IsNil() {
		return json.Marshal(o[valueKeyByte])
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Byte) UnmarshalJSON(data []byte) error {
	//nothing todo if null value
	if string(data) == "null" {
		return nil
	}
	var v byte
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
	*o = OfByte(v)
	return nil
}

func (o *Byte) UnmarshalText(data []byte) error {
	return o.Scan(string(data))
}

func (o *Byte) MarshalText() ([]byte, error) {
	if o == nil {
		return []byte(""), nil
	}
	return []byte(o.String()), nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Byte) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.V(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Byte) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v byte
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfByte(v)
	return nil
}

func (c Byte) Value() (driver.Value, error) {
	v, ok := c.Get()
	if ok {
		return driver.DefaultParameterConverter.ConvertValue(v)
	}
	return driver.DefaultParameterConverter.ConvertValue(nil)
}

func (c *Byte) Scan(input interface{}) (err error) {
	var vv string
	var isvalid = true

	if input == nil {
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
		val, err := scanValueByte(vv)
		if err != nil {
			return err
		}
		*c = OfByte(val)
	}
	return
}
