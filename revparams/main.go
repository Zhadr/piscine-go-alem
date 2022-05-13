package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	src := os.Args
	for i := len(src) - 1; i > 0; i-- {
		for j := 0; j < len(src[i]); j++ {
			z01.PrintRune(rune(src[i][j]))
		}
		z01.PrintRune('\n')

	}
}
