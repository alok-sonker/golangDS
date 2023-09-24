package main

import "fmt"

// There are 2 type of string literals
//
//	1- Interpreted
//	2- Raw
func stringLiterals() {
	// Interpredted string lietrals
	// this string should be in (``)
	interpretedStr := `My name is ALOK Sonker and 
	I have 6 years of experience in software engineering.
	Lets contribute in following area
		AI
	ML
				blockchain
									distributed system
		etc.
	`

	fmt.Println(interpretedStr)
	rawStr := "Hello Github!! its me"
	fmt.Println(rawStr)
}
