package script

import (
	"fmt"
	"regexp"
	"strings"
)

// Stellt einen Token dar
type Token struct {
	// Speichert den Aktuellen Token typen ab
	Type TokenDatatype

	// Speichert den extrahierten Wert ab
	Value string

	// Speichert die Zeile ab, auf welcher sich der Wert befindet
	Line uint64

	// Speichert die Position ab, auf welcher sich der Wert befindet
	Pos uint64
}

/*
Gibt an ob es sich um ein Char handelt
wenn ja gibt diese Funktion ein true zurück andernfalls ein false.
*/
func isChar(obj string) bool {
	for _, item := range CHARS {
		if item == obj {
			return true
		}
	}
	return false
}

/*
Gibt an ob es sich um ein Symbol (Zeichen) handelt,
wenn ja gibt diese Funktion ein true zurück andernfalls ein false.
*/
func isSymbol(obj string) bool {
	for _, item := range SYMBOLS {
		if item == obj {
			return true
		}
	}
	return false
}

/*
Gibt an ob es sich um eine Zahl zwischen 0-9 handelt,
wenn ja gibt diese Funktion ein true zurück andernfalls ein false.
*/
func isNumber(obj string) bool {
	for _, item := range NUMBERS {
		if item == obj {
			return true
		}
	}
	return false
}

/*
Gibt an ob es sich um ein UTF-8 Emoji handelt,
wenn ja gibt diese Funktion ein true zurück andernfalls ein false.
*/
func isEmoji(obj string) bool {
	var emojiRx = regexp.MustCompile(`[\x{1F600}-\x{1F6FF}|[\x{2600}-\x{26FF}]`)
	var s = emojiRx.Match([]byte(obj))
	return s
}

/*
Ließt einen Skriptstring in eine Liste einzelner Token (Zeichen) ein
-> []*Token = Gibt die List der eingelesenen Token zurück
-> error = Gibt einen Fehler an welcher beim einlesen aufgetreten ist:
	-- Weniger als 6 Zeichen innerhalb des Skriptes
	-- Unbekanntes / Nicht zulässiges Zeichen
*/
func TokenizationScriptString(script_str string) ([]*Token, error) {
	// Es wird geprüft ob der String mindestens 6 Zeichen großt ist
	// wenn nicht wird der Vorgang mit einer Fehlermeldung abgerbochen
	if len(script_str) < 6 {
		return []*Token{}, fmt.Errorf("Invalid script")
	}

	// Die Zeichen des Strings werden zeichenweise eingelesen und geprüft
	current_line, current_pos := 1, 0
	extracted_tokens := []*Token{}
	for _, otem := range strings.Split(script_str, "") {
		if isChar(otem) {
			reval_obj := Token{Value: otem, Type: TEXT, Line: uint64(current_line), Pos: uint64(current_pos)}
			extracted_tokens = append(extracted_tokens, &reval_obj)
			current_pos++
		} else if isSymbol(otem) {
			reval_obj := Token{Value: otem, Type: SYMBOL, Line: uint64(current_line), Pos: uint64(current_pos)}
			if otem == NEW_LINE {
				current_line++
				current_pos = 0
			} else {
				current_pos++
			}
			extracted_tokens = append(extracted_tokens, &reval_obj)
		} else if isNumber(otem) {
			reval_obj := Token{Value: otem, Type: NUMBER, Line: uint64(current_line), Pos: uint64(current_pos)}
			extracted_tokens = append(extracted_tokens, &reval_obj)
			current_pos++
		} else if isEmoji(otem) {
			reval_obj := Token{Value: otem, Type: EMOJI, Line: uint64(current_line), Pos: uint64(current_pos)}
			extracted_tokens = append(extracted_tokens, &reval_obj)
			current_pos++
		} else {
			return []*Token{}, fmt.Errorf("Invalid script " + otem)
		}
	}

	// Der Vorgang wrude erfolgreich druchgeführt,
	// es wird eine Liste mit Token zurückgegeben
	return extracted_tokens, nil
}
