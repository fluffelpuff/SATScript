package script

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
	Type PreparedDatatype
	Name string
}

// Stellt eine Operation dar
type ParsedFunctionOperation struct {
}

// Wird verwendet wenn eine Funktionsdeklaration darzustellen
type ParsedFunction struct {
	Operations []*ParsedFunctionOperation
	Arguments  []*ParsedFunctionArgument
	ReturnType []*PreparedDatatype
	Name       string
	IsPublic   bool
}

// Wird verwendet um eine Variable zu Deklarieren
type ParsedVariableDeclaration struct {
}

// Gibt den verwendeten Item Type an
type ParsedScriptItemType string

// Speichert alle Skript einträge ab
type ParsedScriptItem struct {
	ItemType     *ParsedScriptItemType
	CommentValue *PreparedText
	FloatValue   *big.Float
	VarName      string
	IntValue     *big.Int
	StringValue  string
	BoolValue    bool
}

// Gibt ein Geparstes Objekt an
type ParsedScript struct {
	DeclaratedFunctions []*string
	DeclaratedVariabels []*string
	Items               []*ParsedScriptItem
}

// Speichert die Variablen und die Funktionen ab welche verfügbar sind
type ParsedScriptDefines struct {
	// Speichert die Funktionen ab welche Deklariert wurden
	DeclaratedFunctions []*ParsedFunction

	// Speichert die Variablen ab welche Deklariert wurden
	DeclaratedVariabels []*ParsedVariableDeclaration
}
