package scriptparse

import (
	"fmt"
	"math/big"
	"satvm/core/address"
)

// Stellt die Verwendete Version der Skriptsprache dar
type ScriptVersion struct {
	LanguageName      string
	BuildVersion      uint32
	MinjorVersion     uint8
	ReleaseVersion    uint8
	StackHight        int
	AllowNewerVersion bool
}

// Stellt ein bearbeiteten Token dar
type PreparedToken struct {
	// Gibt den Typen des aktuellen Tokens an
	Type *PreparedScriptTokenDataType

	// Gibt den Symbolwert an
	SymbolValue *SymbolToken

	// Gibt die Zeile an auf welcher das Objekt endet
	StrLineEnd uint64

	// Gibt die Zeile an auf welcher das Objekt beginnt
	StartLine uint64

	// Gibt die Startposition an
	StartPos uint64

	// Gibt die Endposition an
	EndPos uint64

	// Gibt ein Kommentar an
	CommentValue *PreparedComment

	// Gibt einen Stringwert an
	StringValue *PreparedString

	// Gibt einen Zahlenwert an
	IntegerValue *PreparedInteger

	// Gibt einen Floatwert an
	FloatValue *PreparedFloat

	// Gibt den Keyword wert an
	KeywordValue *PreparedKeyword

	// Gibt den Adresswert an
	AddressValue *PreparedAddress

	// Gibt den Datentypen wert an
	DatatypeValue *PreparedDatatype

	// Gibt den Regelwert an
	RuleValue *PreparedRules

	// Gibt den Preparierten Text an
	TextValue *PreparedText

	// Gibt die Skriptversion an
	ScriptVersion *ScriptVersion
}

// Stelt ein Parsingfähiges Skript bereit
type PreparedUnparsedScript struct {
	LanguageSpeficationVersion *ScriptVersion
	PreparatedTokens           []*PreparedToken
	currentHight               uint
}

// Gibt das Aktuelle Preparierte Obejkt aus
func (obj *PreparedUnparsedScript) GetCurrentToken() *PreparedToken {
	return obj.PreparatedTokens[obj.currentHight]
}

// Setzt die Angabe des Aktuellen Objekt um 1 nach Oben
func (obj *PreparedUnparsedScript) NextStackHight() {
	if obj.StackIsEnd() {
		return
	}
	obj.currentHight++
}

// Gibt das Aktuelle Objekt aus und Setzt die Höhe um 1 nach Oben
func (obj *PreparedUnparsedScript) GetCANext() *PreparedToken {
	c_obj := obj.GetCurrentToken()
	obj.NextStackHight()
	return c_obj
}

// Setzt die Höhe des Stacks an des Ende der Skript Versions Informaition +1
func (obj *PreparedUnparsedScript) SetToSVHightEnd() {
	obj.currentHight = uint(obj.LanguageSpeficationVersion.StackHight)
}

// Die Aktuelle Stackhöhe wird zurückgesetzt
func (obj *PreparedUnparsedScript) SetToZero() {
	obj.currentHight = 0
}

// Gitb an ob noch ein Objekt verfügbar ist
func (obj *PreparedUnparsedScript) StackIsEnd() bool {
	return len(obj.PreparatedTokens) == int(obj.currentHight)
}

// Gibt den Aktuellen Cursor zurück
func (obj *PreparedUnparsedScript) GetCurrentCursor() PreparedUnparsedScriptCursor {
	return PreparedUnparsedScriptCursor{PreparedUnparsedScriptObject: obj, CurrentHight: obj.currentHight}
}

