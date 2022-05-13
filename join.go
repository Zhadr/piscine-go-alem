package piscine

func Join(strs []string, sep string) string {
	result := strs[0]
	for i := range strs {
		if i != len(strs)-1 {
			result = result + sep + strs[i+1]
		}
	}
	return result
}
