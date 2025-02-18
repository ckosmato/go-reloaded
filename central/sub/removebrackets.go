package sub

import (
	"strconv"
	"strings"
)

func RemoveBrackets(s string) (string, int) {
	s = strings.Trim(s, "()")

	parts := strings.Split(s, ",")
	keyword := parts[0]
	num := 1
	if len(parts) > 1 {
		if count, err := strconv.Atoi(parts[1]); err == nil {
			num = count
		}
	}
	return keyword, num
}