// Diese Funktion ließt die Verwendete Sprache und die Version des Skriptes ein
func getScriptLangAndVersion(token_list []*PreParsedToken) ([]*PreparedToken, *ScriptVersion, []*PreParsedToken, error) {
	// Es werden alle Kommentare entfernt.
	extracted_tokens, phight := []*PreparedToken{}, 0
	for hight, item := range token_list {
		if item.Type == P_COMMENT {
			token_comment := PreparedComment{Multiline: true, Value: item.Value}
			new_token := PreparedToken{Type: &PR_COMMENT, StartPos: item.StartPos, StrLineEnd: item.StrLineEnd, EndPos: item.EndPos, CommentValue: &token_comment}
			extracted_tokens = append(extracted_tokens, &new_token)
		} else {
			phight = hight
			break
		}
	}

	// Das Slice wird angepasst
	new_slice := token_list[phight:]

	// Es wird geprüft ob Mindestens 8 Einträge auf der Token List vorhanden ist
	if len(new_slice) < 8 {
		return []*PreparedToken{}, &ScriptVersion{}, []*PreParsedToken{}, fmt.Errorf("getScriptLangAndVersion: Each script needs to specify the language used and the version used.")
	}

	// Es wird geprüft ob es sich um pragma Schlüsselwort handelt
	if new_slice[0].Type != P_TEXT {
		return []*PreparedToken{}, &ScriptVersion{}, []*PreParsedToken{}, fmt.Errorf("getScriptLangAndVersion: ")
	}
	if new_slice[0].Value != "pragma" {
		return []*PreparedToken{}, &ScriptVersion{}, []*PreParsedToken{}, fmt.Errorf("getScriptLangAndVersion: ")
	}

	// Die Daten werden neu zwischengespeichert
	new_slice = new_slice[1:]

	// Es wird geprüft ob es sich um ein satscript handelt
	if new_slice[0].Type != P_TEXT {
		return []*PreparedToken{}, &ScriptVersion{}, []*PreParsedToken{}, fmt.Errorf("getScriptLangAndVersion: ")
	}
	if new_slice[0].Value != "satscript" {
		return []*PreparedToken{}, &ScriptVersion{}, []*PreParsedToken{}, fmt.Errorf("getScriptLangAndVersion: ")
	}

	// Die Daten werden neu zwischengespeichert
	new_slice = new_slice[1:]

	// Es wird geprüft ob ein Symbol oder eine Nummber vorhanden ist
	force_current_version := false
	if new_slice[0].Type == P_SYMBOL {
		if new_slice[0].SymbolValue != CaretSymbol {
			return []*PreparedToken{}, &ScriptVersion{}, []*PreParsedToken{}, fmt.Errorf("getScriptLangAndVersion: ")
		} else {
			force_current_version = true
			new_slice = new_slice[1:]
		}
	} else {
		if new_slice[0].Type != P_NUMBER {
			return []*PreparedToken{}, &ScriptVersion{}, []*PreParsedToken{}, fmt.Errorf("getScriptLangAndVersion: ")
		}
	}

	// Es wird geprüft ob als nächstes eine Nummer auf dem Stack vorhanden ist
	if new_slice[0].Type != P_NUMBER {
		return []*PreparedToken{}, &ScriptVersion{}, []*PreParsedToken{}, fmt.Errorf("getScriptLangAndVersion: ")
	}

	// Es wird geprüft ob als nächstes ein Punkt vorhanden ist
	if new_slice[1].Type != P_SYMBOL && new_slice[1].SymbolValue != PeriodSymbol {
		return []*PreparedToken{}, &ScriptVersion{}, []*PreParsedToken{}, fmt.Errorf("getScriptLangAndVersion: ")
	}

	// Es wird geprüft ob als nächstes eine Nummber vorhanden ist
	if new_slice[2].Type != P_NUMBER {
		return []*PreparedToken{}, &ScriptVersion{}, []*PreParsedToken{}, fmt.Errorf("getScriptLangAndVersion: ")
	}

	// Es wird geprüft ob als nächstes ein Punkt vorhanden ist
	if new_slice[3].Type != P_SYMBOL && new_slice[3].SymbolValue != PeriodSymbol {
		return []*PreparedToken{}, &ScriptVersion{}, []*PreParsedToken{}, fmt.Errorf("getScriptLangAndVersion: ")
	}

	// Es wird geprüft ob als nächstes eine Nummber vorhanden ist
	if new_slice[4].Type != P_NUMBER {
		return []*PreparedToken{}, &ScriptVersion{}, []*PreParsedToken{}, fmt.Errorf("getScriptLangAndVersion: ")
	}

	// Es wird geprüft ob es sich um eine gültie Build version handelt
	uinted_build_number, err := covertToStringUint32(new_slice[0].Value)
	if err != nil {
		return []*PreparedToken{}, &ScriptVersion{}, []*PreParsedToken{}, fmt.Errorf("getScriptLangAndVersion: " + err.Error())
	}

	// Es wird geprüft ob es sich um eine gültige Minjor version handelt
	minjor_number, err := convertStringToUint8(new_slice[2].Value)
	if err != nil {
		return []*PreparedToken{}, &ScriptVersion{}, []*PreParsedToken{}, fmt.Errorf("getScriptLangAndVersion: " + err.Error())
	}

	// Es wird geprüft ob es sich um eine Release Version handelt
	release_number, err := convertStringToUint8(new_slice[4].Value)
	if err != nil {
		return []*PreparedToken{}, &ScriptVersion{}, []*PreParsedToken{}, fmt.Errorf("getScriptLangAndVersion: " + err.Error())
	}

	// Das Aktuelle Slice wird angepasst
	new_slice = new_slice[5:]

	// Die Antwortendaten werden zurückgegeben
	resolved_data := new(ScriptVersion)
	resolved_data.BuildVersion = uinted_build_number
	resolved_data.MinjorVersion = minjor_number
	resolved_data.ReleaseVersion = release_number
	resolved_data.AllowNewerVersion = force_current_version
	resolved_data.StackHight = len(extracted_tokens) + 1

	// Es wird Informationsobjekt erstellt
	extracted_tokens = append(extracted_tokens, &PreparedToken{Type: &PR_VERSION, ScriptVersion: resolved_data})

	// Die Daten werden zurückgegeben
	return extracted_tokens, resolved_data, new_slice, nil
}

