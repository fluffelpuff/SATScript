package script

import (
	"fmt"
	"strconv"
)

/*
Wandelt einen String in ein Uint8 Integer um.
-> uint8 = Die Eingelesene Zahl
-> error = Der Fehler welcher beim einlesen des Uint8 aufgetreten ist
*/
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

/*
Wandelt einen String in ein Uint32 Integer um.
-> uint32 = Die Eingelesene Zahl
-> error = Der Fehler welcher beim einlesen des Uint32 aufgetreten ist
*/
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

/*
Gibt an ob es sich um einen zulässigen Datentypen handelt
-> bool = True => Gültiger Datentyp, False => Ungültiger Datentyp
-> *PreparedDatatype = Gibt den Datentypen zurück sofern einer gefunden wurde
-> error = Gibt den Fehler an welcher beim ermitteln des Datentypes aufgetreten ist
*/
func isStringADatatype(strvalue string) (bool, *PreparedDatatype, error) {
	for _, item := range DATATYPES_SLICE {
		if string(*item) == strvalue {
			return true, item, nil
		}
	}
	return false, new(PreparedDatatype), nil
}

/*
Gibt an ob es sich um ein zulässiges Schlüsselwort handelt
-> bool = True => Gültiger Datentyp, False => Ungültiger Datentyp
-> *PreparedKeyword = Gibt das Schlüsselwort zurück sofern eines gefunden wurde
-> error = Gibt den Fehler an welcher beim ermitteln des Schlüsselwortes aufgetreten ist
*/
func isStringAKeyword(strvalue string) (bool, *PreparedKeyword, error) {
	for _, item := range KEYWORD_SLICE {
		if string(*item) == strvalue {
			return true, item, nil
		}
	}
	return false, new(PreparedKeyword), nil
}
