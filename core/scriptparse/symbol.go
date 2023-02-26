package scriptparse

type SymbolToken string

var (
	// ₿ Zeichen
	BtcSymbol SymbolToken = "₿"

	// @ Zeichen
	AtSymbol SymbolToken = "@"

	// ! Zeichen
	ExclamationMarkSymbol SymbolToken = "!"

	// # Zeichen
	NumberSignSymbol SymbolToken = "#"

	// $ Zeichen
	DollarSignSymbol SymbolToken = "$"

	// % Zeichen
	PercentSignSymbol SymbolToken = "%"

	// ^ Zeichen
	CaretSymbol SymbolToken = "^"

	// & Zeichen
	AmpersandSymbol SymbolToken = "&"

	// * Zeichen
	AsteriskSymbol SymbolToken = "*"

	// ( Zeichen
	LestParenthesisSymbol SymbolToken = "("

	// ) Zeichen
	RightParenthesisSymbol SymbolToken = ")"

	// _ Zeichen
	UnderscoreSymbol SymbolToken = "_"

	// + Zeichen
	PlusSignSymbol SymbolToken = "+"

	// , Zeichen
	CommaSymbol SymbolToken = ","

	// . Zeichen
	PeriodSymbol SymbolToken = "."

	// / Zeichen
	SlashSymol SymbolToken = "/"

	// | Zeichen
	VerticalBarSymbol SymbolToken = "|"

	// \ Zeichen
	BackslashSymbol SymbolToken = "\\"

	// Tabulator
	TabSymbol SymbolToken = "\\"

	// ` Zeichen
	ApostropheSymbol SymbolToken = "`"

	// - Zeichen
	MinusSignSymbol SymbolToken = "-"

	// = Zeichen
	EqualToSignSymbol SymbolToken = "="

	// < Zeichen
	OpeningAngleBracketSymbol SymbolToken = "<"

	// > Zeichen
	ClosingAngleBracket SymbolToken = ">"

	// ? Zeichen
	QuestionMarkSymbol SymbolToken = "?"

	// { Zeichen
	LeftBraceSymbol SymbolToken = "{"

	// } Zeichen
	RightBraceSymbol SymbolToken = "}"

	// [ Zeichen
	LeftBracketSymbol SymbolToken = "["

	// ] Zeichen
	RightBracketSymbol SymbolToken = "]"

	// : Zeichen
	ColonSymbol SymbolToken = ":"

	// " Zeichen
	QuotationMarkSymbol SymbolToken = "\""

	// ' Zeichen
	MarkSymbol SymbolToken = "'"

	// ; Zeichen
	SemicolonSymbol SymbolToken = ";"
)

// Gibt alle Verfügbaren Symole als Liste aus
func getSymbolList() []*SymbolToken {
	return []*SymbolToken{
		&BtcSymbol,
		&AtSymbol,
		&ExclamationMarkSymbol,
		&NumberSignSymbol,
		&DollarSignSymbol,
		&PercentSignSymbol,
		&CaretSymbol,
		&AmpersandSymbol,
		&AsteriskSymbol,
		&LestParenthesisSymbol,
		&RightParenthesisSymbol,
		&UnderscoreSymbol,
		&PlusSignSymbol,
		&CommaSymbol,
		&PeriodSymbol,
		&SlashSymol,
		&VerticalBarSymbol,
		&BackslashSymbol,
		&ApostropheSymbol,
		&MinusSignSymbol,
		&EqualToSignSymbol,
		&OpeningAngleBracketSymbol,
		&ClosingAngleBracket,
		&QuestionMarkSymbol,
		&LeftBraceSymbol,
		&RightBraceSymbol,
		&LeftBracketSymbol,
		&RightBracketSymbol,
		&ColonSymbol,
		&QuotationMarkSymbol,
		&SemicolonSymbol,
		&TabSymbol,
		&MarkSymbol,
	}
}
