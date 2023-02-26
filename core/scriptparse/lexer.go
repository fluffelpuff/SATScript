package scriptparse

import (
	"fmt"
	"regexp"
	"strings"
)

// Stellt einen Token dar
type Token struct {
	Type  TokenDatatype
	Value string
	Line  uint64
	Pos   uint64
}

// Gibt an ob es sich um ein CHAR handelt
func isChar(obj string) bool {
	for _, item := range CHARS {
		if item == obj {
			return true
		}
	}
	return false
}

// Gibt an ob es sich um ein Symbol handelt
func isSymbol(obj string) bool {
	for _, item := range SYMBOLS {
		if item == obj {
			return true
		}
	}
	return false
}

// Gibt an ob es sich um eine Nummer handelt
func isNumber(obj string) bool {
	for _, item := range NUMBERS {
		if item == obj {
			return true
		}
	}
	return false
}

// Gibt an ob es sich um ein Emoji handelt
func isEmoji(obj string) bool {
	var emojiRx = regexp.MustCompile(`[\x{1F600}-\x{1F6FF}|[\x{2600}-\x{26FF}]`)
	var s = emojiRx.Match([]byte(obj))
	return s
}

// Wandelt einen Scriptstring in eine Liste von Token um
func TokenizationScriptString(script_str string) ([]*Token, error) {
	// Es wird geprüft ob Mindestens 6 Elemente auf dem Stack liegen
	if len(script_str) < 6 {
		return []*Token{}, fmt.Errorf("Invalid script")
	}

	// Die einzelnen Zeichen werden eingelesen
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

	// Die Extrahierten Token werden zurückgegeben
	return extracted_tokens, nil
}
