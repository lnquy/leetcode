package main

import "bytes"

func reverseWords(s string) string {
	sb := []byte(s)
	result := bytes.NewBuffer(make([]byte, 0, len(s)))
	lastWordCharIdx := -1
	firstWordWritten := false

	for i := len(sb) - 1; i >= 0; i-- {
		if sb[i] == ' ' {
			if lastWordCharIdx != -1 {
				if firstWordWritten {
					_ = result.WriteByte(32) // Whitespace
				}
				_, _ = result.Write(sb[i+1 : lastWordCharIdx+1])
				lastWordCharIdx = -1
				firstWordWritten = true
			}
			continue
		}

		if lastWordCharIdx == -1 {
			lastWordCharIdx = i
		}
	}
	if lastWordCharIdx != -1 {
		if firstWordWritten {
			_ = result.WriteByte(32) // Whitespace
		}
		_, _ = result.Write(sb[0 : lastWordCharIdx+1])
	}
	return result.String()
}
