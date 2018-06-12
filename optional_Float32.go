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

var _Float32 = time.Time{}
var __Float32 = optional_scanner.ScanBool

// template type Optional(T,scan)

//swagger:type float32
type Float32 optionalFloat32

type optionalFloat32 []float32

const (
	valueKeyFloat32 = iota
)

func scanValueFloat32(input string) (val float32, err error) {
	v, err := optional_scanner.ScanFloat(input)
	return float32(v), err
}

// Of wraps the value in an optional.
func OfFloat32(value float32) Float32 {
	return Float32{valueKeyFloat32: value}
}

func OfFloat32Ptr(ptr *float32) Float32 {
	if ptr == nil {
		return EmptyFloat32()
	} else {
		return OfFloat32(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyFloat32() Float32 {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Float32) Get() (value float32, ok bool) {
	o.If(func(v float32) {
		value = v
		ok = true
	})
	return
}

func (o Float32) IsNil() bool {
	return o == nil
}

func (o Float32) IsPresent() bool {
	return !o.IsBlank()
}

func (o Float32) IsBlank() bool {
	if o.IsNil() {
		return true
	}
	var emptyVal float32
	if o.V() == emptyVal {
		return true
	}
	return false
}

// If calls the function if there is a value wrapped by this optional.
func (o Float32) If(f func(value float32)) {
	if !o.IsNil() {
		f(o[valueKeyFloat32])
	}
}

func (o Float32) ElseFunc(f func() float32) (value float32) {
	if !o.IsNil() {
		o.If(func(v float32) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Float32) Else(elseValue float32) (value float32) {
	return o.ElseFunc(func() float32 { return elseValue })
}

// V returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Float32) V() (value float32) {
	var zero float32
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Float32) String() string {
	if !o.IsNil() {
		return fmt.Sprintf("%v", o.V())
	}
	return fmt.Sprintf("%v", nil)
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Float32) MarshalJSON() (data []byte, err error) {
	if !o.IsNil() {
		return json.Marshal(o[valueKeyFloat32])
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Float32) UnmarshalJSON(data []byte) error {
	//nothing todo if null value
	if string(data) == "null" {
		return nil
	}
	var v float32
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
	*o = OfFloat32(v)
	return nil
}

func (o *Float32) UnmarshalText(data []byte) error {
	return o.Scan(string(data))
}

func (o *Float32) MarshalText() ([]byte, error) {
	if o == nil {
		return []byte(""), nil
	}
	return []byte(o.String()), nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Float32) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.V(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Float32) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v float32
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfFloat32(v)
	return nil
}

func (c Float32) Value() (driver.Value, error) {
	v, ok := c.Get()
	if ok {
		return driver.DefaultParameterConverter.ConvertValue(v)
	}
	return driver.DefaultParameterConverter.ConvertValue(nil)
}

func (c *Float32) Scan(input interface{}) (err error) {
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
		val, err := scanValueFloat32(vv)
		if err != nil {
			return err
		}
		*c = OfFloat32(val)
	}
	return
}
