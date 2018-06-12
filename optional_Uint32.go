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

var _Uint32 = time.Time{}
var __Uint32 = optional_scanner.ScanBool

// template type Optional(T,scan)

//swagger:type uint32
type Uint32 optionalUint32

type optionalUint32 []uint32

const (
	valueKeyUint32 = iota
)

func scanValueUint32(input string) (val uint32, err error) {
	v, err := optional_scanner.ScanUint(input)
	return uint32(v), err
}

func maybeBlankUint32() bool {
	var emptyVal uint32
	switch reflect.ValueOf(emptyVal).Interface().(type) {
	case string, []byte, bool:
		return true
	default:
		return false
	}
}

// Of wraps the value in an optional.
func OfUint32(value uint32) Uint32 {
	return Uint32{valueKeyUint32: value}
}

func OfUint32Ptr(ptr *uint32) Uint32 {
	if ptr == nil {
		return EmptyUint32()
	} else {
		return OfUint32(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyUint32() Uint32 {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Uint32) Get() (value uint32, ok bool) {
	o.If(func(v uint32) {
		value = v
		ok = true
	})
	return
}

func (o Uint32) IsNil() bool {
	return o == nil
}

func (o Uint32) IsPresent() bool {
	return !o.IsBlank()
}

func (o Uint32) IsBlank() bool {
	if o.IsNil() {
		return true
	}
	if !maybeBlankUint32() {
		return false
	}
	var emptyVal uint32
	if o.V() == emptyVal {
		return true
	}
	return false
}

// If calls the function if there is a value wrapped by this optional.
func (o Uint32) If(f func(value uint32)) {
	if !o.IsNil() {
		f(o[valueKeyUint32])
	}
}

func (o Uint32) ElseFunc(f func() uint32) (value uint32) {
	if !o.IsNil() {
		o.If(func(v uint32) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Uint32) Else(elseValue uint32) (value uint32) {
	return o.ElseFunc(func() uint32 { return elseValue })
}

// V returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Uint32) V() (value uint32) {
	var zero uint32
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Uint32) String() string {
	if !o.IsNil() {
		return fmt.Sprintf("%v", o.V())
	}
	return fmt.Sprintf("%v", nil)
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uint32) MarshalJSON() (data []byte, err error) {
	if !o.IsNil() {
		return json.Marshal(o[valueKeyUint32])
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Uint32) UnmarshalJSON(data []byte) error {
	//nothing todo if null value
	if string(data) == "null" {
		return nil
	}
	var v uint32
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
	*o = OfUint32(v)
	return nil
}

func (o *Uint32) UnmarshalText(data []byte) error {
	return o.Scan(string(data))
}

func (o *Uint32) MarshalText() ([]byte, error) {
	if o == nil {
		return []byte(""), nil
	}
	return []byte(o.String()), nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uint32) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.V(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Uint32) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v uint32
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfUint32(v)
	return nil
}

func (c Uint32) Value() (driver.Value, error) {
	v, ok := c.Get()
	if ok {
		return driver.DefaultParameterConverter.ConvertValue(v)
	}
	return driver.DefaultParameterConverter.ConvertValue(nil)
}

func (c *Uint32) Scan(input interface{}) (err error) {
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
		val, err := scanValueUint32(vv)
		if err != nil {
			return err
		}
		*c = OfUint32(val)
	}
	return
}
