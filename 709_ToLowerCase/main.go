package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(toLowerCase("Hello"))
	fmt.Println(toLowerCase("here"))
	fmt.Println(toLowerCase("LOVELY"))
}

func toLowerCase(str string) string {
	b := bytes.Buffer{}
	for _, r := range str {
		if r >= 'A' && r <= 'Z' {
			r += 32
		}
		b.WriteString(string(r))
	}
	return b.String()
}
