package main

import "fmt"

func main() {

	var usia uint8 = 17
	//  to keep real age until decimal
	var usiaAsli = 12345678
	usiaBohong := 19
	var isExist bool = true

	// displaying string with many quotation with backstick
	var backSticks = `It's "Ega". Alamat Bojonggede. Kerja di Celerates.`

	fmt.Println(usia)
	fmt.Println(usiaAsli)
	fmt.Println(usiaBohong)
	fmt.Println(isExist)
	fmt.Println(backSticks)
}
