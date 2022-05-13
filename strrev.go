package piscine

func StrRev(s string) string {
	count := ""
	for _, ch := range s {
		count = string(ch) + count
	}
	return count
}
