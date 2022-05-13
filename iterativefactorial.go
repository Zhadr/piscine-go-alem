package piscine

func IterativeFactorial(nb int) int {
	count := 1
	if nb < 0 || nb > 21 {
		return 0
	} else {
		for i := 1; i <= nb; i++ {
			count *= int(i)
		}
		return count
	}
}
