package internal

import (
	"ascii-art/option/output"
	"fmt"
)

func Print(words [][]string) {
	word := output.ChangeType(words)

	for i := 0; i < len(word); i++ {
		fmt.Println(word[i])
	}
}
