package main

import "fmt"

func main() {

	var peserta [2]string
	peserta[0] = "Hafidh"
	peserta[1] = "Adhi"

	usia := [3]int{1, 2, 3}

	fmt.Println(peserta)
	fmt.Println(usia)

	for i := 0; i < 3; i++ {
		fmt.Println(usia[i])
	}

}
