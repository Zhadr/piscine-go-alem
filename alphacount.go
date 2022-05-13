package piscine

func AlphaCount(s string) int {
	count := 0
	str := []rune(s)
	for _, char := range str {
		if (char >= 65 && char <= 90) || (char >= 97 && char <= 122) {
			count++
		}
	}
	return count
}
