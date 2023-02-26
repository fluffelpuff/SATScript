package address

import "strings"

func IsUniverseAddress(strv string) (bool, error) {
	if strings.HasPrefix(strv, "universe2p") {

	} else if strings.HasPrefix(strv, "uivlink1") {

	}
	return false, nil
}
