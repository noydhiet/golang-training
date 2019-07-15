package main

import (
	"fmt"
)

func main() {
	var z string //deklarasi variable z dengan tipe data string
	fmt.Printf("Masukkan Hurufnya: ")
	fmt.Scanf("%s", &z) //scan inputan keyboard harus bertipe string
	fmt.Println(Palindromebukan(z))

}

//function palindrome untuk pengecekan apakah termasuk permutasi atau bukan
func Palindromebukan(z string) string {
	if z == reverse(z) {
		return "Permutasi dong"
	}
	return "Bukan dong"
}

//untuk function pembalikan
func reverse(a string) string {
	d := []int32(a)
	for b, l := 0, len(d)-1; b < len(d)/2; b, l = b+1, l-1 {
		d[b], d[l] = d[l], d[b]
	}
	return string(d)
}
