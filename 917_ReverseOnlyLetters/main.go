package main

import "fmt"

func main() {
	fmt.Println(reverseOnlyLetters("ab-cd"))
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj") == "j-Ih-gfE-dCba")
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!") == "Qedo1ct-eeLg=ntse-T!")
}

func reverseOnlyLetters(s string) string {
	b := []byte(s)
	i, lastIdx := 0, len(b)-1
	for i < lastIdx {
		if isLetter(b[i]) {
			for !isLetter(b[lastIdx]) {
				lastIdx--
			}
			b[i], b[lastIdx] = b[lastIdx], b[i]
			lastIdx--
		}
		i++
	}

	return string(b)
}

func isLetter(r uint8) bool {
	return (r >= 65 && r <= 90) || (r >= 97 && r <= 122)
}
