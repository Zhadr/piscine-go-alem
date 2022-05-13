package piscine

func Map(f func(int) bool, a []int) []bool {
	len := len(a)
	boolean := make([]bool, len)
	for i := range a {
		boolean[i] = f(a[i])
	}
	return boolean
}
