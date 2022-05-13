package piscine

func Compact(ptr *[]string) int {
	var arg []string
	count := 0
	for _, i := range *ptr {
		if i != "" {
			arg = append(arg, i)
			count++
		}
	}
	*ptr = arg
	return count
}
