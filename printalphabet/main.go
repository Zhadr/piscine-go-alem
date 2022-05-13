package main

import "github.com/01-edu/z01"

func main() {
	a := 'a'
	b := 'z'
	for i := a; i <= b; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
