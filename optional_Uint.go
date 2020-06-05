// Code generated by gotemplate. DO NOT EDIT.

package optional

import (
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/imiskolee/optional/optional_scanner"
)

var _Uint = time.Time{}
var __Uint = optional_scanner.ScanBool

// template type Optional(T,scan)

//swagger:type uint
type Uint optionalUint

type optionalUint []uint

const (
	valueKeyUint = iota
)

func scanValueUint(input string) (val uint, err error) {
	v, err := optional_scanner.ScanUint(input)
	return uint(v), err
}

func maybeBlankUint() bool {
	var emptyVal uint
	switch reflect.ValueOf(emptyVal).Interface().(type) {
	case string, []byte, bool:
		return true
	default:
		return false
	}
}

// Of wraps the value in an optional.
func OfUint(value uint) Uint {
	return Uint{valueKeyUint: value}
}

func OfUintPtr(ptr *uint) Uint {
	if ptr == nil {
		return EmptyUint()
	} else {
		return OfUint(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyUint() Uint {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Uint) Get() (value uint, ok bool) {
	o.If(func(v uint) {
		value = v
		ok = true
	})
	return
}

func (o Uint) IsNil() bool {
	return o == nil || len(o) == 0
}

func (o Uint) IsPresent() bool {
	return !o.IsBlank()
}

func (o Uint) IsBlank() bool {
	if o.IsNil() {
		return true
	}
	if !maybeBlankUint() {
		return false
	}
	var emptyVal uint
	if o.V() == emptyVal {
		return true
	}
	return false
}

// If calls the function if there is a value wrapped by this optional.
func (o Uint) If(f func(value uint)) {
	if !o.IsNil() {
		f(o[valueKeyUint])
	}
}

func (o Uint) ElseFunc(f func() uint) (value uint) {
	if !o.IsNil() {
		o.If(func(v uint) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Uint) Else(elseValue uint) (value uint) {
	return o.ElseFunc(func() uint { return elseValue })
}

// V returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Uint) V() (value uint) {
	var zero uint
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Uint) String() string {
	if !o.IsNil() {
		return fmt.Sprintf("%v", o.V())
	}
	return fmt.Sprintf("%v", nil)
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uint) MarshalJSON() (data []byte, err error) {
	if !o.IsNil() {
		return json.Marshal(o[valueKeyUint])
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Uint) UnmarshalJSON(data []byte) error {
	//nothing todo if null value
	if string(data) == "null" {
		return nil
	}
	var v uint
	//empty string
	if string(data) == "" || string(data) == "\"\"" || string(data) == "''" {
		*o = OfUint(v)
		return nil
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Bool:
		d, err := strconv.ParseBool(string(data))
		if err != nil {
			return err
		}
		if d {
			data = []byte("true")
		} else {
			data = []byte("false")
		}
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		d, err := strconv.ParseBool(string(data))
		if err == nil {
			if d {
				data = []byte("1")
			} else {
				data = []byte("0")
			}
		}
	}

	err := json.Unmarshal(data, &v)

	//Try unmarshal string numbers with quote
	if err != nil && len(data) > 2 {
		if data[0] == '"' && data[len(data)-1] == '"' {
			data = data[1 : len(data)-1]
		}
		if data[0] == '\'' && data[len(data)-1] == '\'' {
			data = data[1 : len(data)-1]
		}
		err = json.Unmarshal(data, &v)
	}

	//for number to string
	if err != nil {
		d := fmt.Sprintf(`"%s"`, string(data))
		err = json.Unmarshal([]byte(d), &v)
	}

	if err != nil {
		return err
	}
	*o = OfUint(v)
	return nil
}

func (o *Uint) UnmarshalText(data []byte) error {
	return o.Scan(string(data))
}

func (o *Uint) MarshalText() ([]byte, error) {
	if o == nil {
		return []byte(""), nil
	}
	return []byte(o.String()), nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.V(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Uint) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v uint
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfUint(v)
	return nil
}

func (c Uint) Value() (driver.Value, error) {
	v, ok := c.Get()
	if ok {
		return driver.DefaultParameterConverter.ConvertValue(v)
	}
	return driver.DefaultParameterConverter.ConvertValue(nil)
}

func (c *Uint) Scan(input interface{}) (err error) {
	var vv string
	var isvalid = true
	switch reflect.ValueOf(input).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Interface, reflect.Slice:
		if reflect.ValueOf(input).IsNil() {
			isvalid = false
		}
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

	//for empty string
	if vv == "" {
		var zero uint
		*c = OfUint(zero)
		return
	}
	if isvalid {
		val, err := scanValueUint(vv)
		if err != nil {
			return err
		}
		*c = OfUint(val)
	}
	return
}
