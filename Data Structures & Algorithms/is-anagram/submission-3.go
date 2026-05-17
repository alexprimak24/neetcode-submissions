func isAnagram(s string, t string) bool {
	freq := make(map[rune]int)
	freq1 := make(map[rune]int)
	for _,char := range s{
		freq[char] += 1 
	}
	for _, char1 := range t{
		freq1[char1] += 1 
	}

	var longest string
	if len(s) > len(t) {
		longest = s
	} else {
		longest = t
	}
	for _,ch := range longest{
		v, ok := freq[ch]
		if !ok {
			return false
		}
		v1, ok := freq1[ch]
		if !ok {
			return false
		}
		if v1 != v {
			return false
		}
	}

	return true
}
