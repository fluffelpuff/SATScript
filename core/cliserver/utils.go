package cliserver

import "satscript/core/config"

// Diese Funktion setzt die Rechte für eine Datei auf Root
func setPrimToRoot(path string) error {
	return nil
}

// Gibt an ob der Aktuelle Benutzer als Root unterwegs ist
func currentIsPrivUser(conf *config.PathConfigs) bool {
	return false
}
