package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	src := os.Args
	for i := 1; i < len(src); i++ {
		for j := 0; j < len(src[i]); j++ {
			z01.PrintRune(rune(src[i][j]))
		}
		z01.PrintRune('\n')

	}
}
