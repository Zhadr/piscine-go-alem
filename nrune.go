package piscine

func NRune(s string, n int) rune {
	str := []rune(s)
	if len(str) < n || n < 1 {
		return 0
	}
	getRune := rune(str[n-1])
	return getRune
}
