package main

import (
	"fmt"
)

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Input year :")
// 	year, _ := reader.ReadString('\n')
// 	year2, err := strconv.Atoi(year)

// 	if err == nil {
// 		fmt.Println(year2)
// 	} else {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(year2)
// }

func main() {
	var i int
	fmt.Println("Enter your value : ")

	fmt.Scanf("%d", &i)

	fmt.Println(i)
}
