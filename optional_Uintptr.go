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

var _Uintptr = time.Time{}
var __Uintptr = optional_scanner.ScanBool

// template type Optional(T,scan)

//swagger:type uintptr
type Uintptr optionalUintptr

type optionalUintptr []uintptr

const (
	valueKeyUintptr = iota
)

func scanValueUintptr(input string) (val uintptr, err error) {
	v, err := optional_scanner.ScanUint(input)
	return uintptr(v), err
}

// Of wraps the value in an optional.
func OfUintptr(value uintptr) Uintptr {
	return Uintptr{valueKeyUintptr: value}
}

func OfUintptrPtr(ptr *uintptr) Uintptr {
	if ptr == nil {
		return EmptyUintptr()
	} else {
		return OfUintptr(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyUintptr() Uintptr {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Uintptr) Get() (value uintptr, ok bool) {
	o.If(func(v uintptr) {
		value = v
		ok = true
	})
	return
}

func (o Uintptr) IsNil() bool {
	return o == nil
}

func (o Uintptr) IsPresent() bool {
	return !o.IsBlank()
}

func (o Uintptr) IsBlank() bool {
	if o.IsNil() {
		return true
	}
	var emptyVal uintptr
	if o.V() == emptyVal {
		return true
	}
	return false
}

// If calls the function if there is a value wrapped by this optional.
func (o Uintptr) If(f func(value uintptr)) {
	if !o.IsNil() {
		f(o[valueKeyUintptr])
	}
}

func (o Uintptr) ElseFunc(f func() uintptr) (value uintptr) {
	if !o.IsNil() {
		o.If(func(v uintptr) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Uintptr) Else(elseValue uintptr) (value uintptr) {
	return o.ElseFunc(func() uintptr { return elseValue })
}

// V returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Uintptr) V() (value uintptr) {
	var zero uintptr
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Uintptr) String() string {
	if !o.IsNil() {
		return fmt.Sprintf("%v", o.V())
	}
	return fmt.Sprintf("%v", nil)
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uintptr) MarshalJSON() (data []byte, err error) {
	if !o.IsNil() {
		return json.Marshal(o[valueKeyUintptr])
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Uintptr) UnmarshalJSON(data []byte) error {
	//nothing todo if null value
	if string(data) == "null" {
		return nil
	}
	var v uintptr
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
	*o = OfUintptr(v)
	return nil
}

func (o *Uintptr) UnmarshalText(data []byte) error {
	return o.Scan(string(data))
}

func (o *Uintptr) MarshalText() ([]byte, error) {
	if o == nil {
		return []byte(""), nil
	}
	return []byte(o.String()), nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uintptr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.V(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Uintptr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v uintptr
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfUintptr(v)
	return nil
}

func (c Uintptr) Value() (driver.Value, error) {
	v, ok := c.Get()
	if ok {
		return driver.DefaultParameterConverter.ConvertValue(v)
	}
	return driver.DefaultParameterConverter.ConvertValue(nil)
}

func (c *Uintptr) Scan(input interface{}) (err error) {
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
		val, err := scanValueUintptr(vv)
		if err != nil {
			return err
		}
		*c = OfUintptr(val)
	}
	return
}
