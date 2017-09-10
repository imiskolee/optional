package optional_scanner

import (
	"errors"
	"strconv"
)

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
