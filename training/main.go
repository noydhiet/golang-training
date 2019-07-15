package main
import "fmt"

/*
Main class for hello world
*/
func main(){

	var firtsName string = "Adly"
	var lastName string = "Maulana"
	/* new digunakan untuk mencetak data 
	pointer dengan tipe data tertentu */
	address := new((string)) 

	/*komentar kode 
	menampilkan pesan hello-world
	*/
	fmt.Println("hallo", firtsName, lastName, "!!")
	fmt.Printf("halo %s %s!\n", firtsName, lastName)
	
	fmt.Printf(*address) /*variable addrress menampung 
	data bertipe pointer string. Jika ditampilkan yang
	muncul adalah alamat memori nilai tsbt*/
	
	// fmt.println("baris ini tidak akan di eksekusi")
}