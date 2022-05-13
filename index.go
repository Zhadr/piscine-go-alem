package piscine

func Index(s string, toFind string) int {
	len := 0
	sublen := 0
	for i := range s {
		len = i + 1
	}
	for i := range toFind {
		sublen = i + 1
	}
	a := 0
	for i := 0; i < len; i++ {
		j := 0
		a = i
		for j < sublen {
			if a < len && s[a] == toFind[j] {
				j++
				a++
			} else {
				break
			}
		}
		if j == sublen {
			return i
		}
	}
	return -1
}
