package main

func isAnagramSlow(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var firstStrMap map[byte]int = map[byte]int{}
	for i := 0; i < len(s); i++ {
		_, exist := firstStrMap[s[i]]
		if exist {
			firstStrMap[s[i]] += 1
		} else {
			firstStrMap[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		if firstStrMap[t[i]] == 0 {
			return false
		}
		_, exist := firstStrMap[t[i]]
		if exist {
			firstStrMap[t[i]] -= 1
		} else {
			return false
		}
	}

	return true

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var arr [26]int
	n := len(s)

	for i := 0; i < n; i++ {
		arr[s[i]-'a']++
	}

	for i := 0; i < n; i++ {
		arr[t[i]-'a']--
		if arr[t[i]-'a'] == 0 {
			return false
		}
	}
	return true
}
