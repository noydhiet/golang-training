package main
import "fmt"

func palindrome(word string) string {
	
	//palindrome checker
	for i :=0; i < len(word)/2; i++ {

		if (word[i] != word[len(word) -1 - i] ) {
			return "This is not palindrome"
			//fmt.Println("This is not palindrome")
		}
			
	}
	return "This is palindrome"
	//fmt.Println("This is palindrome")
}


func main() {
	// variabel untuk input
	var word string
	//input a word
	fmt.Println("Enter input: ")

	// hasil input
	fmt.Scanf("%s", &word)

	// hasil output polindrome
	fmt.Printf(palindrome(word))

}

