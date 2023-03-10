package main

import (
	"fmt"
	"satscript/core/script"
)

// Printed eine Liste mit allen Contracts
func compileScript(data []string) error {
	// Der Dateipfad wird ermittelt
	file_path := data[0]
	fmt.Println(file_path)

	// Es wird geprüft ob das Skript korrekt ist
	is_correct, err := script.AnalyzeScriptFile(file_path)
	if err != nil {
		return err
	}

	// Es wird geprüft ob die Datei korrekt ist
	if !is_correct {
		return fmt.Errorf("")
	}

	// Die Contracts wurden erfolgreich abgerufen
	return nil
}

// Wird ausgeführt wenn es sich um den Script befehl handelt
func scriptCommand(data []string) error {
	switch data[0] {
	case "compile":
		return compileScript(data[1:])
	default:
		return fmt.Errorf("Unbekannter parameter " + data[0])
	}
}
