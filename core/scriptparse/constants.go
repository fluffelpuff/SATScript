package scriptparse

import (
	"math/big"
	"strings"
)

// Speichert die Maximale größe von Zahlen ab
var MaxInt *big.Int = &big.Int{}
var MaxFloat *big.Float = &big.Float{}

// Legt die Zugelassenene Zeichen fest
var (
	CHARS     = strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZÄÜÖäüö", "")
	SYMBOLS   = strings.Split("@!§$€%&/()=?`*'_:;><,.-#+´ß \n{}[]\"|\\^\t₿", "")
	NUMBERS   = strings.Split("0123456789", "")
	TABULATOR = "\t"
	NEW_LINE  = "\n"
	SPACE     = " "
)

// Gibt die Typen für den Lexer an
var (
	SYMBOL TokenDatatype = "SYMBOL"
	NUMBER TokenDatatype = "NUMBER"
	TEXT   TokenDatatype = "TEXT"
	EMOJI  TokenDatatype = "EMOJI"
)

// Gibt die Typen für den PreParse Vorgang an
var (
	P_COMMENT  PreParsedTokenDataType = "COMMENT"
	P_TEXT_STR PreParsedTokenDataType = "STRING"
	P_SYMBOL   PreParsedTokenDataType = "SYMBOL"
	P_NUMBER   PreParsedTokenDataType = "NUMBER"
	P_TEXT     PreParsedTokenDataType = "TEXT"
)

// Gibt die Typen für den Prepariervorgang an
var (
	PR_DATATYPE PreparedScriptTokenDataType = "DATATYPE"
	PR_ADDRESS  PreparedScriptTokenDataType = "ADDRSSS"
	PR_TEXT_STR PreparedScriptTokenDataType = "STRING"
	PR_KEYWORD  PreparedScriptTokenDataType = "KEYWORD"
	PR_COMMENT  PreparedScriptTokenDataType = "COMMENT"
	PR_INTEGER  PreparedScriptTokenDataType = "INTEGER"
	PR_SYMBOL   PreparedScriptTokenDataType = "SYMBOL"
	PR_TEXT     PreparedScriptTokenDataType = "TEXT"
	PR_VERSION  PreparedScriptTokenDataType = "VERSION"
	PR_FLOAT    PreparedScriptTokenDataType = "FLOAT"
)

// Speichert alle Adresstypen ab
var (
	ADR_TYPE_BITCOIN          AddressType = AddressType("BITCOIN")
	ADR_TYPE_VM_ADDRESS       AddressType = AddressType("VM")
	ADR_TYPE_CONTRACT_ADDRESS AddressType = AddressType("CONTRACT")
	ADR_TYPE_UNIVERSE_ADDRESS AddressType = AddressType("Universe")
)

// Gibt die Möglichen Datentypen an
var (
	DATATYPE_UNIVERSE_EP_ADDRESS PreparedDatatype = PreparedDatatype("UniverseEndPointAddress")
	DATATYPE_CONTRACT_ADDRESS    PreparedDatatype = PreparedDatatype("ContractAddress")
	DATATYPE_ACCOUNT_ADDRESS     PreparedDatatype = PreparedDatatype("Account")
	DATATYPE_CHAIN_ADDRESS       PreparedDatatype = PreparedDatatype("ChainAddress")
	DATATYPE_LN11_ADDRESS        PreparedDatatype = PreparedDatatype("LN11Address")
	DATATYPE_STRING              PreparedDatatype = PreparedDatatype("String")
	DATATYPE_BOOL                PreparedDatatype = PreparedDatatype("bool")
	DATATYPE_INT                 PreparedDatatype = PreparedDatatype("int")
	DATATYPE_FLOAT               PreparedDatatype = PreparedDatatype("float")
	DATATYPE_BYTES               PreparedDatatype = PreparedDatatype("Bytes")
	DATATYPE_LIST                PreparedDatatype = PreparedDatatype("List")
	DATATYPE_JSON                PreparedDatatype = PreparedDatatype("JSON")
	DATATYPE_ARRAY               PreparedDatatype = PreparedDatatype("Array")
	DATATYPE_AMOUNT              PreparedDatatype = PreparedDatatype("Amount")
	DATATYPE_CALLABLE            PreparedDatatype = PreparedDatatype("Callable")
	DATATYPE_URL                 PreparedDatatype = PreparedDatatype("URL")
)

