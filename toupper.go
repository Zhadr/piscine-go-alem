package piscine

func ToUpper(s string) string {
	str := []rune(s)
	str2 := str
	for i := 0; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			str2[i] = str[i] - 32
		}
	}
	return string(str2)
}
