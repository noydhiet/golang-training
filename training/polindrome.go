package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your word : ")
	word, _ := reader.ReadString('\n')
	fmt.Println(word)

	var word string = "30"
	var pal bool = true

	if word != "30" {
		fmt.Println("Jangan lebih dari 30 karakter")
	}
}
