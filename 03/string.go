package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := strings.ToUpper("nik")
	fmt.Println(s1)

	s2 := strings.ToLower("Nik")
	fmt.Println(s2)

	s3 := strings.Title("Nik is a developer")
	fmt.Println(s3)

	s4 := strings.EqualFold("Swe den", "Swe deN")
	fmt.Println(s4)

	s5 := strings.HasPrefix("Swe den", "Sw")
	fmt.Println(s5)

	s6 := strings.HasSuffix("Swe den", "den")
	fmt.Println(s6)

	s7 := strings.Count("Sweden", "e")
	fmt.Println(s7)

	s8 := strings.Index("Sweden", "d")
	fmt.Println(s8)

	s9 := strings.Repeat("Sweden", 5)
	fmt.Println(s9)

	s10 := strings.TrimSpace("    Sweden     ")
	fmt.Println(s10, len(s10))

	s11 := strings.TrimLeft("    Sweden", " ")
	fmt.Println(s11, len(s11))

	s12 := strings.TrimRight("Sweden     ", " ")
	fmt.Println(s12, len(s12))

	fmt.Println(strings.Compare("nik", "nIK"))

	s13 := strings.Fields("Welcome to Go programming")
	fmt.Println(s13, len(s13))

	s14 := strings.Join(s13, "_")
	fmt.Println(s14)

	s15 := strings.Split(s14, "_")
	fmt.Println(s15, len(s15))
	s16 := strings.SplitAfter(s14, "_")
	fmt.Println(s16, len(s16))

	s17 := strings.Replace(s14, "_", "*", -1)
	fmt.Println(s17)

	tf := func(c rune) bool {
		return unicode.IsDigit(c)
	}

	s18 := strings.TrimFunc("123NikhilGOlang789", tf)
	fmt.Println(s18)

}
