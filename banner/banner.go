package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G😁", 6)

	s := "G😁"
	fmt.Println("len:", len(s))
	// code point = rune ~= unicode character
	for i, r := range s {
		fmt.Println(i, r)
		if i == 0 {
			fmt.Printf("%c of type %T\n", r, r)
			// rune (int32)
		}
	}

	b := s[0]
	fmt.Printf("%c of type %T\n", b, b)
	// byte (uint8)

	x, y := 1, "1"
	fmt.Printf("x= %#v, y=%#v\n", x, y) // Use #v in debug/log

	fmt.Printf("%20s!\n", s)

	test := isPalindrome("GOG")
	fmt.Println(test)

}

func isPalindrome(s string) bool {
	rs := []rune(s) // convert s to a slice of runes
	for l, r := 0, len(rs)-1; l < r; {
		if rs[l] != rs[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func banner(text string, width int) {
	padding := (width - utf8.RuneCountInString(text)) / 2 // BUG: len is in bytes
	// padding := (width - len(text)) / 2 // BUG: len is in bytes
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
