package main

import (
	"fmt"
)

func main() {
	var kata string
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter huruf :")
	// name, _ := reader.ReadString('\n')

	fmt.Println("Masukan kata :")
	fmt.Scanf("%s", &kata)

	fmt.Println(faktorial(kata))
}

func faktorial(kata string) string {
	//chek
	for i := 0; i < len(kata)/2; i++ {
		if kata[i] != kata[len(kata)-1-i] {
			return "Bukan Permutasi"
		}
	}
	return "Permutasi"
}
