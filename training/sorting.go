package main

import "fmt"

func main()  {
	
	// urutanAcak := [6]int{10,3,2,1,8,7}
	urutanAcak := [6]int{7,8,4,3,9,0}
	
	
	for j := 1; j <= 6; j++ {
		
	
		for i := 0; i < len(urutanAcak)-1; i++ {

			if urutanAcak[i] > urutanAcak[i+1] {
				urutanAcak[i], urutanAcak[i+1] = urutanAcak[i+1], urutanAcak[i]
			}

		}

		
	}

	

	fmt.Println(urutanAcak)
	
}