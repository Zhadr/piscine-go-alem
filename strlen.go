package piscine

func StrLen(s string) int {
	count := 0
	for i := 0; i < len([]rune(s)); i++ {
		count++
	}
	return count
}
