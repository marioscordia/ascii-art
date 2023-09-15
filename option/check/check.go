package check

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsWords() []string {
	word := ""
	for _, i := range os.Args[1:] {
		if !strings.HasPrefix(i, "--") && IsBanners(i) {
			word = i
		}
	}
	words := strings.ReplaceAll(word, "\\n", "\n")

	for _, c := range words {
		if c < '\n' || c > '~' {
			fmt.Println("Word is not in the correct form")
			os.Exit(0)
		}
	}
	if word == "" {
		os.Exit(0)
	}
	mapArgs := strings.Split(words, "\n")
	return mapArgs
}

func IsFile(file string) string {
	if strings.HasSuffix(strings.Split(file, "--output=")[1], ".txt") {
		return strings.Split(file, "--output=")[1]
	} else {
		fmt.Println("please check the followings:\n1.type of file to be written to\n2.word characters")
		os.Exit(0)
		return "error"
	}
}

func IsColor(color string) string {
	// if strings.Contains(color, "--color=") {
	// 	color = strings.Split(color, "--color=")[1]
	// } else {
	// 	fmt.Println("please check the followings:\n1.type of file to be written to\n2.word characters")
	// 	os.Exit(0)
	// }
	colorArr := []string{
		"black", "red",
		"green", "yellow", "purple",
		"magenta", "teal", "white",
	}
	for i, c := range colorArr {
		if c == color {
			return strconv.Itoa(i)
		} else if i+1 == len(colorArr) {
			return "7"
		}
	}
	return "error"
}

func IsLetters(word string) string {
	for i := 1; i < len(os.Args[1:]); i++ {
		if os.Args[i] != word && IsBanners(os.Args[i]) && !strings.Contains(os.Args[i], "--") {
			return os.Args[i]
		}
	}
	return word
}

func IsBanners(arg string) bool {
	if arg != "standard" && arg != "shadow" && arg != "thinkertoy" {
		return true
	} else {
		return false
	}
}

func IsArguments(args []string) bool {
	var total int
	for i := 0; i < len(args); i++ {
		if IsBanners(args[i]) && (!strings.HasPrefix(os.Args[i], "--color=") ||
			!strings.HasPrefix(os.Args[i], "--align=") ||
			!strings.HasPrefix(os.Args[i], "--output=")) {
			total += 1
		}
	}
	if total > 2 {
		return false
	}
	return true
}
