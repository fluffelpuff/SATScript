package script

// Diese Funktion gibt an ob der Name bereits verwendet wird
func (obj *ParsedScript) IsDeclarated(name string) bool {
	for _, item := range obj.Items {
		if *item.ItemType == PS_ITEM_FUNCTION_DECLARATION {
			if item.FunctionDeclaration.Name == name {
				return true
			}
		} else if *item.ItemType == PS_ITEM_VAR_DECLARATION {
			if item.VarName == name {
				return true
			}
		}
	}
	return false
}

// Registriert eine neue Funktion
func (obj *ParsedScript) RegisterNewFunction(pfnc *ParsedFunction) bool {
	// Es wird geprüft ob die Funktion bereits Registriert wurde
	if obj.IsDeclarated(pfnc.Name) {
		return false
	}

	// Das Item wird erzeugt
	cs_item := new(ParsedScriptItem)
	cs_item.ItemType = &PS_ITEM_FUNCTION_DECLARATION

	// Der Eintrag wird hinzugefügt
	obj.Items = append(obj.Items, cs_item)

	// Die Funktion wurde erfolgreich registriert
	return true
}

// Registriert eine neue Variale
func (obj *ParsedScript) RegisterNewGlobalVariable() bool {
	return false
}

// Erzeugt ein neues ParsedScript Objekt
func NewParsedScript(version *ScriptVersion) *ParsedScript {
	news := new(ParsedScript)
	return news
}
