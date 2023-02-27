package scriptparse

import "fmt"

/*
Ließt ein Kommentar aus einem übergebenen Cursor ein
*/
func parseCommentByBodyCursor(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	// Es wird geprüft ob das Stack am ende ist
	if cursor.IsEnd() {
		return nil, nil
	}

	// Es wird grpüft ob es sich um einen Kommentar handelt
	if *cursor.GetCurrentItem().Type != PR_COMMENT {
		return nil, nil
	}

	// Das Aktuelle Textitem wird extrahiert
	extrcted_text_item := cursor.GetCurrentItem().TextValue

	// Die Aktuelle Aboluthöhe wird gesetzt
	cursor.SetAbolut()

	// Der Stack wird um eins erhöht
	cursor.Next()

	// Es wird ein neues ParsedScriptItem erstellt
	re_pars_item := new(ParsedScriptItem)
	re_pars_item.ItemType = &PS_ITEM_COMMENT_DECLARATION
	re_pars_item.CommentValue = extrcted_text_item

	// Die neue Höhe wird Standardtisiert
	cursor.SetAbolut()

	// Es wird ein neues Kommentarobjekt erzeugt
	return re_pars_item, nil
}

/*
Ließt einen Funktionsaufruf ein
*/
func parseFunctionCall(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	return &ParsedScriptItem{}, nil
}

/*
Ließt einen Variablenwert ein
*/
func parseVariableRead(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	return &ParsedScriptItem{}, nil
}

/*
Ließt Mathematische berechnungen ein
*/
func parseMathOperation(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	return &ParsedScriptItem{}, nil
}

/*
Ließt einen Statischenwert ein
*/
func parseStaticValue(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	return &ParsedScriptItem{}, nil
}

/*
Ließt eine IF Bedingung ein
*/
func parseIfStatement(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	return &ParsedScriptItem{}, nil
}

/*
Ließt ein Switchcase ein
*/
func parseSwitchStatement(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	return &ParsedScriptItem{}, nil
}

/*
Ließt eine Variablen Dekleration ein
*/
func parseVarDeclaration(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	return &ParsedScriptItem{}, nil
}

/*
Ließt eine Variablen Veränderung ein
*/
func parseVarReDeclaration(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	return &ParsedScriptItem{}, nil
}

/*
Ließt Map Operationen ein
*/
func parseMapOperation(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	return &ParsedScriptItem{}, nil
}

/*
Ließt einen Sicherheits Event basierten Funktionsaufruf ein
*/
func parseEmitCall(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	return &ParsedScriptItem{}, nil
}

/*
Ließt eine Foorschleife ein
*/
func parseForLoop(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	return &ParsedScriptItem{}, nil
}

/*
Ließt eone Whileschleife ein
*/
func parseWhileLoop(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	return &ParsedScriptItem{}, nil
}

/*
Ließt einen Datentyp basierten Funktionsaufruf ein
*/
func parseDatatypeBasedFunctionCall(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	return &ParsedScriptItem{}, nil
}

// Diese Funktion ließt einen Return ein
func parseReturnByBodyCursor(cursor *SliceBodyCursor, returns []*PreparedDatatype) (*ParsedScriptItem, error) {
	// Es wird geprüft ob das Stack am ende ist
	if cursor.IsEnd() {
		return nil, nil
	}

	// Es wird grpüft ob es sich um ein 'return' Keyword handelt
	if *cursor.GetCurrentItem().Type != PR_KEYWORD {
		return nil, nil
	}
	if *cursor.GetCurrentItem().KeywordValue != KEYWORD_RETURN {
		return nil, nil
	}

	// Die Aktuelle Aboluthöhe wird gesetzt
	cursor.SetAbolut()

	// Der Stack wird um eins erhöht
	cursor.Next()

	// Es wird geprüft wieviele returns vorhnanden sein müssen
	if len(returns) > 0 {
		if cursor.RestItems() < len(returns) {
			cursor.Reset()
			return nil, fmt.Errorf("parseReturnByBodyCursor: invalid script, need return values")
		}
	}

	// Die Einzelnen Rückgabewerte werden abgerufen
	for _, item := range cursor.GetSlice() {
		fmt.Println(&item)
	}

	return nil, nil
}

