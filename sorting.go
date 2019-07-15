package main

import "fmt"

func main() {

	primes := []int{10, 3, 2, 1, 8, 7}
	length := len(primes)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if primes[j] < primes[i] {
				tmp := primes[i]
				primes[i] = primes[j]
				primes[j] = tmp
			}
		}

	}

	fmt.Println(primes)
}
