package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}

	currSIdx := 0
	for i := 0; i < len(t); i++ {
		if t[i] != s[currSIdx] {
			continue
		}
		currSIdx++
		if currSIdx >= len(s) {
			return true
		}
	}
	return false
}
