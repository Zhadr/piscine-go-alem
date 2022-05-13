package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	temp := 0
	if n/10 == 0 {
		temp = n
		if temp < 0 {
			z01.PrintRune('-')
		}
	} else {
		temp = n % 10
		PrintNbr(n / 10)
	}
	if temp < 0 {
		temp = -temp
	}
	z01.PrintRune(rune(temp) + 48)
}
