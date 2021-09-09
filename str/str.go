package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	s := "«ABC»"

	/*
		var s string
		s = "<<ABC>>"

		var s string = "<<ABC>>"
	*/

	fmt.Println((s))
	fmt.Println("len", len(s))
	fmt.Printf("%T\n", s[0]) //uint8, one single byte

	for _, r := range s {
		// get one rune at a time if you do a for loop on a string
		fmt.Printf("%T\n", r) // int32, rune
		break
	}

	// have to use utf8 package to get actual length of 5 here so it counts the unicode correctly
	fmt.Println("len (runes):", utf8.RuneCountInString(s))

	//path := "c:\to\nowhere" //compliation will not work cause of the backslash with quotes. MUST USE BACKTICKS BABY
	// \t is TAB and \n is new line
	path := `c:\to\nowhere`
	fmt.Println(path)

	/* compliation error, unknown escape sequence
	s1 := "\d"
	fmt.Println(s1)
	*/

	// raw string is when you use back slash for a reason not for like a direction, ie \n being new line vs \n

	// use back ticks to do multi line strings
	poem := `The road goes even on and on,
	Down from the road where it began...`

	fmt.Println(poem)

	// for strings methods, look up string package -- no methods like JS
	// same for regular expressions
	fmt.Printf("|%s|\n", strings.TrimSpace("  hello "))
}
