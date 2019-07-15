package main

import (
	"fmt"
	"sort"
)
func main()  {
	
	urutanAcak := []int{10,3,2,1,8,7}
	
	sort.Ints(urutanAcak)
	fmt.Println(urutanAcak)

}

/*
func main(){
	primes := [6]int{10,3,2,1,8,7}
	length := len(primes)
	for i :=0; i < length; i++{
		for j:= i +1; j < length; j++{
			if primes [j] < primes[i]{
			tmp := primes[i]
			primes[i] = primes[j]
			primes[j] = tmp
		
		}
		
	}
}
*/