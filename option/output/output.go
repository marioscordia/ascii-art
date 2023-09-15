package output

import (
	"ascii-art/option/check"
	"os"
	"strings"
)

func OptionOutput(mapArgs [][]string, n int) {
	words := ChangeType(mapArgs)
	Output(words, n)
}

func ChangeType(mapArgs [][]string) []string {
	var words []string
	var word string
	var total, count int

	for n, v := range os.Args[1:] {
		if strings.HasPrefix(v, "--align=") {
			align := strings.TrimPrefix(os.Args[n+1], "--align=")
			if align != "justify" {
				count = 1
			}
		}
	}
	for _, arg := range check.IsWords() {
		for i := 0; i < 8; i++ {
			word = ""
			for j := 0; j < len(arg)+count; j++ {
				word += mapArgs[j+total][i]
			}
			words = append(words, word)
		}
		total += len(arg)
	}

	return words
}

func Output(words []string, n int) {
	fileName := check.IsFile(os.Args[n+1])
	file, _ := os.Create(fileName)
	for i := 0; i < len(words); i++ {
		file.WriteString(words[i])
		file.WriteString("\n")
	}
}
