package main

import (
	"fmt"
)

func main()  {
	var polindStr string
	fmt.Printf("Enter String : ")

	fmt.Scanf("%s", &polindStr)
	fmt.Println(isPalindrome(polindStr))
}

func reverse(s string) string {
	c := []rune(s)
	for i, j := 0, len(c)-1; i < len(c)/2; i, j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]
	}
	return string(c)
}

func isPalindrome(polindStr string) string {
	if  polindStr == reverse(polindStr) {
		return "PERMUTASI"
	}
	return "BUKAN PERMUTASI"
}