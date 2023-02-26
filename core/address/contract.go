package address

import (
	"strings"

	"github.com/tnakagawa/goref/bech32m"
)

func IsContractAddress(strv string) (bool, error) {
	if strings.HasPrefix(strv, "contract1") {
		hrp, decoded, _, err := bech32m.Decode(strv)
		if err != nil {
			return false, nil
		}
		if hrp != "contract" {
			return false, nil
		}
		if len(decoded) != 66 {
			return false, nil
		}
		return true, nil
	}
	return false, nil
}
