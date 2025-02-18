package sub

import "strings"

func CheckPun(s string) bool {
	const prefixes = ".,!?;:"
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, string(prefix)) {
			return true
		}
	}
	return false
}
