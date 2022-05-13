package piscine

func BasicAtoi2(s string) int {
	var count, sum int
	for _, val := range []rune(s) {
		if val >= '0' && val <= '9' {
			count = 0
			for i := '0'; i < val; i++ {
				count++
			}
			sum = sum*10 + count
		} else if !(val >= '0' && val <= '9') {
			return 0
		}
	}
	return sum
}
