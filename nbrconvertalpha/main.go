package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	src := os.Args[1:]
	tr := 0
	for _, c := range src {
		curr := 0
		for _, v := range c {
			curr = curr*10 + int(v-'0')
		}
		if c == "--upper" {
			tr = 1
		}
		if c == "--upper" || curr > 26 {
			z01.PrintRune(' ')
			continue
		}
		z01.PrintRune(rune((curr + 'a' - 1 - tr*32)))
	}
	z01.PrintRune('\n')
}
