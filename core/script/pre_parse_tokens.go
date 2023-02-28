package script

import "fmt"

// Speichert alle Statischen Symbole zwischen
type PreParsedToken struct {
	// Gibt den Aktuellen Typen an
	Type PreParsedTokenDataType

	// Gibt den Aktuellen Symbolwert an
	SymbolValue SymbolToken

	// GIbt an ob es sich um eine Negative Zahl handelt
	IsMinus bool

	// Gibt an die Line an auf welcher ein String aufhört
	StrLineEnd uint64

	// Gibt den eigentlichen Wert an
	Value string

	// Gibt die Zeile an, auf welcher das Token beginnt
	Line uint64

	// Gibt die Position an, wekche angibt wo der Token beginnt
	StartPos uint64

	// Gibt die Position an, welche angibt wo der der Token endet
	EndPos uint64
}

// Gibt an ob es sich um einen zulässigen String handelt
func nextIsMultiLineCommand(token []*Token) (PreParsedToken, []*Token, bool, error) {
	// Es wird geprüft ob als nächstes ein / vorhanden ist
	if len(token) < 2 {
		return PreParsedToken{}, []*Token{}, false, nil
	}
	if token[0].Type != SYMBOL {
		return PreParsedToken{}, []*Token{}, false, nil
	}
	if token[0].Value != string(SlashSymol) {
		return PreParsedToken{}, []*Token{}, false, nil
	}

	// Es wird geprüft ob als nächstes ein * vorhanden ist
	if token[1].Type != SYMBOL {
		return PreParsedToken{}, []*Token{}, false, nil
	}
	if token[1].Value != string(AsteriskSymbol) {
		return PreParsedToken{}, []*Token{}, false, nil
	}

	// Der Slice wird geupdatet
	new_stra := token[2:]

	// Es wird geprüft ob als nächstes ein zulässiges Item auf dem Parser Stack liegt
	extracted_string := ""
	closer_was_found := false
	extracted := PreParsedToken{Type: P_COMMENT, Line: new_stra[0].Line, StartPos: new_stra[0].Pos - 1}
	for len(new_stra) > 0 {
		item := new_stra[0]
		new_stra = new_stra[1:]

		if item.Type == TEXT {
			extracted_string += item.Value
		} else if item.Type == NUMBER {
			extracted_string += item.Value
		} else {
			if item.Value == string(AsteriskSymbol) {
				if new_stra[0].Value == string(SlashSymol) {
					new_stra = new_stra[1:]
					extracted.EndPos = item.Pos + 3
					extracted.StrLineEnd = item.Line
					closer_was_found = true
					break
				} else {
					extracted_string += item.Value
				}
			} else {
				extracted_string += item.Value
			}
		}
	}

	// Es wird geprüft ob der String geschlossen wurde
	if !closer_was_found {
		fmt.Println(extracted_string)
		return PreParsedToken{}, []*Token{}, true, fmt.Errorf("Invalid script comment at line %d, character %d. Each string requires a \" to open and a \" to close. The string was not closed.\n", int(token[0].Line), int(token[0].Pos))
	}

	// Der Vollständige neue Wert wird hinzugefügt
	extracted.Value = extracted_string

	// Gibt die Daten zurück
	return extracted, new_stra, true, nil
}

