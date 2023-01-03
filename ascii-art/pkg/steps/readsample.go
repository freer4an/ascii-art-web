package steps

import (
	"bufio"
	"os"
)

type Ascii_art struct {
	letter string
	width  int
}

func GetSamples(path string) (map[rune]Ascii_art, string) {
	f, err := os.Open(path)
	if err != nil {
		return nil, ""
	}
	defer f.Close()
	test := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		test = append(test, scanner.Text())
	}
	symbols := ""
	start := 1
	samples := map[rune]Ascii_art{}
	key := ' '
	for key <= '~' {
		symbols += string(key)
		letter := ""
		width := 0
		for i := start; i < start+8; i++ {
			if i == start {
				width = len(test[i])
			}
			letter += test[i] + "\n"
		}
		samples[key] = Ascii_art{letter: letter, width: width}
		start += 9
		key++
	}
	return samples, symbols
}
