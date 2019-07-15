package main
import "fmt"

func main() {
	var usia = 22

	switch {
	case (usia < 25) && (usia > 20):
		fmt.Println("Matang")
	default:
		fmt.Println("Unable") 
	}
}