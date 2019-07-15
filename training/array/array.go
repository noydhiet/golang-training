package main

import "fmt"

func main() {
	var peserta [3]string
	peserta[0] = " hendy"
	peserta[1] = "Lagi"
	peserta[2] = "pusing"

	usia := [2]int{1, 2, 3}

	for i := 0; i < 3; i++ {
		fmt.Println(usia[i])
	}

	fmt.Println(peserta)
	fmt.Println(usia)
}
