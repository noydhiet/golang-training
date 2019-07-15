package main

import (
	"fmt"
)

func main() {
	var dataStr string
	fmt.Println("Enter string : ")

	fmt.Scanf("%s", &dataStr)
	fmt.Println(isPalindrome(dataStr))
}

func reverse(s string) string {
	b := make([]byte, len(s))
	var j int = len(s) - 1
	for i := 0; i <= j; i++ {
		b[j-i] = s[i]
	}

	return string(b)
}

func isPalindrome(dataStr string) string {
	if dataStr == reverse(dataStr) {
		return "PERMUTASI"
	}
	return "BUKAN PERMUTASI"
}
