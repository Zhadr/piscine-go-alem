package piscine

func ConcatParams(args []string) string {
	word := ""
	for _, val := range args {
		word += val + "\n"
	}
	return word[:len(word)-1]
}
