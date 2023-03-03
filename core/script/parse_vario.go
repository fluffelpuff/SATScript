package script

import (
	"fmt"
)

/*
Liest ein Kommentar aus einem übergebenen Cursor ein
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
Liest eine Variablen Dekleration ein
*/
func parseVarDeclaration(cursor *SliceBodyCursor, defines *ParsedScriptDefines) (*ParsedScriptItem, error) {
	// Es wird geprüft ob das Stack am ende ist
	if cursor.IsEnd() {
		return nil, nil
	}

	// Es wird geprüft ob sich ein Stringloser Text auf dem Stack befindet
	if *cursor.GetCurrentItem().Type != PR_TEXT {
		return nil, nil
	}

	// Speichert den Namen der Variable ab
	var_name := *cursor.GetCurrentItem().TextValue

	// Die Höhe auf dem Stack wird erhöht
	cursor.Next()

	// Es wird geprüft ob das Stack am ende ist
	if cursor.IsEnd() {
		cursor.Reset()
		return nil, nil
	}

	// Es wird geprüft ob als nächstes ein : Symbol
	if *cursor.GetCurrentItem().Type != PR_SYMBOL {
		cursor.Reset()
		return nil, nil
	}
	if *cursor.GetCurrentItem().SymbolValue != ColonSymbol {
		cursor.Reset()
		return nil, nil
	}

	// Das Stack wird um 1 Erhöht
	cursor.Next()

	// Es wird geprüft ob sich das Stack am ende befindet
	if cursor.IsEnd() {
		cursor.Reset()
		return nil, nil
	}

	// Es wird geprüft ob als nächstes ein = auf dem Stack liegt
	if *cursor.GetCurrentItem().Type != PR_SYMBOL {
		cursor.Reset()
		return nil, nil
	}
	if *cursor.GetCurrentItem().SymbolValue != EqualToSignSymbol {
		cursor.Reset()
		return nil, nil
	}

	// Das Stack wird um 1 erhöht
	cursor.Next()

	// Es wird geprüft ob sich noch Elemente auf dem Stack befinden
	if cursor.IsEnd() {
		return nil, fmt.Errorf("parseVarDeclaration: invalid variable declaration 4")
	}

	// Das Rückgabe Objekt wird erstellt
	return_value := new(ParsedScriptItem)
	return_value.VarName = string(var_name)
	return_value.ItemType = &PS_ITEM_VAR_DECLARATION

	// Es wird geprüft ob es sich um ein Funktionsaufruf handelt
	func_call, err := parseFunctionCall(cursor)
	if err != nil {
		return nil, fmt.Errorf("parseVarDeclaration: " + err.Error())
	}
	if func_call != nil {
		// Es wird geprüft ob die Variable vorhanden ist
		if !defines.IsADefinedFunction(func_call.FunctionCall.FunctionName) {
			return nil, fmt.Errorf("parseVarDeclaration: invalid variable declaration unkown used variable")
		}

		// Die Daten werden zwischengespeichert
		return_value.VarDeclarationValue = func_call

		// Die Aktuelle Höhe des Stacks wird angepasst
		cursor.SetAbolut()

		// Die Daten werden zurückgegeben
		return return_value, nil
	}

	// Es wird geprüft ob eine andere Variable Verlinkt werden soll
	is_static_value, err := parseStaticValue(cursor)
	if err != nil {
		return nil, fmt.Errorf("parseVarDeclaration: " + err.Error())
	}
	if is_static_value != nil {
		// Die Daten werden zwischengespeichert
		return_value.VarDeclarationValue = is_static_value

		// Die Aktuelle Höhe des Stacks wird angepasst
		cursor.SetAbolut()

		// Die Daten werden zurückgegeben
		return return_value, nil
	}

	// Es wird geprüft ob es sich um eine andere Variable handelt
	is_var_read, err := parseVariableRead(cursor)
	if err != nil {
		return nil, fmt.Errorf("parseVarDeclaration: " + err.Error())
	}
	if is_var_read != nil {
		// Es wird geprüft ob die Variable vorhanden ist
		if !defines.IsADefinedVariable(is_var_read.VarName) {
			// Es wird geprüft ob es sich um eine Funktion handelt
			if !defines.IsADefinedFunction(is_var_read.VarName) {
				return nil, fmt.Errorf("parseVarDeclaration: invalid variable declaration unkown used variable")
			}
		}

		// Die Daten werden zwischengespeichert
		return_value.VarDeclarationValue = is_static_value

		// Die Aktuelle Höhe des Stacks wird angepasst
		cursor.SetAbolut()

		// Die Daten werden zurückgegeben
		return return_value, nil
	}

	// Es handelt sich um einen ungültigen Wert auf dem Stack
	return nil, fmt.Errorf("parseVarDeclaration: invalid variable declaration data")
}

