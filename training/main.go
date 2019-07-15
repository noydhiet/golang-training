package main 

import "fmt"
/*
	bla
	bla
	this is multiline comment
*/


func main() {
	
	//name variable 
	firstName := "firdha"   // := dipake untuk ngunci nilai dan tipe data nya
	var lastName string = "maulidya"
	address := new(string)

	//This line produces Hello World result
	//fmt.Println("Hello ", firstName, lastName, "!!!")  //hasilmya make spasi
	//fmt.Println("Hello "+ firstName+ lastName+ "!!!")  //hasilnya ga pake spasi
	fmt.Printf("Hello %s %s! \n", firstName, lastName)  //tentuin tipe data pake %s, pake Printf
	fmt.Printf(*address)
	//fmt.Printf(address)	
	//end of code


}