package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	isPalindrome := true

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your Name : ")
	inputString, _ := reader.ReadString('\n')

	inputArray := strings.Fields(inputString)

	for i := 0; i < len(inputArray)/2; i++ {
		if inputArray[i] != inputArray[len(inputArray)-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println("PALINDROME !!!")
	} else {
		fmt.Println("BUKAN PALINDROME ...")
	}

}