/*
Liest eine Variablen Veränderung ein
*/
func parseVarReDeclaration(cursor *SliceBodyCursor, defines *ParsedScriptDefines) (*ParsedScriptItem, error) {
	// Es wird geprüft ob das Stack am ende ist
	if cursor.IsEnd() {
		return nil, nil
	}

	// Es wird geprüft ob sich ein Stringloser Text auf dem Stack befindet
	if *cursor.GetCurrentItem().Type != PR_TEXT {
		return nil, nil
	}

	// Speichert den Namen der Variable ab
	var_name := *cursor.GetCurrentItem().TextValue

	// Die Höhe auf dem Stack wird erhöht
	cursor.Next()

	// Es wird geprüft ob das Stack am ende ist
	if cursor.IsEnd() {
		cursor.Reset()
		return nil, nil
	}

	// Es wird geprüft ob als nächstes ein = Symbol vorhanden ist
	if *cursor.GetCurrentItem().Type != PR_SYMBOL {
		cursor.Reset()
		return nil, nil
	}
	if *cursor.GetCurrentItem().SymbolValue != EqualToSignSymbol {
		cursor.Reset()
		return nil, nil
	}

	// Das Stack wird um 1 Erhöht
	cursor.Next()

	// Es wird geprüft ob sich das Stack am ende befindet
	if cursor.IsEnd() {
		return nil, fmt.Errorf("parseVarDeclaration: invalid variable declaration 1")
	}

	// Es wird geprüft ob die Verwendete Variable vorhanden ist
	if !defines.IsADefinedVariable(string(var_name)) {
		return nil, fmt.Errorf("parseVarDeclaration: unkown reference variable")
	}

	// Das Rückgabe Objekt wird erstellt
	return_value := new(ParsedScriptItem)
	return_value.VarName = string(var_name)
	return_value.ItemType = &PS_ITEM_VAR_CHANGE_VALUE

	// Es wird geprüft ob es sich um ein Funktionsaufruf handelt
	func_call, err := parseFunctionCall(cursor)
	if err != nil {
		return nil, fmt.Errorf("parseVarDeclaration: " + err.Error())
	}
	if func_call != nil {
		// Es wird geprüft ob die Variable vorhanden ist
		if !defines.IsADefinedFunction(func_call.FunctionCall.FunctionName) {
			return nil, fmt.Errorf("parseVarDeclaration: invalid variable declaration unkown used variable")
		}

		// Die Daten werden zwischengespeichert
		return_value.VarDeclarationValue = func_call

		// Die Aktuelle Höhe des Stacks wird angepasst
		cursor.SetAbolut()

		// Die Daten werden zurückgegeben
		return return_value, nil
	}

	// Es wird geprüft ob eine andere Variable Verlinkt werden soll
	is_static_value, err := parseStaticValue(cursor)
	if err != nil {
		return nil, fmt.Errorf("parseVarDeclaration: " + err.Error())
	}
	if is_static_value != nil {
		// Die Daten werden zwischengespeichert
		return_value.VarDeclarationValue = is_static_value

		// Die Aktuelle Höhe des Stacks wird angepasst
		cursor.SetAbolut()

		// Die Daten werden zurückgegeben
		return return_value, nil
	}

	// Es wird geprüft ob es sich um eine andere Variable handelt
	is_var_read, err := parseVariableRead(cursor)
	if err != nil {
		return nil, fmt.Errorf("parseVarDeclaration: " + err.Error())
	}
	if is_var_read != nil {
		// Es wird geprüft ob die Variable vorhanden ist
		if !defines.IsADefinedVariable(is_var_read.VarName) {
			// Es wird geprüft ob es sich um eine Funktion handelt
			if !defines.IsADefinedFunction(is_var_read.VarName) {
				return nil, fmt.Errorf("parseVarDeclaration: invalid variable declaration unkown used variable")
			}
		}

		// Die Daten werden zwischengespeichert
		return_value.VarDeclarationValue = is_static_value

		// Die Aktuelle Höhe des Stacks wird angepasst
		cursor.SetAbolut()

		// Die Daten werden zurückgegeben
		return return_value, nil
	}

	// Es handelt sich um einen ungültigen Wert auf dem Stack
	return nil, fmt.Errorf("parseVarDeclaration: invalid variable declaration data")
}

