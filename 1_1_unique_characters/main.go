// Implement an algorithm to determine if a string has all unique characters.
// What if you cannot use additional data structures?
package main

import "fmt"

func main() {
	data := "z tqa"

	result := hasUniqueChars2(data)

	fmt.Printf("Characters checking result: %v", result)
}

func hasUniqueChars1(data string) bool {
	if len(data) > 256 {
		return true
	}

	resultMap := [256]bool{}

	for i := 0; i < len(data); i++ {
		if resultMap[data[i]] {
			return false
		}

		resultMap[data[i]] = true
	}

	return true
}

func hasUniqueChars2(data string) bool {
	checker := 0
	firstChar := 'a'

	for i := 0; i < len(data); i++ {
		symbol := data[i] - byte(firstChar)

		if checker&(1<<symbol) > 0 {
			return false
		}

		checker |= (1 << symbol)
	}

	return true
}
