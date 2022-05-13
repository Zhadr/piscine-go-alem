package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	result := 0
	for i, j := range a[:len(a)-1] {
		if f(j, a[i+1]) > 0 {
			result += 1
		}
	}
	if result == 0 || result == len(a)-1 {
		return true
	} else {
		return false
	}
}
