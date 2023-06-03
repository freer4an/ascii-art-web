package steps

import (
	"strings"
	"unicode"
)

func PrintSamples(text string, sample map[rune]Ascii_art) (string, int) {
	if !isASCII(text) {
		return "", 0
	}
	words := strings.Split(text, "\n")
	if strings.HasPrefix(text, "\n") {
		words = words[1:]
	}

	str := ""

	for _, word := range words {
		if word == "" {
			str += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, ch := range word {
				for key, value := range sample {
					if ch == key {
						startline := sample[key].width*i + i
						endline := startline + sample[key].width
						for _, line := range value.letter[startline:endline] {
							if line != '\n' {
								str += string(line)
							}
						}
					}
				}
			}
			str += "\n"
		}
	}
	return str[:len(str)-1], 1
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
