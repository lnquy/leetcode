package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}

// O(n)
func isValid(s string) bool {
	ps := make([]rune, 0)
	for _, r := range s {
		l := len(ps)
		switch r {
		case '(','[', '{':
			ps = append(ps, r)
		case ')':
			if l == 0 || ps[l-1] != '(' {
				return false
			}
			ps = ps[:l-1]
		case ']':
			if l == 0 || ps[l-1] != '[' {
				return false
			}
			ps = ps[:l-1]
		case '}':
			if l == 0 || ps[l-1] != '{' {
				return false
			}
			ps = ps[:l-1]
		}
	}
	return len(ps) == 0
}
