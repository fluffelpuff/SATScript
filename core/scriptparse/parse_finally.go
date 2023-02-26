package scriptparse

import "fmt"

// Diese Funktion ließt einen Funktionscube aus einem Cursoe ein
func parseFunctionNameReturnDTypeCubeByCursor(cursor *PreparedUnparsedScriptCursor) (string, []*ParsedFunctionArgument, []*PreparedDatatype, error) {
	// Es wird geprüft ob sich mindestens 1 Element auf dem Stack befindet
	if cursor.IsEnd() {
		return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionNameReturnDTypeCubeByCursor: invalid function declaration 1")
	}

	// Es wird geprüft ob als nächstes ein Stringloser Textwert vorhanden ist
	if *cursor.GetCurrentItem().Type != PR_TEXT {
		return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 1")
	}

	// Es wird geprüft ob der Text auf der selben Zeile beginnt wo er aufhört
	if cursor.GetCurrentItem().StrLineEnd != 0 {
		return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 2")
	}

	// Der Textwert wird extrahiert
	function_name := *cursor.GetCurrentItem()

	// Das Stack wird um 1 Erhöht
	cursor.Next()

	// Es wird geprüft ob sich auf dem Stack noch ein Element befindet
	if cursor.IsEnd() {
		return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 4")
	}

	// Es wird gerpüft ob als nächstes ein OpenCube vorhanden ist
	if *cursor.GetCurrentItem().Type != PR_SYMBOL {
		return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 5")
	}
	if *cursor.GetCurrentItem().SymbolValue != LestParenthesisSymbol {
		return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 6")
	}

	// Das Stack wird um eins erhöht
	cursor.Next()

	// Es wird geprüft ob der Cursor zu ende ist, wenn ja wird der Vorgang mit einem Fehler abgebrochen
	if cursor.IsEnd() {
		return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 7")
	}

	// Diese Schleife wird solange ausgeführt bis entwender das Stack leer ist, oder ein Cube Closer gefunden wurde
	has_found_closer := false
	for !cursor.IsEnd() {
		// Das Aktuelle Element wird vom Stack geholt
		c_item := cursor.GetCurrentItemANext()

		// Es wird geprüft ob es sich um einen Cube Close handelt
		if *c_item.Type == PR_SYMBOL {
			if *c_item.SymbolValue == RightParenthesisSymbol {
				has_found_closer = true
				break
			}
		}

		// Es wird geprüft ob es sich um einen Stringlosen Text handelt
		if *c_item.Type != PR_TEXT {

		}

		// Der Texwert wird ausgelesen
	}

	// Es wird geprüft ob der Cube geschlossen wurde
	if !has_found_closer {
		return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 8")
	}

	// Es wird geprüft ob sich mindestens 1 Element auf dem Stack befindet
	if cursor.IsEnd() {
		return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 9")
	}

	// Es wird geprüft ob als nächstes mehrere Daten zurückgegeben werden sollen
	if *cursor.GetCurrentItem().Type == PR_SYMBOL {
		// Es wird geprüft ob es sich um ein OpenCube handelt
		if *cursor.GetCurrentItem().SymbolValue == LestParenthesisSymbol {
			// Das Stack wird um 1 nach oben gezählt
			cursor.Next()

			// Es wird geprüft ob mindestns 1 Item auf dem Stack liegt
			if cursor.IsEnd() {
				return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 11")
			}

			// Die Schleife wird ausgeführt bis der Cube geschlossen wird
			has_found_closer_two := false
			for !cursor.IsEnd() {
				// Das Aktuelle Element wird vom Stack geholt
				c_item := cursor.GetCurrentItemANext()

				// Es wird geprüft ob es sich um einen Cube Close handelt
				if *c_item.Type == PR_SYMBOL {
					if *c_item.SymbolValue == RightParenthesisSymbol {
						has_found_closer_two = true
						break
					}
				}
			}

			// Es wird geprüft ob der Cube geschlossen wurde
			if !has_found_closer_two {
				return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 12")
			}

			// Es wird geprüft ob sich mindesntens 1 Eintrag auf dem Stack befindet
			if cursor.IsEnd() {
				return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 13")
			}
		}
	} else if *cursor.GetCurrentItem().Type == PR_DATATYPE {
		// Der Aktuelle Datentyp wird ermittelt

		// Das Stack wird um 1 Erhöht
		cursor.Next()

		// Es wird geprüft ob mindestnes 1 Element auf dem Stack liegt
		if cursor.IsEnd() {
			return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 14")
		}
	}

	// Die Funktion wurd erfolgreich durchgeführt
	return string(*function_name.TextValue), []*ParsedFunctionArgument{}, []*PreparedDatatype{}, nil
}