// Diese Funktion giht an ob es sich um ein Keyword handelt
func isKeyword(token_list []*PreParsedToken, scrversion *ScriptVersion) (bool, *PreparedToken, []*PreParsedToken, error) {
	// Es wird geprüft ob mindestens 1 Eintrag vorhanden ist
	if len(token_list) < 1 {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}

	// Es wird geprüft ob es sich um einen Stringlosen Text handelt
	if token_list[0].Type != P_TEXT {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}

	// Es wird geprüft ob es sich um ein Schlüsselwort handelt
	is_keyword, keyw, err := isStringAKeyword(token_list[0].Value)
	if err != nil {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}
	if !is_keyword {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}

	// Es wird ein neuer Rückgabewert erzeugt
	new_return := new(PreparedToken)
	new_return.Type = &PR_KEYWORD
	new_return.KeywordValue = keyw
	new_return.StartPos = token_list[0].StartPos
	new_return.EndPos = token_list[0].EndPos
	new_return.StartLine = token_list[0].Line
	new_return.StrLineEnd = token_list[0].StrLineEnd

	// Es wird geprüft ob es sich um ein Keyword handelt
	return true, new_return, token_list[1:], nil
}

// Diese Funktion gibt an ob es sich um eine Krypto Adresse handelt (BTC, LN-BOLT11, LN-BOL12, ContractAdress, UniverseAddress)
func isAddress(token_list []*PreParsedToken, scrversion *ScriptVersion) (bool, *PreparedToken, []*PreParsedToken, error) {
	// Es wird geprüft ob mindestens 1 Eintrag vorhanden ist
	if len(token_list) < 1 {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}

	// Es wird geprüft ob es sich um einen Stringlosen Text handelt
	if token_list[0].Type != P_TEXT {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}

	// Es wird geprüft ob es sich um eine Bitcoin Adresse handelt
	is_btc, err := address.IsBtcAddress(token_list[0].Value)
	if is_btc || err != nil {
		// Der Rückgabewert wird gebaut
		adr_value := new(PreparedAddress)
		adr_value.Value = token_list[0].Value
		adr_value.Type = ADR_TYPE_BITCOIN
		return_value := new(PreparedToken)
		return_value.AddressValue = adr_value
		return_value.StartLine = token_list[0].Line
		return_value.StrLineEnd = token_list[0].StrLineEnd
		return_value.StartPos = token_list[0].StartPos
		return_value.EndPos = token_list[0].EndPos
		return_value.Type = &PR_ADDRESS

		// Die Daten werden zurückgegeben
		return true, return_value, token_list[1:], nil
	}

	// Es wird geprüft ob es sich um einen Account handelt
	is_vm_adr, err := address.IsAccountAddress(token_list[0].Value)
	if is_vm_adr {
		// Der Rückgabewert wird gebaut
		adr_value := new(PreparedAddress)
		adr_value.Value = token_list[0].Value
		adr_value.Type = ADR_TYPE_UNIVERSE_ADDRESS
		return_value := new(PreparedToken)
		return_value.AddressValue = adr_value
		return_value.StartLine = token_list[0].Line
		return_value.StrLineEnd = token_list[0].StrLineEnd
		return_value.StartPos = token_list[0].StartPos
		return_value.EndPos = token_list[0].EndPos
		return_value.Type = &PR_ADDRESS

		// Die Daten werden zurückgegeben
		return true, return_value, token_list[1:], nil
	}

	// Es wird geprüft ob es sich um eine Contract Adresse handelt
	is_contract_adr, err := address.IsContractAddress(token_list[0].Value)
	if is_contract_adr {
		// Der Rückgabewert wird gebaut
		adr_value := new(PreparedAddress)
		adr_value.Value = token_list[0].Value
		adr_value.Type = ADR_TYPE_CONTRACT_ADDRESS
		return_value := new(PreparedToken)
		return_value.AddressValue = adr_value
		return_value.StartLine = token_list[0].Line
		return_value.StrLineEnd = token_list[0].StrLineEnd
		return_value.StartPos = token_list[0].StartPos
		return_value.EndPos = token_list[0].EndPos
		return_value.Type = &PR_ADDRESS

		// Die Daten werden zurückgegeben
		return true, return_value, token_list[1:], nil
	}

	// Es wird geprüft ob es sich um eine Universe Adresse handelt
	is_universe_adr, err := address.IsUniverseAddress(token_list[0].Value)
	if is_universe_adr {
		// Der Rückgabewert wird gebaut
		adr_value := new(PreparedAddress)
		adr_value.Value = token_list[0].Value
		adr_value.Type = ADR_TYPE_UNIVERSE_ADDRESS
		return_value := new(PreparedToken)
		return_value.AddressValue = adr_value
		return_value.StartLine = token_list[0].Line
		return_value.StrLineEnd = token_list[0].StrLineEnd
		return_value.StartPos = token_list[0].StartPos
		return_value.EndPos = token_list[0].EndPos
		return_value.Type = &PR_ADDRESS

		// Die Daten werden zurückgegeben
		return true, return_value, token_list[1:], nil
	}

	// Es handelt sich nicht um eine Adresse
	return false, &PreparedToken{}, []*PreParsedToken{}, nil
}

