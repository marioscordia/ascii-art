package align

import (
	"ascii-art/option/check"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func OptionAlign(words [][]string, n int) [][]string {
	mapArgs := check.IsWords()

	length := Length(words)
	right, left, sep := Align(length, n)

	for _, arg := range mapArgs {
		if len(right) != 0 {
			words = Append(words, "right", right)
		}
		for _, r := range arg {
			if r == ' ' && len(sep) != 0 {
				words = Append(words, "sep", sep)
			}
		}
		if len(left) != 0 {
			words = Append(words, "left", left)
		}
	}
	return words
}

func Append(words [][]string, align, str string) [][]string {
	mapArgs := check.IsWords()
	var arr [][]string
	var word []string
	for j := 0; j < 8; j++ {
		word = append(word, str)
	}
	if align == "left" {
		for i := 0; i < len(mapArgs); i++ {
			words = append(words, word)
		}
	}
	if align == "right" {
		for i := 0; i < len(mapArgs); i++ {
			arr = append(arr, word)
			for _, w := range words {
				arr = append(arr, w)
			}
			words = arr
		}
	}
	if align == "sep" {
		for _, arg := range mapArgs {
			for j, r := range arg {
				if r == ' ' {
					words[j] = word
				}
			}
			arr = append(arr, word)
		}
	}
	return words
}

func Align(lenght, n int) (string, string, string) {
	var left, right, sep int
	word := strings.Join(check.IsWords(), " ")
	space := strings.Count(word, " ")
	align := strings.TrimPrefix(os.Args[n+1], "--align=")
	if align == "left" {
		right = ConsoleSize() - lenght
	} else if align == "right" {
		left = ConsoleSize() - lenght
	} else if align == "center" {
		right = (ConsoleSize() - lenght) / 2
		left = (ConsoleSize() - lenght) / 2
	} else if align == "justify" {
		if strings.Count(word, " ") != 0 {
			sep = (ConsoleSize() + 6*space - lenght) / (space)
		}
	} else {
		fmt.Println("please check the followings:\n1.type of file to be written to\n2.word characters")
		os.Exit(0)
	}
	leftS, rightS, sepS := Spaces(left, right, sep)

	return leftS, rightS, sepS
}

func Spaces(left, right, sep int) (string, string, string) {
	var leftS, rightS, sepS string
	for ; left > 0; left-- {
		leftS += " "
	}
	for ; right > 0; right-- {
		rightS += " "
	}
	for ; sep > 0; sep-- {
		sepS += " "

	}
	return leftS, rightS, sepS
}

func Length(words [][]string) int {
	var total int
	for _, word := range words {
		total += len(word[0])
	}
	return total
}

func ConsoleSize() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()

	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")

	width, _ := strconv.Atoi(sArr[1])

	return width
}
