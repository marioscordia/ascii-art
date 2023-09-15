package main

import (
	"ascii-art/internal"
	"os"
)

func main() {
	format := internal.IsFormat(os.Args[1:]) // 1.0
	Map := internal.CreateMap(format)        // 1.1
	internal.IsOption(os.Args[1:], Map)
}
