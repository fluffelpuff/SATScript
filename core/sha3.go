package main

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Erstellt ein Hash aus einem oderen mehren Strings
func ComputeSHA3_256_Hash(obj []string) (string, error) {
	// Erzeugt aus allen Verfügbaren Strings einen einzigen
	total_str := ""
	for _, i := range obj {
		total_str += i
	}

	// SHA-3 256 Bit Hash
	h256 := sha3.New256()
	h256.Write([]byte(total_str))

	// Hex String
	hexed_str := hex.EncodeToString(h256.Sum(nil))
	return Reverse(hexed_str), nil
}
