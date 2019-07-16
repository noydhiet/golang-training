package main

import "fmt"

// func main()  {
	
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter Year :")
// 	year, _ := reader.ReadString('\n')
// 	fmt.Println(year)
// 	year2,err
// }

func main()  {
	
	var i int
	fmt.Println("enter year : ")

	_, err := fmt.Scanf("%d", &i)

	if err != nil {
		fmt.Println(err)
	}

	k := 0

	for j := 1; j <= i; j++ {
		if i%j == 0 {
			k++
		}
	}

	if k == 2{
		fmt.Println("Bilangan Prima")
	} else {
		fmt.Println("Bukan Bilangan Prima")
	}

	

		
	


}