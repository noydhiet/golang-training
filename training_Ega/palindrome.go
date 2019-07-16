package main

import (
	"fmt"
	"bufio"
	"os"
)

func isPalindrome(input string) string {
	for i:=0; i < len(input)/2; i++ {
		if (input[i] != input[len(input) -1 - i] ) {
			return "BUKAN PERMUTASI"
		}
	}
	return "PERMUTASI"
}

func main() {
	consoleReader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your word\n")
    input, _ := consoleReader.ReadString('\n')
    fmt.Println("Your input is : ", input)
    fmt.Println(isPalindrome(input))

}