package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	src := os.Args[1:]
	Sort(src)
	for i := 0; i < len(src); i++ {

		for j := 0; j < len(src[i]); j++ {
			z01.PrintRune(rune(src[i][j]))
		}
		z01.PrintRune('\n')

	}
}

func Sort(table []string) {
	len := 0
	for i := range table {
		i++
		len++
	}
	i := 1
	for i < len {
		if table[i-1] > table[i] {
			temp := table[i-1]
			table[i-1] = table[i]
			table[i] = temp
			i = 1
		} else {
			i++
		}
	}
}