// Diese Funktion ließt einen Codeblock aus einem Cursor ein
func parseCodeBlockTypeCubeByCursor(cursor *PreparedUnparsedScriptCursor) ([]*PreparedToken, error) {
	// Es wird geprüft ob sich mindestens 1 Element auf dem Stack befindet
	if cursor.IsEnd() {
		return nil, fmt.Errorf("parseCodeBlockTypeCubeByCursor: invalid function declaration 1")
	}

	// Es wird geprüft ob als nächtes ein Codeblock beginnt
	if *cursor.GetCurrentItem().Type != PR_SYMBOL {
		return nil, fmt.Errorf("parseCodeBlockTypeCubeByCursor: invalid function declaration 2")
	}
	if *cursor.GetCurrentItem().SymbolValue != LeftBraceSymbol {
		return nil, fmt.Errorf("parseCodeBlockTypeCubeByCursor: invalid function declaration 2")
	}

	// Die Aktuelle Höhe wird um eins erhöht
	cursor.Next()

	// Es wird geprüft ob sich mindestens noch 1 Element auf dem Stack befidnet
	if cursor.IsEnd() {
		return []*PreparedToken{}, fmt.Errorf("parseCodeBlockTypeCubeByCursor: invalid function declaration 3")
	}

	// Diese Schleife wird solange ausgeführt bis der Cideblock geschlossen wurde
	total_openers, total_closers, body_items := 1, 0, []*PreparedToken{}
	for !cursor.IsEnd() {
		// Das Aktuelle Element wird vom Stack geholt
		c_item := cursor.GetCurrentItemANext()

		// Es wird geprüft ob es sich um einen Cube Close handelt
		if *c_item.Type == PR_SYMBOL {
			if *c_item.SymbolValue == LeftBraceSymbol {
				total_openers++
			} else if *c_item.SymbolValue == RightBraceSymbol {
				body_items = append(body_items, c_item)
				total_closers++
				if total_closers == total_openers {
					break
				} else {
					body_items = append(body_items, c_item)
				}
			} else {
				body_items = append(body_items, c_item)
			}
		}
	}

	// Die Daten werden zurückgegeben
	return body_items, nil
}

// Diese Funktion wird verwendet um eine Funktion einlesen zu können
func parseFunctionDeklaration(prep_script *PreparedUnparsedScript) (*ParsedObject, error) {
	// Es wird geprüft ob das Stack am ende ist
	if prep_script.StackIsEnd() {
		return nil, nil
	}

	// Der Aktuelle Cursor wird abgerufen
	cursor := prep_script.GetCurrentCursor()

	// Es wird grpüft ob ein Visible Keyword angegeben wurde
	is_a_public_func := false
	if *cursor.GetCurrentItem().Type == PR_KEYWORD {
		if *cursor.GetCurrentItem().KeywordValue == KEYWORD_PUBLIC {
			is_a_public_func = true
			cursor.Next()

			// Es wird geprüft ob der Cursor zu ende ist, wenn ja wird der Vorgang mit einem Fehler abgebrochen
			if cursor.IsEnd() {
				return nil, fmt.Errorf("parseFunctionDeklaration: invalid function declaration")
			}

		}
	}

	// Es wird geprüft ob als nächstes ein Keyword mit dem Wert "func" vorhanden ist
	if *cursor.GetCurrentItem().Type == PR_KEYWORD {
		if *cursor.GetCurrentItem().KeywordValue != KEYWORD_FUNCTION {
			if is_a_public_func {
				return nil, fmt.Errorf("parseFunctionDeklaration: invalid script")
			} else {
				return nil, nil
			}
		}
	} else {
		return nil, nil
	}

	// Es wird um 1 Item erhöht
	cursor.Next()

	// Es wird geprüft ob der Cursor zu ende ist, wenn ja wird der Vorgang mit einem Fehler abgebrochen
	if cursor.IsEnd() {
		return nil, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 3")
	}

	// Der Funktionscube wird eingelesen
	func_name, _, _, err := parseFunctionNameReturnDTypeCubeByCursor(&cursor)
	if err != nil {
		return nil, fmt.Errorf("parseFunctionDeklaration: " + err.Error())
	}

	// Die Bodydaten werden eingelesen
	_, err = parseCodeBlockTypeCubeByCursor(&cursor)
	if err != nil {
		return nil, fmt.Errorf("parseFunctionDeklaration: " + err.Error())
	}

	// Alle Relevanten Funktionsdaten wurden abgerufen, die Cursorhöhe wird übermittelt
	cursor.FinallyPushBackHight()

	fmt.Println(func_name)
	_ = is_a_public_func

	// Die Daten werden zurückgegeben
	return &ParsedObject{}, nil
}

