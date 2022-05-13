package piscine

func Capitalize(s string) string {
	str := []rune(s)
	len := 0
	for range str {
		len++
	}
	// str2 := str
	for i, word := range str {
		if i == 0 || !IsAlphaCheck(str[i-1]) {
			if word >= 'a' && word <= 'z' {
				str[i] = word - 32
			}
		} else {
			if word >= 'A' && word <= 'Z' {
				str[i] = word + 32
			}
		}
	}
	return string(str)
}

func IsAlphaCheck(str rune) bool {
	if (str >= 'a' && str <= 'z') || (str >= 'A' && str <= 'Z') || (str >= '0' && str <= '9') {
		return true
	}

	return false
}

// if i == 0 {
// 	if str[i] >= 'a' && str[i] <= 'z' {
// 		str[i] = str[i] - 32
// 		continue
// 	}
// 	if i == ' ' {
// 		i++
// 		if str[i] >= 'a' && str[i] <= 'z' {
// 			str[i] = str[i] - 32
// 			continue
// 		}
// 	} else {
// 		if str[i] >= 'A' && str[i] <= 'Z' {
// 			str[i] = str[i] + 32
// 		}
// 	}

// }
