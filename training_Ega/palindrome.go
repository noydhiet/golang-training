package main

import "fmt"

func main() {

	var word string
	fmt.Println("Enter your word: ")

	fmt.Scanf("%d", &word)
	fmt.Println(isPalindrome(word))

}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func isPalindrome(word string) string {
	if word == reverse(word) {
		return "PERMUTASI"
	}
	return "BUKAN PERMUTASI"
}
