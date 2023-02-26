package scriptparse

import (
	"fmt"
	"strconv"
)

// Diese Funktion wandelt einen String in ein uin8 um
func convertStringToUint8(value string) (uint8, error) {
	in, err := strconv.Atoi(value)
	if in > 255 || in < 0 {
		return 0, fmt.Errorf("")
	}
	if err != nil {
		return 0, err
	}
	uinted := uint8(in)
	return uinted, nil
}

// Diese Funktion wandelt einen String in ein uin32 um
func covertToStringUint32(value string) (uint32, error) {
	in, err := strconv.Atoi(value)
	if in > 4294967295 || in < 0 {
		return 0, fmt.Errorf("")
	}
	if err != nil {
		return 0, err
	}
	uinted := uint32(in)
	return uinted, nil
}

// Gibt an das es sich um einen Datentypen handelt
func isStringADatatype(strvalue string) (bool, *PreparedDatatype, error) {
	for _, item := range DATATYPES_SLICE {
		if string(*item) == strvalue {
			return true, item, nil
		}
	}
	return false, new(PreparedDatatype), nil
}

// Gibt an ob es sich um ein Schlüsselwort handelt
func isStringAKeyword(strvalue string) (bool, *PreparedKeyword, error) {
	for _, item := range KEYWORD_SLICE {
		if string(*item) == strvalue {
			return true, item, nil
		}
	}
	return false, new(PreparedKeyword), nil
}
