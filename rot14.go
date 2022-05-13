package piscine

func Rot14(s string) string {
	str := []rune(s)
	converted := str
	for i := 0; i < len(str); i++ {
		if (str[i] >= 65 && str[i] <= 76) || (str[i] >= 97 && str[i] <= 108) {
			converted[i] = str[i] + 14
		} else if (str[i] >= 77 && str[i] <= 90) || (str[i] >= 109 && str[i] <= 122) {
			converted[i] = str[i] - 12
		}
	}
	return string(converted)
}
