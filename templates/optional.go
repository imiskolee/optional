package templates

import (
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/imiskolee/optional/optional_scanner"
	"time"
)

var _ = time.Time{}
var __ = optional_scanner.ScanBool

// template type Optional(T,scan)
type T string

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Optional optional

type optional []T

const (
	valueKey = iota
)

func scan(input string) (val T, err error) {
	return *new(T), nil
}

func scanValue(input string) (val T, err error) {
	v, err := scan(input)
	return T(v), err
}

// Of wraps the value in an optional.
func Of(value T) Optional {
	return Optional{valueKey: value}
}

func OfOptionalPtr(ptr *T) Optional {
	if ptr == nil {
		return Empty()
	} else {
		return Of(*ptr)
	}
}

// Empty returns an empty optional.
func Empty() Optional {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Optional) Get() (value T, ok bool) {
	o.If(func(v T) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Optional) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o Optional) If(f func(value T)) {
	if o.IsPresent() {
		f(o[valueKey])
	}
}

func (o Optional) ElseFunc(f func() T) (value T) {
	if o.IsPresent() {
		o.If(func(v T) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Optional) Else(elseValue T) (value T) {
	return o.ElseFunc(func() T { return elseValue })
}

// V returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Optional) V() (value T) {
	var zero T
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Optional) String() string {
	if o.IsPresent() {
		return fmt.Sprintf("%v", o.V())
	}
	return fmt.Sprintf("%v", nil)
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Optional) MarshalJSON() (data []byte, err error) {
	if o.IsPresent() {
		return json.Marshal(o[valueKey])
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Optional) UnmarshalJSON(data []byte) error {
	var v T
	err := json.Unmarshal(data, &v)
	//Try unmarshal string numbers with quote
	if err != nil && len(data) > 2 {
		cpy := data[1 : len(data)-2]
		err = json.Unmarshal(cpy, &v)
	}
	if err != nil {
		return err
	}

	*o = Of(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Optional) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.V(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Optional) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v T
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = Of(v)
	return nil
}

func (c Optional) Value() (driver.Value, error) {
	v, ok := c.Get()
	if ok {
		return driver.DefaultParameterConverter.ConvertValue(v)
	}
	return driver.DefaultParameterConverter.ConvertValue(nil)
}

func (c *Optional) Scan(input interface{}) (err error) {
	var vv string
	var isvalid = true
	switch value := input.(type) {
	case string:
		if len(value) > 0 {
			vv = value
		} else {
			isvalid = false
		}
	case []byte:
		if value != nil && len(value) > 0 {
			vv = string(value)
		} else {
			isvalid = false
		}
	default:
		isvalid = false
	}
	if isvalid {
		val, err := scanValue(vv)
		if err != nil {
			return err
		}
		*c = Of(val)
	}
	return
}
