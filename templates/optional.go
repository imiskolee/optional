package templates

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

var _ = time.Time{}
var __ = optional_scanner.ScanBool

// template type Optional(T,scan)
type T string

//swagger:type T
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

func maybeBlank() bool {
	var emptyVal T
	switch reflect.ValueOf(emptyVal).Interface().(type) {
	case string, []byte, bool:
		return true
	default:
		return false
	}
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

func (o Optional) IsNil() bool {
	return o == nil
}

func (o Optional) IsPresent() bool {
	return !o.IsBlank()
}

func (o Optional) IsBlank() bool {
	if o.IsNil() {
		return true
	}
	if !maybeBlank() {
		return false
	}
	var emptyVal T
	if o.V() == emptyVal {
		return true
	}
	return false
}

// If calls the function if there is a value wrapped by this optional.
func (o Optional) If(f func(value T)) {
	if !o.IsNil() {
		f(o[valueKey])
	}
}

func (o Optional) ElseFunc(f func() T) (value T) {
	if !o.IsNil() {
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
	if !o.IsNil() {
		return fmt.Sprintf("%v", o.V())
	}
	return fmt.Sprintf("%v", nil)
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Optional) MarshalJSON() (data []byte, err error) {
	if !o.IsNil() {
		return json.Marshal(o[valueKey])
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Optional) UnmarshalJSON(data []byte) error {
	//nothing todo if null value
	if string(data) == "null" {
		return nil
	}
	var v T
	//empty string
	if string(data) == "" || string(data) == "\"\""  || string(data) == "''"{
		*o = Of(v)
		return nil
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Bool:
		d,err := strconv.ParseBool(string(data))
		if err != nil {
			return err
		}
		if d {
			data = []byte("true")
		}else{
			data = []byte("false")
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
		d := fmt.Sprintf(`"%s"`,string(data))
		err = json.Unmarshal([]byte(d), &v)
	}

	if err != nil {
		return err
	}
	*o = Of(v)
	return nil
}

func (o *Optional) UnmarshalText(data []byte) error {
	return o.Scan(string(data))
}

func (o *Optional) MarshalText() ([]byte, error) {
	if o == nil {
		return []byte(""), nil
	}
	return []byte(o.String()), nil
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
		var zero T
		*c = Of(zero)
		return
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