/*
Liest Map Operationen ein
*/
func parseMapDeclaration(cursor *SliceBodyCursor) (*ParsedScriptItem, error) {
	// Es wird geprüft ob das Stack am ende ist
	if cursor.IsEnd() {
		return nil, nil
	}

	// Es wird geprüft ob sich ein Stringloser Text auf dem Stack befindet
	if *cursor.GetCurrentItem().Type != PR_TEXT {
		return nil, nil
	}

	// Speichert den Namen der Variable ab
	map_name := *cursor.GetCurrentItem().TextValue

	// Die Höhe auf dem Stack wird erhöht
	cursor.Next()

	// Es wird geprüft ob das Stack am ende ist
	if cursor.IsEnd() {
		cursor.Reset()
		return nil, nil
	}

	// Es wird geprüft ob als nächstes ein = Symbol vorhanden ist
	if *cursor.GetCurrentItem().Type != PR_DATATYPE {
		cursor.Reset()
		return nil, nil
	}
	if *cursor.GetCurrentItem().DatatypeValue != DATATYPE_MAP {
		return nil, nil
	}

	// Das Stack wird um 1 Erhöht
	cursor.Next()

	// Es wird geprüft ob das Stack am ende ist
	if cursor.IsEnd() {
		return nil, fmt.Errorf("parseMapDeclaration: invalid map declaration #1")
	}

	// Es wird geprüft ob als nächstes ein < Symbol auf dem Stack liegt
	if *cursor.GetCurrentItem().Type != PR_SYMBOL {
		return nil, fmt.Errorf("parseMapDeclaration: invalid map declaration #2")
	}
	if *cursor.GetCurrentItem().SymbolValue != OpeningAngleBracketSymbol {
		return nil, fmt.Errorf("parseMapDeclaration: invalid map declaration #3")
	}

	// Das Stack wird um 1 Erhöht
	cursor.Next()

	// Es wird geprüft ob das Stack am ende ist
	if cursor.IsEnd() {
		return nil, fmt.Errorf("parseMapDeclaration: invalid map declaration #4")
	}

	// Es wird geprüft ob es sich um einen Zulässigen Datentypen handelt
	if *cursor.GetCurrentItem().Type != PR_DATATYPE {
		return nil, fmt.Errorf("parseMapDeclaration: invalid map declaration #4")
	}
	if *cursor.GetCurrentItem().DatatypeValue == DATATYPE_MAP {
		return nil, fmt.Errorf("parseMapDeclaration: invalid map declaration #5")
	}

	// Extrahiert den Datentypen
	extr_datatype_1 := *cursor.GetCurrentItem().DatatypeValue

	// Die Höhe auf dem Stack wird erhöht
	cursor.Next()

	// Es wird geprüft ob das Stack am ende ist
	if cursor.IsEnd() {
		return nil, fmt.Errorf("parseMapDeclaration: invalid map declaration #6")
	}

	// Es wird geprüft ob als nächstes ein Komma vorhanden ist
	if *cursor.GetCurrentItem().Type != PR_SYMBOL {
		return nil, fmt.Errorf("parseMapDeclaration: invalid map declaration #7")
	}
	if *cursor.GetCurrentItem().SymbolValue != CommaSymbol {
		return nil, fmt.Errorf("parseMapDeclaration: invalid map declaration #8")
	}

	// Das Stack wird um 1 Erhöht
	cursor.Next()

	// Es wird geprüft ob das Stack am ende ist
	if cursor.IsEnd() {
		return nil, fmt.Errorf("parseMapDeclaration: invalid map declaration #9")
	}

	// Es wird geprüft ob es sich um einen Zulässigen Datentypen handelt
	if *cursor.GetCurrentItem().Type != PR_DATATYPE {
		return nil, fmt.Errorf("parseMapDeclaration: invalid map declaration #10")
	}
	if *cursor.GetCurrentItem().DatatypeValue == DATATYPE_MAP {
		return nil, fmt.Errorf("parseMapDeclaration: invalid map declaration #11")
	}

	// Extrahiert den Datentypen
	extr_datatype_2 := *cursor.GetCurrentItem().DatatypeValue

	// Die Höhe auf dem Stack wird erhöht
	cursor.Next()

	// Es wird geprüft ob das Stack am ende ist
	if cursor.IsEnd() {
		return nil, fmt.Errorf("parseMapDeclaration: invalid map declaration #12")
	}

	// Es wird geprüft ob als nächstes ein < Symbol auf dem Stack liegt
	if *cursor.GetCurrentItem().Type != PR_SYMBOL {
		return nil, fmt.Errorf("parseMapDeclaration: invalid map declaration #13")
	}
	if *cursor.GetCurrentItem().SymbolValue != ClosingAngleBracket {
		return nil, fmt.Errorf("parseMapDeclaration: invalid map declaration #14")
	}

	// Das Stack wird um 1 Erhöht
	cursor.Next()

	// Die Aktuelle Höhe des Stacks wird neu gesetzt
	cursor.SetAbolut()

	// Das Rückgabe Objekt wird gebaut
	map_declaration := new(ParsedScriptItem)
	map_declaration.VarName = string(map_name)
	map_declaration.ItemType = &PS_ITEM_MAP_DECLARATION
	map_declaration.MapDeclaration = new(ParsedScriptMapDeclaration)
	map_declaration.MapDeclaration.L_Type = &extr_datatype_1
	map_declaration.MapDeclaration.R_Type = &extr_datatype_2

	// Die Daten werden zurückgegeben
	return map_declaration, nil
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
		return_value.FloatValue = cursor.GetCurrentItem().FloatValue.Value

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