// Diese Funktion gibt an ob es sich um einen Datentypen handelt
func isDatatype(token_list []*PreParsedToken, scrversion *ScriptVersion) (bool, *PreparedToken, []*PreParsedToken, error) {
	// Es wird geprüft ob mindestens 1 Eintrag vorhanden ist
	if len(token_list) < 1 {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}

	// Es wird geprüft ob es sich um einen Stringlosen Text handelt
	if token_list[0].Type != P_TEXT {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}

	// Es wird geprüft ob es sich um ein Schlüsselwort handelt
	is_keyword, adr_v, err := isStringADatatype(token_list[0].Value)
	if err != nil {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}
	if !is_keyword {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}

	// Es wird ein neuer Rückgabewert erzeugt
	new_return := new(PreparedToken)
	new_return.Type = &PR_DATATYPE
	new_return.DatatypeValue = adr_v
	new_return.StartPos = token_list[0].StartPos
	new_return.EndPos = token_list[0].EndPos
	new_return.StartLine = token_list[0].Line
	new_return.StrLineEnd = token_list[0].StrLineEnd

	// Es wird geprüft ob es sich um ein Keyword handelt
	return true, new_return, token_list[1:], nil
}

// Diese Funktion gibt an ob es sich um eine Sonderregel handelt
func isRule(token_list []*PreParsedToken, scrversion *ScriptVersion) (bool, *PreparedToken, []*PreParsedToken, error) {
	// Es wird geprüft ob mindestens 3 Einträge vorhanden sind
	if len(token_list) < 3 {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}

	// Es wird geprüft ob es sich um ein Symbol handelt
	if token_list[0].Type != P_SYMBOL {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}

	// Es wird geprüft ob es sich um ein Hashtag handelt
	if token_list[0].SymbolValue != NumberSignSymbol {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}

	// Es wird geprüft ob als nächstes ein String vorhanden ist
	if token_list[1].Type != P_TEXT {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}

	// Es wird geprüft ob es sich um ein zulässige Schlüsselwort handelt
	if token_list[1].Value != "enable" && token_list[1].Value != "disable" {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}

	// Es wird geprüft ob als nächstes ein String vorhanden ist
	if token_list[2].Type != P_TEXT {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}

	// Das Regelobjekt wird gebaut
	rules_object := new(PreparedRules)
	rules_object.RuleName = token_list[2].Value
	rules_object.RuleState = token_list[1].Value == "enable"

	// Das Rückantwortobjekt wird gebaut
	return_value := new(PreparedToken)
	return_value.RuleValue = rules_object
	return_value.StartPos = token_list[0].StartPos
	return_value.StartLine = token_list[0].Line
	return_value.StrLineEnd = token_list[2].StrLineEnd
	return_value.EndPos = token_list[2].EndPos

	// Die Daten werden zurückgegeben
	return true, return_value, token_list[3:], nil
}

