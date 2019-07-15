package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter String :")
	dataStr, _ := reader.ReadString('\n')
	fmt.Println(isPalindrome(dataStr))
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func isPalindrome(dataStr string) string {
	if dataStr == reverse(dataStr) {
		return "BUKAN PERMUTASI"
	}
	return "PERMUTASI"
}
