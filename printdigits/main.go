package main

import "github.com/01-edu/z01"

func main() {
	a := '0'
	b := '9'
	for i := a; i <= b; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
