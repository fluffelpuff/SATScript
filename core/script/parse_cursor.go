package script

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

// Gibt an ob sich der Cursor am ende befindet
func (obj *PreparedUnparsedScriptCursor) IsEnd() bool {
	return len(obj.PreparedUnparsedScriptObject.PreparatedTokens) == int(obj.CurrentHight)
}

// Gibt das Aktuelle Item zurück und erhöhrt um eins
func (obj *PreparedUnparsedScriptCursor) GetCurrentItemANext() *PreparedToken {
	reval := *obj.PreparedUnparsedScriptObject.PreparatedTokens[obj.CurrentHight]
	obj.CurrentHight++
	return &reval
}

// Diese Funktion überträgt die Aktuelle Höhe zurück
func (obj *PreparedUnparsedScriptCursor) FinallyPushBackHight() {
	obj.PreparedUnparsedScriptObject.currentHight = obj.CurrentHight
}

// Diese Funktion erstellt einen SliceCursor aus einem PreparedUnparsedScriptCursor
func (obj *PreparedUnparsedScriptCursor) ToSliceCursor() *SliceBodyCursor {
	slic := new(SliceBodyCursor)
	slic.CurrentHight = obj.CurrentHight
	slic.Items = obj.PreparedUnparsedScriptObject.PreparatedTokens[obj.CurrentHight:]
	return slic
}

type SliceBodyCursor struct {
	Items        []*PreparedToken
	AbolutHight  uint
	CurrentHight uint
}

// Gibt das Aktuelle Item aus
func (obj *SliceBodyCursor) GetCurrentItem() *PreparedToken {
	return obj.Items[obj.CurrentHight]
}

// Dreht die Aktuelle Höhe um eins Nach oben
func (obj *SliceBodyCursor) Next() {
	obj.CurrentHight++
}

// Gibt an ob sich der Cursor am ende befindet
func (obj *SliceBodyCursor) IsEnd() bool {
	return len(obj.Items) == int(obj.CurrentHight)
}

// Gibt das Aktuelle Item zurück und erhöhrt um eins
func (obj *SliceBodyCursor) GetCurrentItemANext() *PreparedToken {
	reval := *obj.Items[obj.CurrentHight]
	obj.CurrentHight++
	return &reval
}

// Setzt den Cursor auf 0 zurück
func (pbj *SliceBodyCursor) Reset() {
	pbj.CurrentHight = pbj.AbolutHight
}

// Setzt die Abolut Höhe auf die Aktuelle Höhe
func (pbj *SliceBodyCursor) SetAbolut() {
	pbj.AbolutHight = pbj.CurrentHight
}

// Gibt die Gesamtgröße des Restlichen Stacks an
func (pbj *SliceBodyCursor) RestItems() int {
	return len(pbj.Items) - int(pbj.CurrentHight)
}

// Gibt den Aktuellen Slice aus
func (obj *SliceBodyCursor) GetSlice() []*PreparedToken {
	return obj.Items[obj.CurrentHight:]
}

// Überträgt die Änderungen eines Slices an den einen Cursor
func transportStateToCursor(m *SliceBodyCursor, s *PreparedUnparsedScriptCursor) bool {
	s.CurrentHight = m.CurrentHight
	return true
}

// Das gleiche wie transportStateToCursor nur dass die Änderungen danach an den Master zurück übergeben werden
func transportStateToCursorAndFinallyPushBackHight(m *SliceBodyCursor, s *PreparedUnparsedScriptCursor) bool {
	if !transportStateToCursor(m, s) {
		return false
	}
	s.FinallyPushBackHight()
	return true
}
