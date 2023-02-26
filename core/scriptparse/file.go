package scriptparse

import (
	"errors"
	"fmt"
	"os"

	"github.com/juju/fslock"
)

// Diese Funktion versucht eine Skriptdatei zu laden und zu überprüfen ob sie Korrekt ist
func AnalyzeScriptFile(filepath string) (bool, error) {
	// Es wird geprüft ob die Datei vorhanden ist
	file, err := os.Stat(filepath)
	if errors.Is(err, os.ErrNotExist) {
		return false, err
	}

	// Die Datei wird gelockt
	lock := fslock.New(filepath)
	lockErr := lock.TryLock()
	if lockErr != nil {
		fmt.Println("falied to acquire lock > " + lockErr.Error())
		return false, lockErr
	}

	// Entsperrt die verwendete Skriptdatei
	defer lock.Unlock()

	// Es wird geprüft ob die Datei größer als 100 MB ist
	if file.Size() >= 104857600 {
	}

	// Die Skriptdatei wird eingelesen, vollständig
	dat, err := os.ReadFile(filepath)
	if err != nil {
		return false, err
	}

	// Es wird versucht das Script zu Tokenisieren
	readed_string := string(dat[:])
	extracted_token, err := TokenizationScriptString(readed_string)
	if err != nil {
		return false, err
	}

	// Das Skript wird für das Parsen Vorbereitet
	preparsed_script, err := PreParseTokenList(extracted_token)
	if err != nil {
		return false, err
	}

	// Die Vorgeparste Token Liste wird vorbereitet
	prepareted_preparsed_script, err := PreparePreParsedTokenList(preparsed_script)
	if err != nil {
		return false, err
	}

	err = ParsePreparatedScript(prepareted_preparsed_script)
	if err != nil {
		return false, err
	}

	// Das Skript wurde erfolgreich getestet
	return true, nil
}

// Diese Funktion wird verwendet
func ParseScriptFile(file_path string, output_file_path string) error {
	return nil
}
