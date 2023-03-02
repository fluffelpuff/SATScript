package script

/*
Liest eine Variablen Dekleration ein
*/
func parseVarDeclaration(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	return &ParsedScriptItem{}, nil
}

/*
Liest eine Variablen Veränderung ein
*/
func parseVarReDeclaration(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	return &ParsedScriptItem{}, nil
}

/*
Liest Map Operationen ein
*/
func parseMapOperation(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	return &ParsedScriptItem{}, nil
}

/*
Liest einen Statischenwert ein
*/
func parseStaticValue(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	// Es wird geprüft ob das Stack am ende ist
	if cursor.IsEnd() {
		return nil, nil
	}

	// Es wird grpüft ob es sich um ein 'return' Keyword handelt
	if *cursor.GetCurrentItem().Type == PR_KEYWORD {
		if *cursor.GetCurrentItem().KeywordValue == KEYWORD_TRUE {
			// Es wird um 1 erhöht auf dem Stack
			cursor.Next()

			// Die Aktuelle Höhe wird gesetzt
			cursor.SetAbolut()

			// Es wird ein neues Script Item erzeugt
			return_value := new(ParsedScriptItem)
			return_value.ItemType = &PS_ITEM_STATIC_BOOL_VALUE
			return_value.BoolValue = true

			// Die Daten werden zurückgegeben
			return return_value, nil
		} else if *cursor.GetCurrentItem().KeywordValue == KEYWORD_FALSE {
			// Es wird um 1 erhöht auf dem Stack
			cursor.Next()

			// Die Aktuelle Höhe wird gesetzt
			cursor.SetAbolut()

			// Es wird ein neues Script Item erzeugt
			return_value := new(ParsedScriptItem)
			return_value.ItemType = &PS_ITEM_STATIC_BOOL_VALUE
			return_value.BoolValue = false

			// Die Daten werden zurückgegeben
			return return_value, nil
		} else {
			return nil, nil
		}
	} else if *cursor.GetCurrentItem().Type == PR_TEXT_STR {
		// Es wird ein neues Script Item erzeugt
		return_value := new(ParsedScriptItem)
		return_value.ItemType = &PS_ITEM_STATIC_STRING_VALUE
		return_value.StringValue = string(*cursor.GetCurrentItem().StringValue)

		// Es wird um 1 erhöht auf dem Stack
		cursor.Next()

		// Die Aktuelle Höhe wird gesetzt
		cursor.SetAbolut()

		return return_value, nil
	} else if *cursor.GetCurrentItem().Type == PR_FLOAT {
		// Es wird ein neues Script Item erzeugt
		return_value := new(ParsedScriptItem)
		return_value.ItemType = &PS_ITEM_STATIC_FLOAT_VALUE
		return_value.FloatValue = *&cursor.GetCurrentItem().FloatValue.Value

		// Es wird um 1 erhöht auf dem Stack
		cursor.Next()

		// Die Aktuelle Höhe wird gesetzt
		cursor.SetAbolut()

		return return_value, nil
	} else if *cursor.GetCurrentItem().Type == PR_INTEGER {
		// Es wird ein neues Script Item erzeugt
		return_value := new(ParsedScriptItem)
		return_value.ItemType = &PS_ITEM_STATIC_INTEGER_VALUE
		return_value.IntValue = cursor.GetCurrentItem().IntegerValue.Value

		// Es wird um 1 erhöht auf dem Stack
		cursor.Next()

		// Die Aktuelle Höhe wird gesetzt
		cursor.SetAbolut()

		return return_value, nil
	} else {
		return nil, nil
	}
}

/*
Liest einen Variablenwert ein
*/
func parseVariableRead(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	// Es wird geprüft ob das Stack am ende ist
	if cursor.IsEnd() {
		return nil, nil
	}

	// Es wird geprüft ob sich ein Stringloser Text auf dem Stack befindet
	if *cursor.GetCurrentItem().Type != PR_TEXT {
		return nil, nil
	}

	// Das Rückgabe Objekt wird erstellt
	return_values := new(ParsedScriptItem)
	return_values.ItemType = &PS_ITEM_READ_VAR
	return_values.VarName = string(*cursor.GetCurrentItem().TextValue)

	// Das Stack wird um Eins erhöht
	cursor.Next()

	// Die Absolute Höhe wird neu gesetzt
	cursor.SetAbolut()

	// Das Erzeugte Objekt wird zurückgegeben
	return return_values, nil
}
