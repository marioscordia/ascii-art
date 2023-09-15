package internal

import (
	"ascii-art/option/align"
	"ascii-art/option/check"
	"ascii-art/option/color"
	"ascii-art/option/fs"
	"ascii-art/option/output"
	"os"
	"strings"
)

func IsOption(args []string, Map map[rune][]string) {
	words := fs.OptionFs(Map)

	for i, option := range args {
		if strings.HasPrefix(option, "--color=") {
			words = color.OptionColor(Map, words, i)
			break
		}
	}
	for i, option := range args {
		if strings.HasPrefix(option, "--align=") {
			words = align.OptionAlign(words, i) // исправить
			break
		}
	}
	for i, option := range args {
		if strings.HasPrefix(option, "--output=") {
			output.OptionOutput(words, i)
			os.Exit(0)
		}
	}
	Print(words)
}

func IsFormat(args []string) string {
	format := ""

	if args[len(args)-1] == "standard" || args[len(args)-1] == "shadow" || args[len(args)-1] == "thinkertoy" {
		format = args[len(args)-1]
	} else if len(args) == 1 {
		format = "standard"
	} else if strings.HasPrefix(args[0], "--color=") && check.IsArguments(args) && len(args) == 3 {
		format = "standard"
	} else if len(args) <= 5 && check.IsArguments(args) {
		format = "standard"
	} else {
		os.Exit(0)
	}
	return format
}
