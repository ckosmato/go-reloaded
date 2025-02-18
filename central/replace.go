package central

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"reloaded/central/sub"
)

func Replace(r io.Reader, w io.Writer) error {
	// handling of cases to be replaced
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		words := strings.Fields(line)

		for i := len(words) - 1; i >= 0; i-- {
			word := words[i]
			if word[len(word)-1] == ')' && word[0] != '(' {
				words[i-1] += word
				words = append(words[:i], words[i+1:]...)
			}
			if sub.CheckBrackets(word) {
				keyword, num := sub.RemoveBrackets(word)
				startIdx := i - num
				if startIdx < 0 {
					startIdx = 0
				}

				switch keyword {
				case "hex":
					for j := startIdx; j < i; j++ {
						if j < len(words) {
							hexValue := words[j]
							if decValue, err := strconv.ParseInt(hexValue, 16, 64); err == nil {
								words[j] = strconv.FormatInt(decValue, 10)
							}
						}
					}
				case "bin":
					for j := startIdx; j < i; j++ {
						if j < len(words) {
							binValue := words[j]
							if decValue, err := strconv.ParseInt(binValue, 2, 64); err == nil {
								words[j] = strconv.FormatInt(decValue, 10)
							}
						}
					}
				case "up":
					for j := startIdx; j < i; j++ {
						if j < len(words) {
							words[j] = strings.ToUpper(words[j])
						}
					}
				case "low":
					for j := startIdx; j < i; j++ {
						if j < len(words) {
							words[j] = strings.ToLower(words[j])
						}
					}
				case "cap":
					for j := startIdx; j < i; j++ {
						if j < len(words) {
							words[j] = sub.Capitalize(words[j])
						}
					}
				}

				words = append(words[:i], words[i+1:]...)
			}

		}

		for i := len(words) - 1; i >= 0; i-- {
			word := words[i]
			if sub.CheckPun(word) && i > 0 {
				if len(word) > 1 && !sub.CheckPun(string(word[1])) {
					words[i-1] += string(word[0])
					words[i] = strings.Trim(words[i], string(word[0]))
				} else {
					words[i-1] += word
					words = append(words[:i], words[i+1:]...)
				}
			}
		}
		for i := len(words) - 1; i >= 0; i-- {
			word := words[i]
			add := "n"
			if word == "a" || word == "A" {
				if sub.CheckVowels(words[i+1]) {
					words[i] += add
				}
			}
		}

		firsttime := true
		for i := len(words) - 1; i >= 0; i-- {
			word := words[i]
			if firsttime && word == "'" && i > 0 {
				words[i-1] += "'"
				words = append(words[:i], words[i+1:]...)
				firsttime = false
			}
			if !firsttime && word == "'" && i < len(words)-1 {
				words[i+1] = word + words[i+1]
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}

		newLine := strings.Join(words, " ")
		newLine = strings.ReplaceAll(newLine, "  ", " ")
		if _, err := io.WriteString(w, newLine+"\n"); err != nil {
			return err
		}
	}
	return sc.Err()
}