// Gibt an ob es sich um einen Zulässigen Einzeiligen String
func nexIsSignleLineCommentText(token []*Token) (PreParsedToken, []*Token, bool, error) {
	// Es wird geprüft ob als nächstes ein / vorhanden ist
	if len(token) < 2 {
		return PreParsedToken{}, []*Token{}, false, nil
	}
	if token[0].Type != SYMBOL {
		return PreParsedToken{}, []*Token{}, false, nil
	}
	if token[0].Value != string(SlashSymol) {
		return PreParsedToken{}, []*Token{}, false, nil
	}

	// Es wird geprüft ob als nächstes ein * vorhanden ist
	if token[1].Type != SYMBOL {
		return PreParsedToken{}, []*Token{}, false, nil
	}
	if token[1].Value != string(SlashSymol) {
		return PreParsedToken{}, []*Token{}, false, nil
	}

	// Der Slice wird geupdatet
	new_stra := token[2:]

	// Es wird geprüft ob als nächstes ein zulässiges Item auf dem Parser Stack liegt
	extracted_string, closer_was_found, last_pos := "", false, 0
	extracted := PreParsedToken{Type: P_COMMENT, Line: new_stra[0].Line, StartPos: new_stra[0].Pos - 1}
	for len(new_stra) > 0 {
		item := new_stra[0]
		new_stra = new_stra[1:]
		last_pos = int(item.Pos)
		extracted.StrLineEnd = item.Line

		if item.Type == TEXT {
			extracted_string += item.Value
		} else if item.Type == NUMBER {
			extracted_string += item.Value
		} else {
			if item.Value == NEW_LINE {
				if len(new_stra) < 1 {
					extracted.EndPos = item.Pos + 1
					closer_was_found = true
					break
				} else {
					new_stra = new_stra[1:]
					extracted.EndPos = item.Pos + 1
					closer_was_found = true
					break
				}
			} else {
				extracted_string += item.Value
			}
		}
	}

	// Es wird geprüft ob der String geschlossen wurde
	if !closer_was_found {
		if len(new_stra) > 1 {
			fmt.Println(extracted_string)
			return PreParsedToken{}, []*Token{}, true, fmt.Errorf("Invalid script comment at line %d, character %d. Each string requires a \" to open and a \" to close. The string was not closed.\n", int(token[0].Line), int(token[0].Pos))
		} else {
			extracted.EndPos = uint64(last_pos + 2)
		}
	}

	// Der Vollständige neue Wert wird hinzugefügt
	extracted.Value = extracted_string

	// Gibt die Daten zurück
	return extracted, new_stra, true, nil
}

// Gibt an ob es sich um einen zulässigen String handelt
func nexIsStringText(token []*Token) (PreParsedToken, []*Token, bool, error) {
	// Es wird geprüft ob als nächstes ein " vorhanden ist
	if len(token) < 1 {
		return PreParsedToken{}, []*Token{}, false, nil
	}
	if token[0].Type != SYMBOL {
		return PreParsedToken{}, []*Token{}, false, nil
	}
	if token[0].Value != string(QuotationMarkSymbol) {
		return PreParsedToken{}, []*Token{}, false, nil
	}

	// Der Slice wird geupdatet
	new_stra := token[1:]

	// Es wird geprüft ob als nächstes ein zulässiges Item auf dem Parser Stack liegt
	extracted_string := ""
	closer_was_found := false
	extracted := PreParsedToken{Type: P_TEXT_STR, Line: new_stra[0].Line, StartPos: new_stra[0].Pos}
	for hight, item := range new_stra {
		if item.Type == TEXT {
			extracted_string += item.Value
		} else if item.Type == NUMBER {
			extracted_string += item.Value
		} else {
			if item.Value == string(QuotationMarkSymbol) {
				extracted.EndPos = item.Pos + 2
				extracted.StrLineEnd = item.Line
				new_stra = new_stra[hight+1:]
				closer_was_found = true
				break
			} else {
				extracted_string += item.Value
			}
		}
	}

	// Es wird geprüft ob der String geschlossen wurde
	if !closer_was_found {
		return PreParsedToken{}, []*Token{}, true, fmt.Errorf("Invalid script starting at line %d, character %d. Each string requires a \" to open and a \" to close. The string was not closed.\n", int(token[0].Line), int(token[0].Pos))
	}

	// Der Vollständige neue Wert wird hinzugefügt
	extracted.Value = extracted_string

	// Gibt die Daten zurück
	return extracted, new_stra, true, nil
}

// Gibt an ob es sich um einen Stringlosen Text handelt
func nextIsStringlessTexx(token []*Token) (PreParsedToken, []*Token, bool, error) {
	// Es wird geprüft ob es sich um einen Text handelt
	if len(token) < 1 {
		return PreParsedToken{}, []*Token{}, false, nil
	}
	if token[0].Type != TEXT {
		return PreParsedToken{}, []*Token{}, false, nil
	}

	// Der Slice wird angepasst
	new_stra := token[:]

	// Es wird geprüft ob als nächstes ein zulässiges Item auf dem Parser Stack liegt
	extracted_string, last_lhight := "", 0
	extracted := PreParsedToken{Type: P_TEXT, Line: new_stra[0].Line, StartPos: new_stra[0].Pos + 1}
	for hight, item := range new_stra {
		last_lhight = hight
		if item.Type == TEXT {
			extracted_string += item.Value
		} else if item.Type == NUMBER {
			if hight == 0 {
				return PreParsedToken{}, []*Token{}, false, fmt.Errorf("nextIsStringlessTexx: Is invalid stringless text")
			} else {
				extracted_string += item.Value
			}
		} else {
			extracted.EndPos = item.Pos + 1
			break
		}
	}

	// Der Vollständige neue Wert wird hinzugefügt
	extracted.Value = extracted_string

	// Der Token Slice wird angepasst
	if len(new_stra) == (last_lhight + 1) {
		new_stra = new_stra[last_lhight+1:]
	} else {
		new_stra = new_stra[last_lhight:]
	}

	// Gibt die Daten zurück
	return extracted, new_stra, true, nil
}

