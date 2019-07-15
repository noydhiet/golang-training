package main

import "fmt"

func main() {

var Polind string 
	fmt.Println ("Masukkan kata tertentu:")
	
	fmt.Scanf ("%s", &Polind)
	fmt.Println (isPalindrome(Polind))
}

func reverse(s string) string {
	k := []rune(s)
	for i, j := 0, len(k)-1; i < len(k)/2; i, j = i+1, j-1 {
		k[i], k[j] = k[j], k[i]
	}
	return string(k)
}

func isPalindrome(Polind string) string {
	if  Polind == reverse(Polind) {
		return "PERMUTASI"
	}
	return "BUKAN PERMUTASI"
}
