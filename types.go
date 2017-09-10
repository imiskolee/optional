package optional

import (
	"errors"
	"strconv"
)

//go:generate gotemplate "github.com/imiskolee/optional/templates" Bool(bool,scanBool)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Byte(byte,scanByte)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Float32(float32,scanFloat)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Float64(float64,scanFloat)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Int(int,scanInt)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Int16(int16,scanInt)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Int32(int32,scanInt)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Int64(int64,scanInt)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Int8(int8,scanInt)
//go:generate gotemplate "github.com/imiskolee/optional/templates" String(string,scanString)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Uint(uint,scanUint)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Uint16(uint16,scanUint)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Uint32(uint32,scanUint)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Uint64(uint64,scanUint)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Uint8(uint8,scanUint)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Uintptr(uintptr,scanUint)

func scanBool(input string) (bool, error) {
	return strconv.ParseBool(input)
}
func scanInt(input string) (int64, error) {
	return strconv.ParseInt(input, 10, 64)
}

func scanUint(input string) (uint64, error) {
	return strconv.ParseUint(input, 10, 64)
}

func scanFloat(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}

func scanString(input string) (string, error) {
	return input, nil
}

func scanByte(input string) (byte, error) {
	if len(input) > 0 {
		return byte(input[0]), nil
	}
	return 0, errors.New("nil string")
}
