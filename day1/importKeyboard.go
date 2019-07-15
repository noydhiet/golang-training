package main

import (
	"fmt"
)

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter year : ")
// 	year, _ := reader.ReadString('\n')
// 	// _ : menangkap error : gak dipakai pun gpp
// 	fmt.Println(year)
// 	year2, err := strconv.Atoi(year)
// if err = nil {
// 	fmt.Println(err)
// } else {
// 	fmt.Println(year2)
// }

// }

func main() {
	var i int
	fmt.Print("Masukkan i : ")
	fmt.Scanf("%d", &i)

	k := 0
	for j := 1; j <= i; j++ {
		if i%j == 0 {
			k++
		}
	}
	if k == 2 {
		fmt.Println("Bilangan prima")
	} else {
		fmt.Println("Bukan prima")
	}

}
