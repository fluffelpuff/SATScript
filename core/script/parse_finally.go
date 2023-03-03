package script

import "fmt"

/*
Gibt an ob es sich um eine Funktionsdekleration handelt, wenn ja wird dise zurückgegebn,
wenn nein wird ein nil wert zurückgegeben.
*/
func parseFunctionDeklaration(prep_script *PreparedUnparsedScript) (*ParsedFunction, error) {
	// Es wird geprüft ob das Stack am ende ist
	if prep_script.StackIsEnd() {
		return nil, nil
	}

	// Der Aktuelle Cursor wird abgerufen
	cursor := prep_script.GetCurrentCursor()

	// Es wird grpüft ob ein Visible Keyword angegeben wurde
	is_a_exporting_func := false
	if *cursor.GetCurrentItem().Type == PR_KEYWORD {
		if *cursor.GetCurrentItem().KeywordValue == KEYWORD_EXPORT {
			is_a_exporting_func = true
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
			if is_a_exporting_func {
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

	// Speichert die Variablen und Funktionen ab welche verfügbar sind
	var_and_funcs := new(ParsedScriptDefines)

	// Der Funktionscube wird eingelesen
	func_name, arguments, rdtypes, err := parseFunctionNameReturnDTypeCubeByCursor(&cursor)
	if err != nil {
		return nil, fmt.Errorf("parseFunctionDeklaration: " + err.Error())
	}

	// Die Bodydaten werden eingelesen
	operations, err := parseCodeBlockTypeCubeByCursor(&cursor, rdtypes, var_and_funcs)
	if err != nil {
		return nil, fmt.Errorf("parseFunctionDeklaration: " + err.Error())
	}

	// Alle Relevanten Funktionsdaten wurden abgerufen, die Cursorhöhe wird übermittelt
	cursor.FinallyPushBackHight()

	// Es wird ein Finales Funktionsobjekt gebaut
	func_obj := new(ParsedFunction)
	func_obj.Arguments = arguments
	func_obj.ReturnType = rdtypes
	func_obj.Name = func_name
	func_obj.IsExporting = is_a_exporting_func
	func_obj.Operations = operations

	// Die Daten werden zurückgegeben
	return func_obj, nil
}

/*
Gibt an ob es sich um einen Kommentar handelt, wenn ja wird diser zurückgegebn,
wenn nein wird ein nil wert zurückgegeben.
*/
func parseCommentDeclaration(prep_script *PreparedUnparsedScript) (*PreparedText, error) {
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

	// Das Aktuelle Textitem wird extrahiert
	extrcted_text_item := cursor.GetCurrentItem().TextValue

	// Der Stack wird um eins erhöht
	cursor.Next()

	// Die Daten werden zurückgegeben
	cursor.FinallyPushBackHight()

	// Es wird ein neues Kommentarobjekt erzeugt
	return extrcted_text_item, nil
}

/*
Gibt an ob es sich um eine Variablen Dekleration handelt, wenn ja wird disee zurückgegebn,
wenn nein wird ein nil wert zurückgegeben.
*/
func parseVariableDeclaration(prep_script *PreparedUnparsedScript) (*PreparedText, error) {
	r := PreparedText("")
	return &r, nil
}

/*
Parst ein Vorberitetes Skript und bereitet es auf das Kompilieren vor
*/
func ParsePreparatedScript(prep_script *PreparedUnparsedScript) error {
	// Die Höhe des Skriptes wird auf das Ende
	prep_script.SetToSVHightEnd()

	// Speichert das Aktuelle Skript ab
	current_script := new(ParsedScript)

	// Die Schleife wird solange ausgeführt, bis alle Einträge auf dem Stack abgearbeitet wurden
	for !prep_script.StackIsEnd() {
		// Es wird geprüft ob es sich um einen Kommentar handelt
		is_comment_result, err := parseCommentDeclaration(prep_script)
		if err != nil {
			return fmt.Errorf("ParsePreparatedScript: " + err.Error())
		}
		if is_comment_result != nil {
			// Es wird geprüft ob der Name bereits Deklariert wurde
			continue
		}

		// Es wird geprüft ob als nächstes eine Funktion Deklariert wurde
		fnc_result, err := parseFunctionDeklaration(prep_script)
		if err != nil {
			return fmt.Errorf("ParsePreparatedScript: " + err.Error())
		}
		if fnc_result != nil {
			// Es wird geprüft ob die Funktion bereits Deklariert ist
			if current_script.NameAlwaysDeclarated(fnc_result.Name) {
				return fmt.Errorf("ParsePreparatedScript: 1")
			}

			// Der Name wird deklariert
			current_script.DeclaratedFunctions = append(current_script.DeclaratedFunctions, &fnc_result.Name)

			// Es wird ein neues Item erzeugt
			cs_item := new(ParsedScriptItem)
			cs_item.ItemType = &PS_ITEM_FUNCTION_DECLARATION

			// Das Item wird zwischengespeichert
			current_script.Items = append(current_script.Items, cs_item)

			// Die Operation wird erzeugt
			continue
		}

		// Es wird ein Fehler ausgelöst, da es sich um einen unbekannten vorgang handelt
		return fmt.Errorf("ParsePreparatedScript: Invalid script, parsing aborted")
	}

	return nil
}
