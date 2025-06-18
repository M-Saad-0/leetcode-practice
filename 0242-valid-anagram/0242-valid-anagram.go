func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var arr [26]int
	n := len(s)

	for i := 0; i < n; i++ {
		arr[s[i] - 'a']++
	}

	for i := 0; i < n; i++ {
		arr[t[i] - 'a']--
		if arr[t[i] - 'a'] < 0 {
			return false
		}
	}
	return true
}