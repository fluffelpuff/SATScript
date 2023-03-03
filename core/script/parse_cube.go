package script

import "fmt"

/*
Liest einen Funktionsaufrufs Cube ein
*/
func parseFuntionCallCube(cursor *SliceBodyCursor) ([]*ParsedScriptItem, error) {
	// Es wird geprüft ob das Stack am ende ist
	if cursor.IsEnd() {
		return nil, nil
	}

	// Es wird geprüft ob als nächstes ein ( Symbol auf dem Stack vorhanden ist
	if *cursor.GetCurrentItem().Type != PR_SYMBOL {
		return nil, nil
	}
	if *cursor.GetCurrentItem().SymbolValue != LestParenthesisSymbol {
		return nil, nil
	}

	// Das Stack wird um eins erhöht
	cursor.Next()

	// Es wird geprüft ob sich noch ein Item auf dem Stack befindet
	if cursor.IsEnd() {
		cursor.Reset()
		return nil, nil
	}

	// Die Höhe des Stacks wird neu Standardisiert
	cursor.SetAbolut()

	// Die Schleife wird ausgeführt bis Entweder das Stack leer ist
	// oder eine Cube Closer gefunden wurde
	has_cube_closer_found, next_is_end_or_comma, extracted := false, false, []*ParsedScriptItem{}
	for !cursor.IsEnd() && !has_cube_closer_found {
		// Es wird geprüft ob als nächstes ein Comma oder
		if next_is_end_or_comma {
			if *cursor.GetCurrentItem().Type != PR_SYMBOL {
				return nil, fmt.Errorf("parseFuntionCallCube: invalid script a1")
			}
			if *cursor.GetCurrentItem().SymbolValue == RightParenthesisSymbol {
				cursor.Next()
				cursor.SetAbolut()
				has_cube_closer_found = true
				next_is_end_or_comma = false
				break
			} else if *cursor.GetCurrentItem().SymbolValue == CommaSymbol {
				cursor.Next()
				cursor.SetAbolut()
				next_is_end_or_comma = false
				continue
			} else {
				return nil, fmt.Errorf("parseFuntionCallCube: invalid script a2")
			}
		} else {
			if *cursor.GetCurrentItem().Type == PR_SYMBOL {
				if *cursor.GetCurrentItem().SymbolValue == RightParenthesisSymbol {
					has_cube_closer_found = true
					cursor.Next()
					cursor.SetAbolut()
					break
				}
			}
		}

		// Es wird geprüft ob es sich um einen Statischen Wert handelt
		is_static_value, err := parseStaticValue(cursor)
		if err != nil {
			return nil, fmt.Errorf("parseReturnByBodyCursor: " + err.Error())
		}
		if is_static_value != nil {
			extracted = append(extracted, is_static_value)
			next_is_end_or_comma = true
			continue
		}

		// Es wird geprüft ob eine Variable eingelesen werden soll
		is_var_read, err := parseVariableRead(cursor)
		if err != nil {
			return nil, fmt.Errorf("parseReturnByBodyCursor: " + err.Error())
		}
		if is_var_read != nil {
			extracted = append(extracted, is_static_value)
			next_is_end_or_comma = true
			continue
		}

		// Es handelt sich um einen ungültigen Cube
		return nil, fmt.Errorf("parseFuntionCallCube: invalid script a3")
	}

	// Die neue Absolute Höhe wird geschrieben
	cursor.SetAbolut()

	// Es wird geprüft ob der Cube geschlossen wurde
	if !has_cube_closer_found {
		return nil, fmt.Errorf("parseFuntionCallCube: invalid script, function call cube has no closer")
	}

	// Es wird geprüft ob ein Comma oder Closer als nächstes vorhadnen sein muss
	if next_is_end_or_comma {
		return nil, fmt.Errorf("parseFuntionCallCube: invalid script, function call cube has no closer or comma")
	}

	// Die Daten werden zurückgegebn
	return extracted, nil
}

/*
Liest einen Funktionscube aus einem übergebenen Cursor ein.
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
		arg_object.Type = &variable_dtype

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

				// Es wird geprüft ob als nächstes ein Komma vorhanden ist
				if *cursor.GetCurrentItem().Type == PR_SYMBOL {
					if *cursor.GetCurrentItem().SymbolValue == CommaSymbol {
						cursor.Next()
					} else if *cursor.GetCurrentItem().SymbolValue != RightParenthesisSymbol {
						return "", []*ParsedFunctionArgument{}, []*PreparedDatatype{}, fmt.Errorf("parseFunctionDeklaration: invalid function declaration 20")
					}
					continue
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