// Diese Funktion gibt an ob es sich um ein Integer oder ein Float handelt
func isIntegerOrFloat(token_list []*PreParsedToken, scrversion *ScriptVersion) (bool, *PreparedToken, []*PreParsedToken, error) {
	// Es wird geprüft ob sich mindestens 1 Eintrag auf dem Stack befindet
	if len(token_list) < 1 {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}

	// Es wird geprüft ob es sich um eine Nummer handelt
	if token_list[0].Type != P_NUMBER {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}

	// Es wird versucht den Wert einzulesen
	int_value, err := new(big.Int).SetString(token_list[0].Value, 0)
	if !err {
		return false, &PreparedToken{}, []*PreParsedToken{}, fmt.Errorf("isIntegerOrFloat: invalid number input: " + token_list[0].Value + " " + string(token_list[0].Type))
	}

	// Es wird ein neuer Floatwert erzeugt
	new_return := new(PreparedToken)
	new_return.Type = &PR_INTEGER
	new_return.IntegerValue = &PreparedInteger{IsNegative: token_list[0].IsMinus, Value: int_value}
	new_return.StartPos = token_list[0].StartPos
	new_return.EndPos = token_list[0].EndPos
	new_return.StartLine = token_list[0].Line
	new_return.StrLineEnd = token_list[0].StrLineEnd

	// Es wird geprüft ob noch 2 Weitere Einträge auf dem Stack vorhanden sind
	if len(token_list[1:]) < 2 {
		return true, new_return, token_list[1:], nil
	}

	// Es wird geprüft ob als nächstes ein Punkt auf dem Stack vorhanden ist
	if token_list[1].Type != P_SYMBOL {
		return true, new_return, token_list[1:], nil
	}
	if token_list[1].SymbolValue != PeriodSymbol {
		return true, new_return, token_list[1:], nil
	}

	// Es wird geprüft ob als nächstes eine nicht Negative Nummer auf dem Stack vorhanden ist
	if token_list[2].Type != P_NUMBER {
		return true, new_return, token_list[1:], nil
	}
	if token_list[2].IsMinus {
		return true, new_return, token_list[1:], nil
	}

	// Es wird versucht die zweite Nummer einzulesen
	_, err = new(big.Int).SetString(token_list[0].Value, 0)
	if !err {
		return false, &PreparedToken{}, []*PreParsedToken{}, fmt.Errorf("isIntegerOrFloat: invalid number input: " + token_list[0].Value + " " + string(token_list[0].Type))
	}

	// Es wird ein neuer Float String ersellt
	float_string := token_list[0].Value + "." + token_list[2].Value

	// Es wird versucht den Wert einzulesen
	float_value, err := new(big.Float).SetString(float_string)
	if !err {
		return false, &PreparedToken{}, []*PreParsedToken{}, fmt.Errorf("isIntegerOrFloat: invalid number input: " + float_string)
	}

	// Der Integerwert wird entfernt und durch ein Floatwert setzt
	new_return.FloatValue = &PreparedFloat{IsNegative: token_list[0].IsMinus, Value: float_value}
	new_return.IntegerValue = nil

	// Die Endposition sowie die letzte Zeile werden geändert
	new_return.EndPos = token_list[2].EndPos
	new_return.StrLineEnd = token_list[2].StrLineEnd

	// Der Typ wird angepasst
	new_return.Type = &PR_FLOAT

	// Die Daten werden zurückgegeben
	return true, new_return, token_list[3:], nil
}

