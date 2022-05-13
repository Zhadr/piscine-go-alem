package piscine

func Abort(a, b, c, d, e int) int {
	args := []int{a, b, c, d, e}
	count := 5
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
	return args[2]
}
