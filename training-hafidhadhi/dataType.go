package main

import "fmt"

func main() {

	var usia uint8 = 17
	// to keep real age until decimal
	var usiaAsli = 20.4
	usiaBohong := 20

	var isExist bool = true

	var backTicks = `Nama saya "Hafidh",
			Alamat Jeruk Purut,
			Kerja di "Celerates"`

	fmt.Println(usia)
	fmt.Println(usiaAsli)
	fmt.Println(usiaBohong)
	fmt.Println(isExist)
	fmt.Println(backTicks)

}
