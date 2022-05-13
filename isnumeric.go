package piscine

func IsNumeric(s string) bool {
	num := []rune(s)
	for i := 0; i < len(num); i++ {
		if num[i] < '0' || num[i] > '9' {
			return false
		}
	}
	return true
}