// Ermittelt ob als nächstes ein Symbol vorhanden ist
func nextIsSymbol(token []*Token) (PreParsedToken, []*Token, bool, error) {
	// Es wird geprüft ob es sich um ein Zulässiges Symbol handelt
	// dh kein Leerzeichen oder neue Zeilen
	if len(token) < 1 {
		return PreParsedToken{}, []*Token{}, false, nil
	}
	if token[0].Type != SYMBOL {
		return PreParsedToken{}, []*Token{}, false, nil
	}
	if token[0].Value == NEW_LINE || token[0].Value == SPACE || token[0].Value == TABULATOR {
		return PreParsedToken{}, []*Token{}, false, nil
	}

	// Die Tokenliste wird zwischengespeichert
	new_stra := token[:]

	// Es wird geprüft ob als nächstes ein zulässiges Item auf dem Parser Stack liegt
	extracted := PreParsedToken{}
	for hight, item := range new_stra {
		if item.Type == SYMBOL {
			found := false
			for _, extr_item := range getSymbolList() {
				if string(*extr_item) == item.Value {
					extracted = PreParsedToken{SymbolValue: *extr_item, Line: item.Line, StartPos: new_stra[0].Pos + 1, EndPos: item.Pos + 1, Type: P_SYMBOL}
					new_stra = new_stra[hight+1:]
					found = true
					break
				}
			}
			if found {
				break
			} else {
				fmt.Println(item.Type, item.Value)
				return PreParsedToken{}, []*Token{}, false, fmt.Errorf("Unkown internal error 1")
			}
		} else {
			return PreParsedToken{}, []*Token{}, false, fmt.Errorf("Unkown internal error 1")
		}
	}

	// Gibt die Daten zurück
	return extracted, new_stra, true, nil
}

// Ermittelt ob als nächstes eine Zahl vorhanden ist
func nextIsNumber(token []*Token) (PreParsedToken, []*Token, bool, error) {
	// Es wird geprüft ob als nächtes eine Nummber kommt
	if len(token) < 1 {
		return PreParsedToken{}, []*Token{}, false, nil
	}
	if token[0].Type != NUMBER {
		if token[0].Type == SYMBOL {
			if SymbolToken(token[0].Value) != MinusSignSymbol {
				return PreParsedToken{}, []*Token{}, false, nil
			}
		} else {
			return PreParsedToken{}, []*Token{}, false, nil
		}
	}

	// Die Tokenliste wird zwischengespeichert
	new_stra := token[:]

	// Es wird geprüft ob als nächstes ein zulässiges Item auf dem Parser Stack liegt
	extracted_string, last_lhight, is_minus, has_num_add := "", 0, false, false
	extracted := PreParsedToken{Type: P_NUMBER, Line: new_stra[0].Line, StartPos: uint64(token[0].Pos)}
	for hight, item := range new_stra {
		last_lhight = hight
		if item.Type == NUMBER {
			extracted_string += item.Value
			has_num_add = true
		} else if item.Type == SYMBOL {
			if item.Value == "-" {
				if hight == 0 {
					is_minus = true
					extracted_string += "-"
				} else {
					extracted.EndPos = uint64(hight)
					break
				}
			} else {
				extracted.EndPos = uint64(hight)
				break
			}
		} else {
			extracted.EndPos = uint64(hight)
			break
		}
	}

	// Es wird geprüt ob es sich nur um ein Minuse handelt
	if extracted_string == "-" || !has_num_add {
		return PreParsedToken{}, []*Token{}, false, nil
	}

	// Der Vollständige neue Wert wird hinzugefügt
	extracted.Value = extracted_string

	// Der Rückgabewert wird angepasst
	extracted.IsMinus = is_minus

	// Der Token Slice wird angepasst
	if len(new_stra) == (last_lhight) {
		new_stra = new_stra[last_lhight+1:]
	} else {
		new_stra = new_stra[last_lhight:]
	}

	// Gibt die Daten zurück
	return extracted, new_stra, true, nil
}

