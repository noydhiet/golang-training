package main

import "fmt"

func isPalindrome(str string) string  {
	
	for x:= 0; x < len(str)/2; x++{
		
		if str[x]!=  str[len(str)-1-x]{

			return "Bukan palindrom"
		}
	}
	return "Baru palindrom"
}

func main()  {
	
	var str string
	fmt.Println("Write input: ")
	fmt.Scanf("%s", &str)
	fmt.Printf(isPalindrome(str))

}
	
