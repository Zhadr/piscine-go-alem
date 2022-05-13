package piscine

func Sqrt(nb int) int {
	count := 0
	for i := 0; i*i <= nb; i++ {
		count = i
	}
	if count*count == nb {
		return count
	} else {
		return 0
	}
}
