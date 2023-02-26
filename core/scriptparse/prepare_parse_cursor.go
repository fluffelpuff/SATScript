package scriptparse

type PreparedUnparsedScriptCursor struct {
	PreparedUnparsedScriptObject *PreparedUnparsedScript
	CurrentHight                 uint
}

// Gibt das Aktuelle Item aus
func (obj *PreparedUnparsedScriptCursor) GetCurrentItem() *PreparedToken {
	return obj.PreparedUnparsedScriptObject.PreparatedTokens[obj.CurrentHight]
}

// Dreht die Aktuelle Höhe um eins Nach oben
func (obj *PreparedUnparsedScriptCursor) Next() {
	obj.CurrentHight++
}

// Gibt an ob sich der Cursor am ende Befindet
func (obj *PreparedUnparsedScriptCursor) IsEnd() bool {
	return len(obj.PreparedUnparsedScriptObject.PreparatedTokens) == int(obj.CurrentHight)
}

// Gibt das Aktuelle Item zurück und erhöhrt um eins
func (obj *PreparedUnparsedScriptCursor) GetCurrentItemANext() *PreparedToken {
	reval := obj.PreparedUnparsedScriptObject.PreparatedTokens[obj.CurrentHight]
	obj.CurrentHight++
	return reval
}

// Diese Funktion überträgt die Aktuelle Höhe zurück
func (obj *PreparedUnparsedScriptCursor) FinallyPushBackHight() {
	obj.PreparedUnparsedScriptObject.currentHight = obj.CurrentHight
}
