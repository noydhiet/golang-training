package main

import "fmt"

func main() {

	var peserta [4]string
	peserta[0] = "Inggrid"
	peserta[1] = "Melanika"
	peserta[2] = "Jeremi"
	peserta[3] = "Deardo"

	usia := [3]int{1, 2, 3}

	// show var usia with looping
	for i := 0; i < 3; i++ {
		fmt.Println(usia[i])
	}

	fmt.Println(peserta[0], peserta[1], peserta[2], peserta[3])
	fmt.Println("Jumlah index : \t", len(peserta)) // count the index in single array
}
