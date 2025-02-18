package sub

import "strings"

func CheckVowels(s string) bool {
	const vowels = "aeiouhAEIOUH"
	for _, vowel := range vowels {
		if strings.HasPrefix(s, string(vowel)) {
			return true
		}
	}
	return false
}
