package script

// Gibt an ob es sich um eine bekannte Funktion handelt
func (obj *ParsedScriptDefines) IsADefinedFunction(fname string) bool {
	return true
}

// Gibt an ob es sich um eine bekannte Variable handelt
func (obj *ParsedScriptDefines) IsADefinedVariable(varName string) bool {
	return true
}

// Gibt den Datentypen einer Variable zurück
func (obj *ParsedScriptDefines) GetVariableDataType(varName string) *PreparedDatatype {
	return &DATATYPE_INT
}

// Gibt die Returndatentypen einer Funktion an
func (obj *ParsedScriptDefines) GetFunctionReturnTypes(fName string) []*PreparedDatatype {
	return []*PreparedDatatype{&DATATYPE_INT}
}

// Gibt die Parameter einer Funktion zurück
func (obj *ParsedScriptDefines) GetFunctionParameter(fName string) []*ParsedFunctionArgument {
	return []*ParsedFunctionArgument{}
}
