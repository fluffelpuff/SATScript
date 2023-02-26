package address

import (
	"strings"

	"github.com/tnakagawa/goref/bech32m"
)

// Es wird geprüft ob es sich um eine VmAdresse handelt
func IsAccountAddress(strv string) (bool, error) {
	if strings.HasPrefix(strv, "ssca1") {
		hrp, decoded, _, err := bech32m.Decode(strv)
		if err != nil {
			return false, nil
		}
		if hrp != "ssca" {
			return false, nil
		}
		if len(decoded) != 55 {
			return false, nil
		}
		return true, nil
	}
	return false, nil
}
