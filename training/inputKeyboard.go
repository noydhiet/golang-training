package main

import "fmt"

// func main() {

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter your name : ")
// 	name, _ := reader.ReadString('\n')
// 	fmt.Println(name)
// }

/*func main() {

//untuk menampilkan inputan keyboard dalam bentuk interger
	var i int
	fmt.Println("Enter year value : ")

	fmt.Scanf("%d", &i)

	fmt.Println(i)
}*/

func main() {

	var bil int
	j := 0
	fmt.Println("Masukan Bilangan Bulat Positif : ")
	fmt.Scanf("%d", &bil) //Scanf maksudnya ialah scan and format
	//fmt.Println(bil)

	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			j++
		}
	}
	if j == 2 {
		fmt.Println("Bilangan Prima")
	} else {
		fmt.Println("Bukan Bilangan Prima")
	}
}
