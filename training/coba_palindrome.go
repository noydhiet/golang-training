package main

import (
	"fmt"
)

func palindromo(word string) bool {

	
	for i:=0; i < len(word)/2; i++ {

		if (word[i] != word[len(word) -1 - i] ) {
			return false
		}
				
	}
	
	return true
}

func main() {
	result := palindromo("A B C D B A")
	fmt.Println(result)
}