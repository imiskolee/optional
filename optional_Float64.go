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

var _Float64 = time.Time{}
var __Float64 = optional_scanner.ScanBool

// template type Optional(T,scan)

//swagger:type float64
type Float64 optionalFloat64

type optionalFloat64 []float64

const (
	valueKeyFloat64 = iota
)

func scanValueFloat64(input string) (val float64, err error) {
	v, err := optional_scanner.ScanFloat(input)
	return float64(v), err
}

// Of wraps the value in an optional.
func OfFloat64(value float64) Float64 {
	return Float64{valueKeyFloat64: value}
}

func OfFloat64Ptr(ptr *float64) Float64 {
	if ptr == nil {
		return EmptyFloat64()
	} else {
		return OfFloat64(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyFloat64() Float64 {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Float64) Get() (value float64, ok bool) {
	o.If(func(v float64) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Float64) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o Float64) If(f func(value float64)) {
	if o.IsPresent() {
		f(o[valueKeyFloat64])
	}
}

func (o Float64) ElseFunc(f func() float64) (value float64) {
	if o.IsPresent() {
		o.If(func(v float64) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Float64) Else(elseValue float64) (value float64) {
	return o.ElseFunc(func() float64 { return elseValue })
}

// V returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Float64) V() (value float64) {
	var zero float64
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Float64) String() string {
	if o.IsPresent() {
		return fmt.Sprintf("%v", o.V())
	}
	return fmt.Sprintf("%v", nil)
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Float64) MarshalJSON() (data []byte, err error) {
	if o.IsPresent() {
		return json.Marshal(o[valueKeyFloat64])
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Float64) UnmarshalJSON(data []byte) error {
	//nothing todo if null value
	if string(data) == "null" {
		return nil
	}
	var v float64
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
	*o = OfFloat64(v)
	return nil
}

func (o *Float64) UnmarshalText(data []byte) error {
	return o.Scan(string(data))
}

func (o *Float64) MarshalText() ([]byte, error) {
	if o == nil {
		return []byte(""), nil
	}
	return []byte(o.String()), nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Float64) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.V(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Float64) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v float64
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfFloat64(v)
	return nil
}

func (c Float64) Value() (driver.Value, error) {
	v, ok := c.Get()
	if ok {
		return driver.DefaultParameterConverter.ConvertValue(v)
	}
	return driver.DefaultParameterConverter.ConvertValue(nil)
}

func (c *Float64) Scan(input interface{}) (err error) {
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
		val, err := scanValueFloat64(vv)
		if err != nil {
			return err
		}
		*c = OfFloat64(val)
	}
	return
}
