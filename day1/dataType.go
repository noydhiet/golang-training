package main

import "fmt"

func main() {
	var usia uint8 = 17
	// to keep real age until decimal
	var usiaAsli = 20.5
	usiaBohong := 20

	isExist := true

	var backSticks = `
					Nama saya <nama>. 
					Tinggal di <alamat>. 
					Kerja di "Celerates".
					`

	fmt.Println(usia)
	fmt.Println(usiaAsli)
	fmt.Println(usiaBohong)
	fmt.Println(isExist)
	fmt.Println(backSticks)
}
