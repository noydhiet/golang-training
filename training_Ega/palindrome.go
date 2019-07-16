package main

import (
	"fmt"
	"bufio"
	"os"
)

func isPalindrome(word string) string {
	for i:=0; i < len(word)/2; i++ {
		if (word[i] != word[len(word) -1 - i] ) {
			return "BUKAN PERMUTASI"
		}
	}
	return "PERMUTASI"
}

func main() {
	consoleReader := bufio.NewReader(os.Stdin)
    	fmt.Print("Enter your word\n")
    	word, _ := consoleReader.ReadString('\n')
    	fmt.Println("Your input is : ", word)
    	fmt.Println(isPalindrome(word))

}
