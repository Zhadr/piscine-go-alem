package piscine

func RecursiveFactorial(nb int) int {
	// count := 1
	if nb < 0 || nb > 21 {
		return 0
	}
	if nb == 1 {
		return nb
	} else if nb == 0 {
		return 1
	} else {
		return nb * RecursiveFactorial(nb-1)
	}
}