/*
Ließt einen Funktionscube aus einem übergebenen Cursor ein.
# (var_name DataType, var_name_2 DataType) (bool, string, etc...) #
*/
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
	has_found_closer, function_parms := false, []*ParsedFunctionArgument{}
	for !cursor.IsEnd() {
		// Es wird geprüft ob es sich um einen Cube Close handelt
		if *cursor.GetCurrentItem().Type == PR_SYMBOL {
			if *cursor.GetCurrentItem().SymbolValue == RightParenthesisSymbol {
				has_found_closer = true
				cursor.Next()
				break
			} else if *cursor.GetCurrentItem().SymbolValue == CommaSymbol {
				cursor.Next()
			}
		}

		// Es wird geprüft ob es sich noch ein Element auf dem Stack befindet
		if cursor.IsEnd() {
			return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: 1")
		}

		// Es wird geprüft ob ein Text angegeben wurde
		if *cursor.GetCurrentItem().Type != PR_TEXT {
			return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: 2")
		}

		// Der Name der Variable wird extrahiert
		variable_name := *cursor.GetCurrentItem().TextValue

		// Es wird das nächste Item auf dem Stack asugewählt
		cursor.Next()

		// Es wird geprüft ob noch ein Eintrag vorhanden ist
		if cursor.IsEnd() {
			return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: 3")
		}

		// Es wird geprüft ob als nächstes ein Datentyp angegeben wurde
		if *cursor.GetCurrentItem().Type != PR_DATATYPE {
			fmt.Println(*cursor.GetCurrentItem().Type)
			return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: 4")
		}

		// Der Datentyp wird extrahiert
		variable_dtype := *cursor.GetCurrentItem().DatatypeValue

		// Es wird das nächste Item auf dem Stack asugewählt
		cursor.Next()

		// Es wird geprüft ob noch ein Eintrag vorhanden ist
		if cursor.IsEnd() {
			return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: 3")
		}

		// Das Objekt wird erzeugt
		arg_object := new(ParsedFunctionArgument)
		arg_object.Name = string(variable_name)
		arg_object.Type = variable_dtype

		// Das Objekt wird zwischengespeichert
		function_parms = append(function_parms, arg_object)
	}

	// Es wird geprüft ob der Cube geschlossen wurde
	if !has_found_closer {
		return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 8")
	}

	// Es wird geprüft ob sich mindestens 1 Element auf dem Stack befindet
	if cursor.IsEnd() {
		return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 9c")
	}

	// Es wird geprüft ob als nächstes mehrere Daten zurückgegeben werden sollen
	extracted_data_types := []*PreparedDatatype{}
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

				// Es wird geprüft ob es sich um einen Datentyp handelt
				if *c_item.Type != PR_DATATYPE {
					return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 17")
				}

				// Der Datentyp wird extrahiert
				extracted_data_types = append(extracted_data_types, c_item.DatatypeValue)

				// Das Item wird von dem Stack extrahiert
				cursor.Next()

				// Es wird geprüft ob sich mindestens noch ein Eintrag auf dem Stack befindet
				if cursor.IsEnd() {
					return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 18")
				}

				// Es wird geprüft ob als nächstes ein Komma vorhanden ist
				if *cursor.GetCurrentItem().Type == PR_SYMBOL {
					if *cursor.GetCurrentItem().SymbolValue == CommaSymbol {
						cursor.Next()
					} else if *cursor.GetCurrentItem().SymbolValue != RightParenthesisSymbol {
						return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 20")
					}
				} else {
					return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 21")
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
		// Der Datentyp wird hinzugefügt
		extracted_data_types = append(extracted_data_types, cursor.GetCurrentItem().DatatypeValue)

		// Das Stack wird um 1 Erhöht
		cursor.Next()

		// Es wird geprüft ob mindestnes 1 Element auf dem Stack liegt
		if cursor.IsEnd() {
			return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 14")
		}
	}

	// Die Funktion wurd erfolgreich durchgeführt
	return string(*function_name.TextValue), function_parms, extracted_data_types, nil
}

/*
Ließt einen Codeblock ein, dieses kommt z.b innerhalb von Funktionen, IF-Bedingungen, Schleifen oder Switch Cases vor.
*/
func parseCodeBlockTypeCubeByCursor(cursor *PreparedUnparsedScriptCursor, returns []*PreparedDatatype) ([]*ParsedFunctionOperation, error) {
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
		return nil, fmt.Errorf("parseCodeBlockTypeCubeByCursor: invalid function declaration 3")
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
				total_closers++
				if total_closers == total_openers {
					break
				} else {
					body_items = append(body_items, c_item)
				}
			} else {
				body_items = append(body_items, c_item)
			}
		} else {
			body_items = append(body_items, c_item)
		}
	}

	// Es wird ein neuer Cursor aus den Bodydaten erzeugt
	body_cursor := new(SliceBodyCursor)
	body_cursor.Items = body_items
	body_cursor.CurrentHight = 0

	// Der Body wird abgearbeitet
	for !body_cursor.IsEnd() {
		// Es wird geprüft ob es sich um ein Kommentar handelt
		pars_script_item, err := parseCommentByBodyCursor(body_cursor)
		if err != nil {
			return nil, fmt.Errorf("parseCodeBlockTypeCubeByCursor: " + err.Error())
		}
		if pars_script_item != nil {
			fmt.Println("COMMENT_READ")
			continue
		}

		// Es wird geprüft ob es sich um einen Funktionsaufruf handelt
		pars_func_call, err := parseFunctionCall(body_cursor)
		if err != nil {
			return nil, fmt.Errorf("parseCodeBlockTypeCubeByCursor: " + err.Error())
		}
		if pars_func_call != nil {
			fmt.Println("COMMENT_READ")
			continue
		}

		// Es wird geprüft ob es sich um eine Variablen Dekleration handelt
		is_var_dec, err := parseFunctionCall(body_cursor)
		if err != nil {
			return nil, fmt.Errorf("parseCodeBlockTypeCubeByCursor: " + err.Error())
		}
		if is_var_dec != nil {
			fmt.Println("COMMENT_READ")
			continue
		}

		// Es wird geprüft ob es sich um ein Return handelt
		pars_returns, err := parseReturnByBodyCursor(body_cursor, returns)
		if err != nil {
			return nil, fmt.Errorf("parseCodeBlockTypeCubeByCursor: " + err.Error())
		}
		if pars_returns != nil {
			break
		}

		// Es wird ein Fehler ausgelöst, es handelt sich um ein unbekannten Eintrag auf dem Stack
		return nil, fmt.Errorf("parseCodeBlockTypeCubeByCursor: invalid script stack, unkown item")
	}

	// Die Daten werden zurückgegeben
	return []*ParsedFunctionOperation{}, nil
}