// Diese Funktion wandelt alle Restlichen Werte um
func nextRead(token_list []*PreParsedToken, scrversion *ScriptVersion) (bool, *PreparedToken, []*PreParsedToken, error) {
	// Es wird geprüft ob mindestens 1 Eintrag vorhanden ist
	if len(token_list) < 1 {
		return false, &PreparedToken{}, []*PreParsedToken{}, nil
	}

	// Es wird geprüft ob es sich um einen Strinlosenwert handelt
	if token_list[0].Type == P_TEXT {
		// Es wird ein Prepaiertes String Objekt erstellt
		preparated_str := PreparedText(token_list[0].Value)

		// Das Ergbnissobjekt wird gebaut
		resolve_erg := new(PreparedToken)
		resolve_erg.Type = &PR_TEXT
		resolve_erg.TextValue = &preparated_str
		resolve_erg.StartPos = token_list[0].StartPos
		resolve_erg.EndPos = token_list[0].EndPos
		resolve_erg.StartLine = token_list[0].Line
		resolve_erg.StrLineEnd = token_list[0].StrLineEnd

		// Die Daten werden zurückgegeben
		return true, resolve_erg, token_list[1:], nil
	}

	// Es wird geprüft ob es sich um einen Textstring handelt
	if token_list[0].Type == P_TEXT_STR {
		// Es wird ein Prepaiertes String Objekt erstellt
		preparated_str := PreparedString(token_list[0].Value)

		// Das Ergbnissobjekt wird gebaut
		resolve_erg := new(PreparedToken)
		resolve_erg.Type = &PR_TEXT_STR
		resolve_erg.StringValue = &preparated_str
		resolve_erg.StartPos = token_list[0].StartPos
		resolve_erg.EndPos = token_list[0].EndPos
		resolve_erg.StartLine = token_list[0].Line
		resolve_erg.StrLineEnd = token_list[0].StrLineEnd

		// Die Daten werden zurückgegeben
		return true, resolve_erg, token_list[1:], nil
	}

	// Es wird geprüft ob es sich um ein Symbol handelt
	if token_list[0].Type == P_SYMBOL {
		// Das Ergbnissobjekt wird gebaut
		resolve_erg := new(PreparedToken)
		resolve_erg.Type = &PR_SYMBOL
		resolve_erg.SymbolValue = &token_list[0].SymbolValue
		resolve_erg.StartPos = token_list[0].StartPos
		resolve_erg.EndPos = token_list[0].EndPos
		resolve_erg.StartLine = token_list[0].Line

		// Die Daten werden zurückgegeben
		return true, resolve_erg, token_list[1:], nil
	}

	// Es wird geprüft ob es sich um ein Kommentar handelt
	if token_list[0].Type == P_COMMENT {
		// Es wird ein Prepaiertes String Objekt erstellt
		preparated_str := new(PreparedComment)
		preparated_str.Multiline = token_list[0].StrLineEnd < token_list[0].Line
		preparated_str.Value = token_list[0].Value

		// Das Ergbnissobjekt wird gebaut
		resolve_erg := new(PreparedToken)
		resolve_erg.Type = &PR_COMMENT
		resolve_erg.CommentValue = preparated_str
		resolve_erg.StartPos = token_list[0].StartPos
		resolve_erg.EndPos = token_list[0].EndPos
		resolve_erg.StartLine = token_list[0].Line
		resolve_erg.StrLineEnd = token_list[0].StrLineEnd

		// Die Daten werden zurückgegeben
		return true, resolve_erg, token_list[1:], nil
	}

	// Es handelt sich um einen unbekannten Typen
	return false, &PreparedToken{}, []*PreParsedToken{}, fmt.Errorf("nextObject: Unkown data type: " + string(token_list[0].Type))
}

