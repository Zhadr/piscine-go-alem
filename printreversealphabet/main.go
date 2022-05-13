package main

import "github.com/01-edu/z01"

func main() {
	a := 'a'
	b := 'z'
	for i := b; i >= a; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
