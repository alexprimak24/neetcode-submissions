func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	freq, freq1 := make(map[rune]int),make(map[rune]int)
	for i, char := range s{
		freq[char]++
		freq1[rune(t[i])]++ 
	}

	for k,v := range freq{
		if freq1[k] != v {
			return false
		}
	}

	return true
}
