package piscine

func MakeRange(min, max int) []int {
	var newarg []int

	if min >= max {
		return newarg
	}

	args := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		args[i] = i + min
	}
	return args
}