// Definiert die Schlüsselwörter
var (
	KEYWORD_FUNCTION PreparedKeyword = PreparedKeyword("func")
	KEYWORD_PAYABLE  PreparedKeyword = PreparedKeyword("payable")
	KEYWORD_PUBLIC   PreparedKeyword = PreparedKeyword("public")
	KEYWORD_PRIVATE  PreparedKeyword = PreparedKeyword("private")
	KEYWORD_FALSE    PreparedKeyword = PreparedKeyword("false")
	KEYWORD_ELSE     PreparedKeyword = PreparedKeyword("else")
	KEYWORD_NULL     PreparedKeyword = PreparedKeyword("null")
	KEYWORD_BREAK    PreparedKeyword = PreparedKeyword("break")
	KEYWORD_TRUE     PreparedKeyword = PreparedKeyword("true")
	KEYWORD_FINALLY  PreparedKeyword = PreparedKeyword("finally")
	KEYWORD_IS       PreparedKeyword = PreparedKeyword("is")
	KEYWORD_RETURN   PreparedKeyword = PreparedKeyword("return")
	KEYWORD_CONTINUE PreparedKeyword = PreparedKeyword("continue")
	KEYWORD_FOR      PreparedKeyword = PreparedKeyword("for")
	KEYWORD_WHILE    PreparedKeyword = PreparedKeyword("while")
	KEYWORD_ASSERT   PreparedKeyword = PreparedKeyword("assert")
	KEYWORD_ABORT    PreparedKeyword = PreparedKeyword("abort")
	KEYWORD_DELETE   PreparedKeyword = PreparedKeyword("delete")
	KEYWORD_ELSEIF   PreparedKeyword = PreparedKeyword("elseif")
	KEYWORD_IF       PreparedKeyword = PreparedKeyword("if")
	KEYWORD_IN       PreparedKeyword = PreparedKeyword("in")
	KEYWORD_ENUM     PreparedKeyword = PreparedKeyword("eum")
	KEYWORD_SWITCH   PreparedKeyword = PreparedKeyword("switch")
	KEYWORD_NEW      PreparedKeyword = PreparedKeyword("new")
	KEYWORD_DEFAULT  PreparedKeyword = PreparedKeyword("default")
	KEYWORD_EMIT     PreparedKeyword = PreparedKeyword("emit")
)

// Definiert den ParsedSkript Item eintrag an
var (
	PS_ITEM_FUNCTION_DECLARATION ParsedScriptItemType = ParsedScriptItemType("function_declaration")
	PS_ITEM_COMMENT_DECLARATION  ParsedScriptItemType = ParsedScriptItemType("comment_declaration")
)

// Gibt alle bekannten Datentypen an
var DATATYPES_SLICE = []*PreparedDatatype{
	&DATATYPE_UNIVERSE_EP_ADDRESS, &DATATYPE_CONTRACT_ADDRESS, &DATATYPE_CHAIN_ADDRESS,
	&DATATYPE_LN11_ADDRESS, &DATATYPE_ACCOUNT_ADDRESS, &DATATYPE_STRING, &DATATYPE_BOOL,
	&DATATYPE_INT, &DATATYPE_FLOAT, &DATATYPE_BYTES, &DATATYPE_LIST, &DATATYPE_JSON,
	&DATATYPE_ARRAY, &DATATYPE_AMOUNT, &DATATYPE_CALLABLE, &DATATYPE_URL,
}

// Gibt alle bekannten Keywörter an
var KEYWORD_SLICE = []*PreparedKeyword{
	&KEYWORD_FUNCTION, &KEYWORD_PAYABLE, &KEYWORD_PUBLIC,
	&KEYWORD_FALSE, &KEYWORD_ELSE, &KEYWORD_NULL, &KEYWORD_BREAK,
	&KEYWORD_TRUE, &KEYWORD_FINALLY, &KEYWORD_IS, &KEYWORD_RETURN, &KEYWORD_CONTINUE,
	&KEYWORD_FOR, &KEYWORD_WHILE, &KEYWORD_ASSERT, &KEYWORD_ABORT, &KEYWORD_DELETE,
	&KEYWORD_ELSEIF, &KEYWORD_IF, &KEYWORD_IN, &KEYWORD_ENUM, &KEYWORD_SWITCH,
	&KEYWORD_NEW, &KEYWORD_DEFAULT, &KEYWORD_EMIT,
}
