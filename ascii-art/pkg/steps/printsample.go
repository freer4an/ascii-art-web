package steps

import (
	"strings"
)

func PrintSamples(text string, contents string, sample map[rune]Ascii_art) string {
	words := strings.Split(text, "\n")
	if strings.HasPrefix(text, "\n") {
		words = words[1:]
	}

	str := ""

	for _, word := range words {
		if !strings.ContainsAny(word, contents) {
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
	return str
}
