package main

import "fmt"

func main (){

	var i int 
	fmt.Println ("enter year value:")

	fmt.Scanf ("%d", &i)
	k:= 0
	for bil:= 1 ; bil <=i ; bil ++ {
		if i% bil == 0{
			k++
		}
	}

		if k == 2{
			fmt.Println("Bilangan Prima")
		} else {
			fmt.Println("Bukan Bilangan Prima")	
		}

	}

	
	