// Gibt an ob es sich um ein Kommentar handelt
func parseCommentDeclaration(prep_script *PreparedUnparsedScript) (*ParsedObject, error) {
	// Es wird geprüft ob das Stack am ende ist
	if prep_script.StackIsEnd() {
		return nil, nil
	}

	// Der Aktuelle Cursor wird abgerufen
	cursor := prep_script.GetCurrentCursor()

	// Es wird grpüft ob es sich um einen Kommentar handelt
	if cursor.GetCurrentItem().Type != &PR_COMMENT {
		return nil, nil
	}

	// Der Stack wird um eins erhöht
	cursor.Next()

	// Die Daten werden zurückgegeben
	cursor.FinallyPushBackHight()
	return nil, nil
}

// Diese Funktion wird verwendet um ein Vobereitestet Script zu Parsen
func ParsePreparatedScript(prep_script *PreparedUnparsedScript) error {
	// Die Höhe des Skriptes wird auf das Ende
	prep_script.SetToSVHightEnd()

	// Die Schleife wird solange ausgeführt, bis alle Einträge auf dem Stack abgearbeitet wurden
	for !prep_script.StackIsEnd() {
		// Es wird geprüft ob es sich um einen Kommentar handelt
		is_comment_result, err := parseCommentDeclaration(prep_script)
		if err != nil {
			return fmt.Errorf("ParsePreparatedScript: " + err.Error())
		}
		if is_comment_result != nil {
			continue
		}

		// Es wird geprüft ob als nächstes eine Funktion Deklariert wurde
		fnc_result, err := parseFunctionDeklaration(prep_script)
		if err != nil {
			return fmt.Errorf("ParsePreparatedScript: " + err.Error())
		}
		if fnc_result != nil {
			continue
		}

		// Das Nächste Objekt wir
		item := prep_script.GetCANext()

		if item.CommentValue != nil {
			fmt.Println(*item.Type, item.CommentValue.Value)
		} else if item.TextValue != nil {
			fmt.Println(*item.Type, *item.TextValue)
		} else if item.SymbolValue != nil {
			fmt.Println(*item.Type, *item.SymbolValue)
		} else if item.KeywordValue != nil {
			fmt.Println(*item.Type, *item.KeywordValue)
		} else if item.DatatypeValue != nil {
			fmt.Println(*item.Type, *item.DatatypeValue)
		} else if item.IntegerValue != nil {
			fmt.Println(*item.Type, *item.IntegerValue)
		} else if item.FloatValue != nil {
			fmt.Println(*item.Type, *item.FloatValue)
		} else if item.StringValue != nil {
			fmt.Println(*item.Type, *item.StringValue)
		} else if item.AddressValue != nil {
			fmt.Println(*item.Type, *&item.AddressValue.Value)
		} else {
			fmt.Println(*item.Type)
		}
	}

	return nil
}
