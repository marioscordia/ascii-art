package fs

import (
	"ascii-art/option/check"
)

func OptionFs(Map map[rune][]string) [][]string {
	mapArgs := check.IsWords()
	words := [][]string{}

	var total int
	for _, arg := range mapArgs {
		for _, r := range arg {
			if r != '\n' {
				words = append(words, Map[r][0:8])
				total++
			}
		}
	}
	return words
}
