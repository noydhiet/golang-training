package main

import "fmt"

// func main()  {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Printf("Enter your city : ")
// 	city, _ := reader.ReadString('\n')
// 	fmt.Printf("You live in " + city)
// }

func main() {
	var i string

	fmt.Println("Enter an integer value : ")

	_, err := fmt.Scanf("%d", &i)

	if err != nil {
		fmt.Println(err)
	}
	k := 0
	for j := 1; j <= i; j++ {
		if i%j == 0 {
			k++
		}
	}
	if k == 2 {
		fmt.Println("Bilangan Prima")
	} else {
		fmt.Println("Bukan Bilangan Prima")
	}
	
	fmt.Println(i)
}
