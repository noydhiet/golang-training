package main

import "fmt"

func main() {

	var usia uint8 = 17
	// to keep real age until decimal
	var usiaAsli = 20.5
	usiaBohong := 25.90

	var isExist bool = true

	var backSticks = `Nama saya "Muhammad Hanif".
	Alamat Bogor.
	Kerja di "Celerates".`

	fmt.Println(usia)
	fmt.Println(usiaAsli)
	fmt.Println(usiaBohong)
	fmt.Println(isExist)
	fmt.Println(backSticks)

}
