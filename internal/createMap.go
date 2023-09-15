package internal

import (
	"os"
	"strings"
)

func CreateMap(format string) map[rune][]string {
	data, _ := os.ReadFile("internal/data/" + format + ".txt")
	res := strings.Split(strings.ReplaceAll(string(data), "\\n", "\n"), "\n")
	Map := make(map[rune][]string)

	count := 1
	for i := ' '; i <= '~'; i++ {
		Map[i] = res[count : count+8]
		count += 9
	}
	return Map
}
