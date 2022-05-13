package piscine

func CountIf(f func(string) bool, tab []string) int {
	l := 0
	for _, index1 := range tab {
		if f(index1) {
			l++
		}
	}
	return l
}
