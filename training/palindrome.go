package main

import (
	"fmt"
	"strings"
)

func main() {

	var word string

	fmt.Println("Enter your word :\t")
	fmt.Scanf("%s", &word)

	if word == reverse(word) {
		fmt.Println(word, "=>", strings.ToUpper(word))
		fmt.Println("PERMUTASI")
	} else {
		fmt.Println(word, "=>", strings.ToUpper(word))
		fmt.Println("BUKAN PERMUTASI")
	}
}

// using reverse to variable word
func reverse(s string) string {
	words := []rune(s)
	for x, y := 0, len(words)-1; x < y; x, y = x+1, y-1 {
		words[x], words[y] = words[y], words[x]
	}

	return string(words)
}
