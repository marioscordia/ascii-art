package color

import (
	"ascii-art/option/check"
	"flag"
	"fmt"
	"strings"
)

func OptionColor(Map map[rune][]string, mapWords [][]string, n int) [][]string {
	words := strings.Join(check.IsWords(), "")
	color := flag.String("color", "", "EX: go run . --color=<color>")
	flag.Parse()
	total := check.IsColor(*color)
	var letters string

	wordsColor := mapWords

	letters = check.IsLetters(words)

	wordsColor = Color(mapWords, words, letters, total)
	return wordsColor
}

func ColorString(s, total string) string {
	return fmt.Sprintf("\033[1;3"+total+"m%s\033[0m", s)
}

func Color(mapWords [][]string, mapArgs, letters, n string) [][]string {
	wordsColor := mapWords
	for i, arg := range mapArgs {
		if strings.Contains(letters, string(arg)) {
			for j := 0; j < 8; j++ {
				wordsColor[i][j] = ColorString(wordsColor[i][j], n)
			}
		}
	}
	return wordsColor
}
