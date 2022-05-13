package piscine

func LastRune(s string) rune {
	str := []rune(s)
	last := str[len(str)-1]
	return last
}
