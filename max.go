package piscine

func Max(a []int) int {
	length := 0
	for i := range a {
		i++
		length++
	}
	i := 1
	for i < length {
		if a[i-1] > a[i] {
			temp := a[i-1]
			a[i-1] = a[i]
			a[i] = temp
			i = 1
		} else {
			i++
		}
	}
	return a[length-1]
}
