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

var _Uint16 = time.Time{}
var __Uint16 = optional_scanner.ScanBool

// template type Optional(T,scan)

//swagger:type uint16
type Uint16 optionalUint16

type optionalUint16 []uint16

const (
	valueKeyUint16 = iota
)

func scanValueUint16(input string) (val uint16, err error) {
	v, err := optional_scanner.ScanUint(input)
	return uint16(v), err
}

func maybeBlankUint16() bool {
	var emptyVal uint16
	switch reflect.ValueOf(emptyVal).Interface().(type) {
	case string, []byte, bool:
		return true
	default:
		return false
	}
}

// Of wraps the value in an optional.
func OfUint16(value uint16) Uint16 {
	return Uint16{valueKeyUint16: value}
}

func OfUint16Ptr(ptr *uint16) Uint16 {
	if ptr == nil {
		return EmptyUint16()
	} else {
		return OfUint16(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyUint16() Uint16 {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Uint16) Get() (value uint16, ok bool) {
	o.If(func(v uint16) {
		value = v
		ok = true
	})
	return
}

func (o Uint16) IsNil() bool {
	return o == nil || len(o) == 0
}

func (o Uint16) IsPresent() bool {
	return !o.IsBlank()
}

func (o Uint16) IsBlank() bool {
	if o.IsNil() {
		return true
	}
	if !maybeBlankUint16() {
		return false
	}
	var emptyVal uint16
	if o.V() == emptyVal {
		return true
	}
	return false
}

// If calls the function if there is a value wrapped by this optional.
func (o Uint16) If(f func(value uint16)) {
	if !o.IsNil() {
		f(o[valueKeyUint16])
	}
}

func (o Uint16) ElseFunc(f func() uint16) (value uint16) {
	if !o.IsNil() {
		o.If(func(v uint16) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Uint16) Else(elseValue uint16) (value uint16) {
	return o.ElseFunc(func() uint16 { return elseValue })
}

// V returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Uint16) V() (value uint16) {
	var zero uint16
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Uint16) String() string {
	if !o.IsNil() {
		return fmt.Sprintf("%v", o.V())
	}
	return fmt.Sprintf("%v", nil)
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uint16) MarshalJSON() (data []byte, err error) {
	if !o.IsNil() {
		return json.Marshal(o[valueKeyUint16])
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Uint16) UnmarshalJSON(data []byte) error {
	//nothing todo if null value
	if string(data) == "null" {
		return nil
	}
	var v uint16
	//empty string
	if string(data) == "" || string(data) == "\"\"" || string(data) == "''" {
		*o = OfUint16(v)
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
	*o = OfUint16(v)
	return nil
}

func (o *Uint16) UnmarshalText(data []byte) error {
	return o.Scan(string(data))
}

func (o *Uint16) MarshalText() ([]byte, error) {
	if o == nil {
		return []byte(""), nil
	}
	return []byte(o.String()), nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uint16) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.V(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Uint16) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v uint16
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfUint16(v)
	return nil
}

func (c Uint16) Value() (driver.Value, error) {
	v, ok := c.Get()
	if ok {
		return driver.DefaultParameterConverter.ConvertValue(v)
	}
	return driver.DefaultParameterConverter.ConvertValue(nil)
}

func (c *Uint16) Scan(input interface{}) (err error) {
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
		var zero uint16
		*c = OfUint16(zero)
		return
	}
	if isvalid {
		val, err := scanValueUint16(vv)
		if err != nil {
			return err
		}
		*c = OfUint16(val)
	}
	return
}
