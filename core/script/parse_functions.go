package script

import (
	"fmt"
)

/*
Liest einen Funktionsaufruf ein
*/
func parseFunctionCall(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	// Es wird geprüft ob das Stack am ende ist
	if cursor.IsEnd() {
		return nil, nil
	}

	// Es wird grpüft ob es sich um ein 'return' Keyword handelt
	if *cursor.GetCurrentItem().Type != PR_TEXT {
		if cursor.GetCurrentItem().Type == &PR_SYMBOL {
			fmt.Println("NO_" + *cursor.GetCurrentItem().SymbolValue)
		}
		return nil, nil
	}

	// Speichert den Aktuellen Namen der Aufgerufenenen Funktion ab
	function_call_name := *cursor.GetCurrentItem().TextValue

	// Das Stack wird um 1 erhöht
	cursor.Next()

	// Es wird geprüft ob es sich noch ein Element auf dem Stack befindet
	if cursor.IsEnd() {
		return nil, nil
	}

	// Es wird geprüft ob als nächstes ein FunctionCallCubevorhanden ist
	func_cube_result, err := parseFuntionCallCube(cursor)
	if err != nil {
		return nil, err
	}
	if func_cube_result == nil {
		cursor.Reset()
		return nil, nil
	}

	// Es wird ein Funktionsaufruf erstetllt
	func_call := new(ParsedScriptFunctionCall)
	func_call.FunctionName = string(function_call_name)
	func_call.Arguments = func_cube_result

	// Ein neues Skript Item wird erzeugt
	return_obj := new(ParsedScriptItem)
	return_obj.ItemType = &PS_ITEM_CALL_FUNCTION
	return_obj.FunctionCall = func_call

	// Die Aktuelle Höhe des Stacks wird angepasst
	cursor.SetAbolut()

	// Gibt die Daten zurück
	return return_obj, nil
}

/*
Überprüft ob die Typen für einen Funktionsaufruf mit den Datentypen der Funktion übereinstimmen
*/
func matchFunctionArgDataTypeForCall(pfarg ParsedFunctionArgument, parsm ParsedScriptItem, defines *ParsedScriptDefines) (bool, error) {
	return true, nil
}

