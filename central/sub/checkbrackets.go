package sub

import "strings"

func CheckBrackets(s string) bool {
	return strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") && len(s) > 2
}