// Bereitet eine Tokenliste auf
func PreParseTokenList(token_list []*Token) ([]*PreParsedToken, error) {
	// Es wird geprüft ob Einträge in der Tokenliste vorhanden sind
	if len(token_list) < 1 {
		return []*PreParsedToken{}, fmt.Errorf("PreParseTokenList: invalid script token list")
	}

	// Die Tokenliste wird aufgeräumt und vorarb gepürft
	cr_otem, extracted, last_len := token_list, []*PreParsedToken{}, -1
	for len(cr_otem) > 0 {
		// Es es sich um ein Leerzeichen oder eine neue Zeile handelt, wird der Eintrag übersprungen
		if cr_otem[0].Value == NEW_LINE || cr_otem[0].Value == SPACE || cr_otem[0].Value == TABULATOR {
			cr_otem = cr_otem[1:]
			continue
		}

		// Es wird geprüftob die Länge des Aktuellen Slice mit dem des letzten Slice übereinstimmt
		// sollte diese Bedingung zutreffen, wird das Skript aus Sicherheitsgründen abgebrochen
		if len(cr_otem) == last_len {
			for _, item := range cr_otem {
				fmt.Println(item.Value)
			}
			return []*PreParsedToken{}, fmt.Errorf("This is an invalid script, the process was aborted for security reasons")
		}

		// Die größe des Slice wird Aktualisiert
		last_len = len(cr_otem)

		// Es wird geprüft ob als nächstes Mehrzeiliger Kommentar vorhanden ist
		resolved, new_tokens, is_obj, err := nextIsMultiLineCommand(cr_otem)
		if err != nil {
			return []*PreParsedToken{}, err
		}
		if is_obj {
			extracted = append(extracted, &resolved)
			//fmt.Println("/*"+resolved.Value+"*/", resolved.StartPos, resolved.EndPos, resolved.Line, resolved.StrLineEnd)
			cr_otem = new_tokens
			continue
		}

		// Es wird geprüft ob als nächstes ein Einzeiliger String vorhanden ist
		resolved, new_tokens, is_obj, err = nexIsSignleLineCommentText(cr_otem)
		if err != nil {
			return []*PreParsedToken{}, err
		}
		if is_obj {
			extracted = append(extracted, &resolved)
			//fmt.Println("//"+resolved.Value, resolved.StartPos, resolved.EndPos, resolved.Line, resolved.StrLineEnd)
			cr_otem = new_tokens
			continue
		}

		// Es wird geprüft ob als nächstes ein String vorhanden ist
		resolved, new_tokens, is_obj, err = nexIsStringText(cr_otem)
		if err != nil {
			return []*PreParsedToken{}, err
		}
		if is_obj {
			extracted = append(extracted, &resolved)
			//fmt.Println("'"+resolved.Value+"'", resolved.StartPos, resolved.EndPos, resolved.Line, resolved.StrLineEnd)
			cr_otem = new_tokens
			continue
		}

		// Es wird geprüft ob als nächstes ein Stringloser Text vorhanden ist
		resolved, new_tokens, is_obj, err = nextIsStringlessTexx(cr_otem)
		if err != nil {
			return []*PreParsedToken{}, err
		}
		if is_obj {
			extracted = append(extracted, &resolved)
			//fmt.Println(resolved.Value, resolved.StartPos, resolved.EndPos, resolved.Line)
			cr_otem = new_tokens
			continue
		}

		// Es wird geprüft ob als nächstes eine Zahl vorhanden ist
		resolved, new_tokens, is_obj, err = nextIsNumber(cr_otem)
		if err != nil {
			return []*PreParsedToken{}, err
		}
		if is_obj {
			extracted = append(extracted, &resolved)
			//fmt.Println(resolved.Value, resolved.StartPos, resolved.EndPos, resolved.Line)
			cr_otem = new_tokens
			continue
		}

		// Es wird geprüft ob als nächstes ein Symbol kommt
		resolved, new_tokens, is_obj, err = nextIsSymbol(cr_otem)
		if err != nil {
			return []*PreParsedToken{}, err
		}
		if is_obj {
			extracted = append(extracted, &resolved)
			//fmt.Println(resolved.SymbolValue, resolved.StartPos, resolved.EndPos, resolved.Line)
			cr_otem = new_tokens
			continue
		}

		// Es handelt sich um einen unbekanntes Zeichen
		return []*PreParsedToken{}, fmt.Errorf("Invalid script")
	}

	// Die Extrahieten Daten werden zurücgegeben
	return extracted, nil
}