/*
Liest Returnargumente ein
*/
func parseReturnByBodyCursor(cursor *SliceBodyCursor, returns []*PreparedDatatype, defines *ParsedScriptDefines) ([]*ParsedScriptItem, error) {
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

	// Der Stack wird um eins erhöht
	cursor.Next()

	// Die Aktuelle Aboluthöhe wird gesetzt
	cursor.SetAbolut()

	// Es wird geprüft wieviele returns vorhnanden sein müssen
	if len(returns) > 0 {
		if cursor.RestItems() < len(returns) {
			cursor.Reset()
			return nil, fmt.Errorf("parseReturnByBodyCursor: invalid script, need return values")
		}
	}

	// Die Rückgabewerte werden eingelesen
	inner_items, next_is_comma_or_closer := []*ParsedScriptItem{}, false
	for !cursor.IsEnd() && len(inner_items) < len(returns) {
		// Es wird geprüft ob als nächstes ein Komma vorhanden ist
		if next_is_comma_or_closer {
			if *cursor.GetCurrentItem().Type != PR_SYMBOL {
				return nil, fmt.Errorf("parseFuntionCallCube: invalid script a1")
			}
			if *cursor.GetCurrentItem().SymbolValue == CommaSymbol {
				next_is_comma_or_closer = false
				cursor.Next()
				cursor.SetAbolut()
				continue
			} else {
				return nil, fmt.Errorf("parseFuntionCallCube: invalid script a2")
			}
		}

		// Es wird geprüft ob ein Statischerwert zurückgegeben werden soll
		is_static_value, err := parseStaticValue(cursor)
		if err != nil {
			return nil, fmt.Errorf("parseReturnByBodyCursor: " + err.Error())
		}
		if is_static_value != nil {
			// Es wird geprüft ob der Datentyp der passende ist
			if is_static_value.ItemType == &PS_ITEM_STATIC_BOOL_VALUE {
				if *returns[len(inner_items)] != DATATYPE_BOOL {
					return nil, fmt.Errorf("parseReturnByBodyCursor: invalid return data type, 2")
				}
			} else if is_static_value.ItemType == &PS_ITEM_STATIC_FLOAT_VALUE {
				if *returns[len(inner_items)] != DATATYPE_FLOAT {
					return nil, fmt.Errorf("parseReturnByBodyCursor: invalid return data type, 3")
				}
			} else if is_static_value.ItemType == &PS_ITEM_STATIC_INTEGER_VALUE {
				if *returns[len(inner_items)] != DATATYPE_INT {
					return nil, fmt.Errorf("parseReturnByBodyCursor: invalid return data type, 4")
				}
			} else if is_static_value.ItemType == &PS_ITEM_STATIC_STRING_VALUE {
				if *returns[len(inner_items)] != DATATYPE_STRING {
					return nil, fmt.Errorf("parseReturnByBodyCursor: invalid return data type, 5")
				}
			} else if is_static_value.ItemType == &PS_ITEM_STATIC_ADDRESS {
				if *returns[len(inner_items)] != DATATYPE_LN11_ADDRESS {
					if *returns[len(inner_items)] != DATATYPE_CHAIN_ADDRESS {
						if *returns[len(inner_items)] != DATATYPE_ACCOUNT_ADDRESS {
							if *returns[len(inner_items)] != DATATYPE_CONTRACT_ADDRESS {
								if *returns[len(inner_items)] != DATATYPE_UNIVERSE_EP_ADDRESS {
									return nil, fmt.Errorf("parseReturnByBodyCursor: invalid return data type, 5")
								}
							}
						}
					}
				}
			} else if is_static_value.ItemType == &PS_ITEM_STATIC_BYTES {
				if *returns[len(inner_items)] != DATATYPE_BYTES {
					return nil, fmt.Errorf("parseReturnByBodyCursor: invalid return data type, 5")
				}
			} else if is_static_value.ItemType == &PS_ITEM_STATIC_LIST {
				if *returns[len(inner_items)] != DATATYPE_LIST {
					return nil, fmt.Errorf("parseReturnByBodyCursor: invalid return data type, 5")
				}
			} else if is_static_value.ItemType == &PS_ITEM_STATIC_JSON {
				if *returns[len(inner_items)] != DATATYPE_JSON {
					return nil, fmt.Errorf("parseReturnByBodyCursor: invalid return data type, 5")
				}
			} else if is_static_value.ItemType == &PS_ITEM_STATIC_ARRAY {
				if *returns[len(inner_items)] != DATATYPE_ARRAY {
					return nil, fmt.Errorf("parseReturnByBodyCursor: invalid return data type, 5")
				}
			} else if is_static_value.ItemType == &PS_ITEM_STATIC_AMOUNT {
				if *returns[len(inner_items)] != DATATYPE_AMOUNT {
					return nil, fmt.Errorf("parseReturnByBodyCursor: invalid return data type, 5")
				}
			} else if is_static_value.ItemType == &PS_ITEM_STATIC_URL {
				if *returns[len(inner_items)] != DATATYPE_URL {
					return nil, fmt.Errorf("parseReturnByBodyCursor: invalid return data type, 5")
				}
			} else {
				return nil, fmt.Errorf("parseReturnByBodyCursor: invalid return data type, 6")
			}

			// Der Rückgabewerte wird zwiscnegspeichert
			inner_items = append(inner_items, is_static_value)

			// Signalisiert das als nächstes ein Komma oder ein Cube Closer vorhanden ist
			if len(inner_items) < len(returns) {
				next_is_comma_or_closer = true
			} else {
				break
			}

			// Nächste Runde starten
			continue
		}

		// Es wird geprüt ob ein Funktionsaufruf durchgeführt werden soll
		is_function_call, err := parseFunctionCall(cursor)
		if err != nil {
			return nil, fmt.Errorf("parseReturnByBodyCursor: " + err.Error())
		}
		if is_function_call != nil {
			// Es wird geprüft ob es sich um eine Funktion handelt
			if !defines.IsADefinedFunction(is_function_call.FunctionCall.FunctionName) {
				return nil, fmt.Errorf("parseReturnByBodyCursor: invalid script, unkown function call " + is_function_call.FunctionCall.FunctionName)
			}

			// Die Funktions Argumente werden Abgerufen
			fnc_args := defines.GetFunctionParameter(is_function_call.FunctionCall.FunctionName)

			// Es wird geprüft ob die Anazahl der benötigten
			if len(fnc_args) != len(is_function_call.FunctionCall.Arguments) {
				return nil, fmt.Errorf("parseReturnByBodyCursor: invalid stack, function call need args")
			}

			// Die Datentypen der Parameter werden auf übereinstimmung geprüft
			if len(fnc_args) > 0 {
				for i := range fnc_args {
					is_ok, err := matchFunctionArgDataTypeForCall(*fnc_args[i], *is_function_call.FunctionCall.Arguments[i], defines)
					if err != nil {
						return nil, fmt.Errorf("parseReturnByBodyCursor: " + err.Error())
					}
					if !is_ok {
						return nil, fmt.Errorf("parseReturnByBodyCursor: invalid function argument datatype")
					}
				}
			}

			// Die Rückgabewerte der Funktion werden abgerufen
			fnc_returns := defines.GetFunctionReturnTypes(is_function_call.FunctionCall.FunctionName)

			// Es wird geprüft ob die Funktion mindestens einen Wert zurückgibt
			if len(fnc_returns) < 1 {
				return nil, fmt.Errorf("parseReturnByBodyCursor: called function has no return")
			}

			// Es wird geprüft ob die Anzahl der Rückgaben der Funktion die Anzahl der benötigten Rückgaben übersteigt
			if len(inner_items)+len(fnc_returns) > len(returns) {
				return nil, fmt.Errorf("parseReturnByBodyCursor: invalid stack, parsing return to many values for current function")
			}

			// Es wird geprüft ob die Rückgabe Typen der Funktion übereinstimmen
			start_hight := len(inner_items)
			for h := range fnc_returns {
				if *fnc_returns[h] != *returns[start_hight] {
					return nil, fmt.Errorf("parseReturnByBodyCursor: invalid stack, invalid function call return")
				}

				// Die Aktuelle höhe wird nach oben gesetzt
				start_hight++

				// Es wird geprüft ob die Aktuele Höhe das Maximum überschreitet
				if start_hight > len(returns) {
					return nil, fmt.Errorf("parseReturnByBodyCursor: invalid stack, parsing return function is failed")
				}
			}

			// Das Item wird dem Stack hinzugefügt
			inner_items = append(inner_items, is_function_call)

			// Es wird Signalisiert dass als nächstes ein Komma oder ein Cube Closer vorhanden ist
			if len(inner_items) < len(returns) {
				next_is_comma_or_closer = true
			} else {
				break
			}

			// Die Nächste Runde
			continue
		}

		// Es wird geprüft ob eine Variable eingelesen werden soll
		is_var_read, err := parseVariableRead(cursor)
		if err != nil {
			return nil, fmt.Errorf("parseReturnByBodyCursor: " + err.Error())
		}
		if is_var_read != nil {
			// Es wird geprüft ob es sich um eine gültige Variable handelt, sollte es sich nicht um eine Variable handeln so wird geprüft
			// ob es sich um eine Aufrufbare Funktion handelt, wenn nicht wird der Vorgang mit einem Fehler abgebrochen.
			if defines.IsADefinedVariable(is_var_read.VarName) {
				if *returns[len(inner_items)] != *defines.GetVariableDataType(is_var_read.VarName) {
					return nil, fmt.Errorf("parseReturnByBodyCursor: invalid stack, parsing return function is failed, mismatch data type")
				}
			} else if defines.IsADefinedFunction(is_var_read.VarName) {
				if *returns[len(inner_items)] != DATATYPE_CALLABLE {
					return nil, fmt.Errorf("parseReturnByBodyCursor: invalid stack, parsing return function is failed, mismatch want callable")
				}
			} else {
				return nil, fmt.Errorf("parseReturnByBodyCursor: invalid stack, parsing return function is failed, unkown return")
			}

			// Der Rückgabewerte wird zwiscnegspeichert
			inner_items = append(inner_items, is_var_read)

			// Signalisiert das als nächstes ein Komma oder ein Cube Closer vorhanden ist
			if len(inner_items) < len(returns) {
				next_is_comma_or_closer = true
			} else {
				break
			}

			// Nächste Runde starten
			continue
		}

		// Es handet sich um einen unbekannten Stackeintrag
		return nil, fmt.Errorf("parseReturnByBodyCursor: invalid stack, parsing return function is failed")
	}

	// Es wird geprüft ob ein Komma erwartet wurde
	if next_is_comma_or_closer {
		return nil, fmt.Errorf("parseReturnByBodyCursor: invalid return data type, 11")
	}

	// Es wird geprüft ob alle Return Parameter gefunden wurden
	if len(inner_items) != len(returns) {
		return nil, fmt.Errorf("parseReturnByBodyCursor: invalid return data type, 9")
	}

	// Die Operationen werden zurückgegeben
	return inner_items, nil
}

