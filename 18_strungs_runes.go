package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const sRus = "ЫЙхю"
	const sThai = "สวัสดี"

	fmt.Println("len (ru):", len(sRus))
	fmt.Println("len (thai):", len(sThai))

	fmt.Print(sRus + ": ")
	for i := 0; i < len(sRus); i++ {
		fmt.Printf("%x ", sRus[i])
	}
	fmt.Println()

	fmt.Println("Rune count (Rus):", utf8.RuneCountInString(sRus))
	fmt.Println("Rune count (Thai):", utf8.RuneCountInString(sThai))

	for idx, runeValue := range sThai {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(sThai); i += w {
		runeValue, width := utf8.DecodeRuneInString(sThai[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
