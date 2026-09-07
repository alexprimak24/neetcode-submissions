func isValid(s string) bool {
  pairs := map[rune]rune{
	'}':'{',
	']':'[',
	')':'(',
  }

  temp := []rune{}	
  for _, char := range s {
	counterChar, ok := pairs[char]
	// opposite parenheses
	if !ok {
		temp = append(temp, char)
	} else {
		// all fine, remove last because we found pair
		if len(temp) > 0 && temp[len(temp)-1] == counterChar {
			temp = temp[:len(temp)-1]
		} else {
			return false
		}
	}
  }

  if len(temp) == 0 {
	return true
  }

  return false
}
