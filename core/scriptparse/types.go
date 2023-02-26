package scriptparse

import "math/big"

// Wird vom Lexer verwendet
type TokenDatatype string

// Wird von Tokenisierten Skripten verwendet welche als nächstes Prepariert werden sollen
type PreParsedTokenDataType string

// Wird von Preparierten aber noch nicht Geparsten Skript Objekten verwendet
type PreparedScriptTokenDataType string

// Stellt ein Schlüsselwort dar
type PreparedKeyword string

// Stellt einen Datentypen dar
type PreparedDatatype string

// Stellt einen String dar
type PreparedString string

// Stellt einen Text dar
type PreparedText string

// Stellt einen Adresstypen dar
type AddressType string

// Stellt eine Adresse dar
type PreparedAddress struct {
	Type  AddressType
	Value string
}

// Stellt eine Regel dar
type PreparedRules struct {
	RuleState bool
	RuleName  string
}

// Stellt eine Nummer dar
type PreparedInteger struct {
	IsNegative bool
	Value      *big.Int
}

// Stellt einen Float dar
type PreparedFloat struct {
	IsNegative bool
	Value      *big.Float
}

// Stellt ein Kommentar dar
type PreparedComment struct {
	Multiline bool
	Value     string
}

// Gibt ein Argument für eine Funktionsdeklaration an
type ParsedFunctionArgument struct {
	Type PreparedScriptTokenDataType
	Name string
}

// Wird verwendet wenn eine Funktionsdeklaration darzustellen
type ParsedFunction struct {
	Arguments []*ParsedFunctionArgument
	Name      string
	IsPublic  bool
}

// Gibt ein Geparstes Objekt an
type ParsedObject struct {
}
