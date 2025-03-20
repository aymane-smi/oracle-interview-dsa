package main

import (
	"fmt"
	"strings"
	"unicode"
)

func modify(s string) string {
	var b strings.Builder
	b.WriteByte(byte(s[0]))
	for i := 1; i < len(s); i++ {
		if unicode.IsUpper(rune(s[0])) {
			b.WriteByte(byte(unicode.ToUpper(rune(s[i]))))
		} else if unicode.IsLower(rune(s[0])) {
			b.WriteByte(byte(unicode.ToLower(rune(s[i]))))
		}
	}
	return b.String()
}

func main() {
	fmt.Print(modify(("Abcd")))
}
