package piscine

func SplitWhiteSpaces(s string) []string {
	var word []rune
	var str []string

	for i, val := range s {
		if val != 32 {
			word = append(word, val)
		}
		if val == 32 || i == len(s)-1 {
			if len(word) > 0 {
				str = append(str, string(word))
				word = []rune{}
			}
		}
	}
	return str
}