// Diese Funktion wird ausgeführt um eine PreParsedTokenList auf den
func PreparePreParsedTokenList(token_list []*PreParsedToken) (*PreparedUnparsedScript, error) {
	// Es wird geprüft ob Mindestens 10 Items auf dem Stack liegen
	if len(token_list) < 10 {
		return &PreparedUnparsedScript{}, fmt.Errorf("PreparePreParsedTokenList: invalid script")
	}

	// Es wird geprüft ob das Skript mit der Angabe der Sprache ound der Version beginnt
	retrived_tokens, c_version, current_token_stack, err := getScriptLangAndVersion(token_list)
	if err != nil {
		return &PreparedUnparsedScript{}, fmt.Errorf("Invalid script")
	}

	// Es wird ein neues Skriptobjekt erstellt
	current_scirpt_obj := new(PreparedUnparsedScript)
	current_scirpt_obj.PreparatedTokens = retrived_tokens
	current_scirpt_obj.LanguageSpeficationVersion = c_version

	// Die Einzelnen Einträge werden abgerufen und Spiziell geprüft
	for len(current_token_stack) > 0 {
		// Es wird geprüft ob es sich um einen Datentypen handelt
		is_rule, new_token, new_list, err := isRule(current_token_stack, c_version)
		if err != nil {
			return &PreparedUnparsedScript{}, fmt.Errorf("PreparePreParsedTokenList: " + err.Error())
		}
		if is_rule {
			current_scirpt_obj.PreparatedTokens = append(current_scirpt_obj.PreparatedTokens, new_token)
			current_token_stack = new_list
			continue
		}

		// Es wird geprüft ob es sich um ein Keyword handelt
		is_keyworkd, new_token, new_list, err := isKeyword(current_token_stack, c_version)
		if err != nil {
			return &PreparedUnparsedScript{}, fmt.Errorf("PreparePreParsedTokenList: " + err.Error())
		}
		if is_keyworkd {
			current_token_stack = new_list
			current_scirpt_obj.PreparatedTokens = append(current_scirpt_obj.PreparatedTokens, new_token)
			continue
		}

		// Es wird geprüft ob es sich um eine Adresse handelt
		is_address, new_token, new_list, err := isAddress(current_token_stack, c_version)
		if err != nil {
			return &PreparedUnparsedScript{}, fmt.Errorf("PreparePreParsedTokenList: " + err.Error())
		}
		if is_address {
			current_token_stack = new_list
			current_scirpt_obj.PreparatedTokens = append(current_scirpt_obj.PreparatedTokens, new_token)
			continue
		}

		// Es wird geprüft ob es sich um einen Datentypen handelt
		is_dtype, new_token, new_list, err := isDatatype(current_token_stack, c_version)
		if err != nil {
			return &PreparedUnparsedScript{}, fmt.Errorf("PreparePreParsedTokenList: " + err.Error())
		}
		if is_dtype {
			current_token_stack = new_list
			current_scirpt_obj.PreparatedTokens = append(current_scirpt_obj.PreparatedTokens, new_token)
			continue
		}

		// Es wird geprüft ob es sich um einen Integer oder einer Float handelt
		is_float, new_token, new_list, err := isIntegerOrFloat(current_token_stack, c_version)
		if err != nil {
			return &PreparedUnparsedScript{}, fmt.Errorf("PreparePreParsedTokenList: " + err.Error())
		}
		if is_float {
			current_token_stack = new_list
			current_scirpt_obj.PreparatedTokens = append(current_scirpt_obj.PreparatedTokens, new_token)
			continue
		}

		// Die Restlichen Objekte werden umgewandelt
		is_obj, new_token, new_list, err := nextRead(current_token_stack, c_version)
		if err != nil {
			return &PreparedUnparsedScript{}, fmt.Errorf("PreparePreParsedTokenList: " + err.Error())
		}
		if is_obj {
			current_token_stack = new_list
			current_scirpt_obj.PreparatedTokens = append(current_scirpt_obj.PreparatedTokens, new_token)
			continue
		}

		// Es handelt sich um ein ungültiges Skript
		return &PreparedUnparsedScript{}, fmt.Errorf("PreparePreParsedTokenList: It is an invalid script")
	}

	// Die Prepaierten Skriptendaten werden aufgegeben
	return current_scirpt_obj, nil
}