/*
Liest einen Codeblock ein, dieses kommt z.b innerhalb von Funktionen, IF-Bedingungen, Schleifen oder Switch Cases vor.
*/
func parseCodeBlockTypeCubeByCursor(cursor *PreparedUnparsedScriptCursor, returns []*PreparedDatatype, defines *ParsedScriptDefines) ([]*ParsedFunctionOperation, error) {
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

	// Diese Schleife wird solange ausgeführt bis der Codeblock geschlossen wurde
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
			fmt.Println("FUNC_CALL")
			continue
		}

		// Es wird geprüft ob es sich um einen Variablen Dekleration handelt
		parse_var_dec, err := parseVarDeclaration(body_cursor, defines)
		if err != nil {
			return nil, fmt.Errorf("parseCodeBlockTypeCubeByCursor: " + err.Error())
		}
		if parse_var_dec != nil {
			fmt.Println("SET_VAR")
			continue
		}

		// Es wird geprüft ob es sich um eine Variablen veränderung handelt
		parse_var_change, err := parseVarReDeclaration(body_cursor, defines)
		if err != nil {
			return nil, fmt.Errorf("parseCodeBlockTypeCubeByCursor: " + err.Error())
		}
		if parse_var_change != nil {
			fmt.Println("CHANGE_VAR_VALUE")
			continue
		}

		// Es wird geprüft ob es sich um eine Map dekleration handelt
		parse_map_dec, err := parseMapDeclaration(body_cursor)
		if err != nil {
			return nil, fmt.Errorf("parseCodeBlockTypeCubeByCursor: " + err.Error())
		}
		if parse_map_dec != nil {
			fmt.Println("MAP_DECLARED")
			continue
		}

		// Es wird geprüft ob es sich um ein Return handelt
		pars_returns, err := parseReturnByBodyCursor(body_cursor, returns, defines)
		if err != nil {
			return nil, fmt.Errorf("parseCodeBlockTypeCubeByCursor: " + err.Error())
		}
		if pars_returns != nil {
			fmt.Println("RETURN_VALUE")
			continue
		}

		// Es wird ein Fehler ausgelöst, es handelt sich um ein unbekannten Eintrag auf dem Stack
		return nil, fmt.Errorf("parseCodeBlockTypeCubeByCursor: invalid script stack, unkown item")
	}

	// Die Daten werden zurückgegeben
	return []*ParsedFunctionOperation{}, nil
}
