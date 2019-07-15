package main

import (
	"fmt"
)


urutanAcak := [6] int [6]int{10, 3, 2, 1, 8, 7}

func bubbleSort(input[6]int){
	//kita set n di angka 6 untuk increment sesuai jumlah angka di array
	n := 6
	//kita set dulu kondisi belum terswap
	swapped := false
	for swapped{
		//membandingkan angka sebelum dan sesudah di urutan array
		//jika lebih besar maka akan diukar posisinya jadi di belakang
		for i =0; i<n; i++{
			if input[i-1] > input[i]{
				input [i], input[i-1] = input[i-1], input[i]
				swapped= true
			}
		}
	}
	fmt.Println(input)
}

func main(){
	bubbleSort(urutanAcak)
}

