package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	src := []rune(os.Args[0])
	for i := 2; i < len(src); i++ {
		z01.PrintRune(rune(src[i]))
	}
	z01.PrintRune('\n')
}
