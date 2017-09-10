package optional

import (
	"errors"
	"strconv"
)

//go:generate gotemplate "github.com/imiskolee/optional/templates" Bool(bool,ScanBool)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Byte(byte,ScanByte)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Float32(float32,ScanFloat)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Float64(float64,ScanFloat)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Int(int,ScanInt)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Int16(int16,ScanInt)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Int32(int32,ScanInt)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Int64(int64,ScanInt)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Int8(int8,ScanInt)
//go:generate gotemplate "github.com/imiskolee/optional/templates" String(string,ScanString)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Uint(uint,ScanUint)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Uint16(uint16,ScanUint)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Uint32(uint32,ScanUint)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Uint64(uint64,ScanUint)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Uint8(uint8,ScanUint)
//go:generate gotemplate "github.com/imiskolee/optional/templates" Uintptr(uintptr,ScanUint)

func ScanBool(input string) (bool, error) {
	return strconv.ParseBool(input)
}
func ScanInt(input string) (int64, error) {
	return strconv.ParseInt(input, 10, 64)
}

func ScanUint(input string) (uint64, error) {
	return strconv.ParseUint(input, 10, 64)
}

func ScanFloat(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}

func ScanString(input string) (string, error) {
	return input, nil
}

func ScanByte(input string) (byte, error) {
	if len(input) > 0 {
		return byte(input[0]), nil
	}
	return 0, errors.New("nil string")
}
