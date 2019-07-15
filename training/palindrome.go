package main

import "fmt"

func main() {
	//set masukan huruf yang akan dimasukan ke palindrome
	var word string
	fmt.Println("Masukan huruf : ")
	fmt.Scanf("%d", &word)
	fmt.Println(isPalindrome(word))
}

// fungsi untuk membagi dua huruf dan membadingkan huruf awal dan di ujung
func isPalindrome(word string) bool {
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}
	return true
}
/*
Ini cara kedua dengan metode reverse tetapi kurang efisien karena banyak looping

func isPalindrome(word string) bool {
	if word == reverse(word) {
		return true
	}
	return false
}
*/
