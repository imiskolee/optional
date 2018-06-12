// Code generated by gotemplate. DO NOT EDIT.

package optional

import (
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"fmt"

	"time"

	"github.com/imiskolee/optional/optional_scanner"
)

var _Uint8 = time.Time{}
var __Uint8 = optional_scanner.ScanBool

// template type Optional(T,scan)

//swagger:type uint8
type Uint8 optionalUint8

type optionalUint8 []uint8

const (
	valueKeyUint8 = iota
)

func scanValueUint8(input string) (val uint8, err error) {
	v, err := optional_scanner.ScanUint(input)
	return uint8(v), err
}

// Of wraps the value in an optional.
func OfUint8(value uint8) Uint8 {
	return Uint8{valueKeyUint8: value}
}

func OfUint8Ptr(ptr *uint8) Uint8 {
	if ptr == nil {
		return EmptyUint8()
	} else {
		return OfUint8(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyUint8() Uint8 {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Uint8) Get() (value uint8, ok bool) {
	o.If(func(v uint8) {
		value = v
		ok = true
	})
	return
}

func (o Uint8) IsNil() bool {
	return o == nil
}

func (o Uint8) IsPresent() bool {
	return !o.IsBlank()
}

func (o Uint8) IsBlank() bool {
	if o.IsNil() {
		return true
	}
	var emptyVal uint8
	if o.V() == emptyVal {
		return true
	}
	return false
}

// If calls the function if there is a value wrapped by this optional.
func (o Uint8) If(f func(value uint8)) {
	if !o.IsNil() {
		f(o[valueKeyUint8])
	}
}

func (o Uint8) ElseFunc(f func() uint8) (value uint8) {
	if !o.IsNil() {
		o.If(func(v uint8) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Uint8) Else(elseValue uint8) (value uint8) {
	return o.ElseFunc(func() uint8 { return elseValue })
}

// V returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Uint8) V() (value uint8) {
	var zero uint8
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Uint8) String() string {
	if !o.IsNil() {
		return fmt.Sprintf("%v", o.V())
	}
	return fmt.Sprintf("%v", nil)
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uint8) MarshalJSON() (data []byte, err error) {
	if !o.IsNil() {
		return json.Marshal(o[valueKeyUint8])
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Uint8) UnmarshalJSON(data []byte) error {
	//nothing todo if null value
	if string(data) == "null" {
		return nil
	}
	var v uint8
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
	*o = OfUint8(v)
	return nil
}

func (o *Uint8) UnmarshalText(data []byte) error {
	return o.Scan(string(data))
}

func (o *Uint8) MarshalText() ([]byte, error) {
	if o == nil {
		return []byte(""), nil
	}
	return []byte(o.String()), nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uint8) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.V(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Uint8) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v uint8
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfUint8(v)
	return nil
}

func (c Uint8) Value() (driver.Value, error) {
	v, ok := c.Get()
	if ok {
		return driver.DefaultParameterConverter.ConvertValue(v)
	}
	return driver.DefaultParameterConverter.ConvertValue(nil)
}

func (c *Uint8) Scan(input interface{}) (err error) {
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
		val, err := scanValueUint8(vv)
		if err != nil {
			return err
		}
		*c = OfUint8(val)
	}
	return
}
