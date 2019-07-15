    
package main

import "fmt"

func main() {
	var Datastr string
	fmt.Println("Enter String: ")

	fmt.Scanf("%s", &Datastr)
	fmt.Println(isPalindrome(Datastr))
}

func isPalindrome(input string) string {
	for i := 0; i < (len(input)/2)-1; i++ {
		if input[i] != input[len(input)-i] {
			return "BUKAN PERMUTASI"
		}
	}
	return "PERMUTASI"
}