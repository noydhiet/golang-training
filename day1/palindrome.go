package main

import "fmt"

func main() {
	// fmt.Printf(reverse("abcdefg"))
	var i string
	fmt.Print("Masukkan kata (lowercase tanpa spasi) : ")
	fmt.Scanf("%s", &i)
	// fmt.Printf(reverse(i))
	if i == reverse(i) {
		fmt.Print(i)
		fmt.Print(" = ")
		fmt.Printf(reverse(i))
	} else {
		fmt.Print(i)
		fmt.Print(" != ")
		fmt.Printf(reverse(i))
	}
}

func reverse(s string) string {
	words := []rune(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return string(words)
}
