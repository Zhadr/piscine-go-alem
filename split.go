package piscine

func Split(s, sep string) []string {
	for i := 0; i < len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			s = s[:i] + " " + s[i+len(sep):]
		}
	}
	return SplitWhiteSpaces(s)
}
