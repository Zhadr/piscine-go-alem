package piscine

func AppendRange(min, max int) []int {
	args := []int(nil)
	if min > max {
		return nil
	}
	for i := min; i < max; i++ {
		args = append(args, i)
	}
	return args
}
