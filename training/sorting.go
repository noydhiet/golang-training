package main

import "fmt"

func main()  {
	
	urutanAcak := [7]int{10, 4, 3, 2, 1, 8, 7}
	length := len(urutanAcak)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if urutanAcak[j] < urutanAcak[i] {
				tmp := urutanAcak[i]
				urutanAcak[i] = urutanAcak[j]
				urutanAcak[j] = tmp
			}
		}
	}
	fmt.Println("Sorting number :", urutanAcak)
}