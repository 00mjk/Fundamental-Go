package main

import "fmt"

func main() {
	fmt.Printf("What is the candy title? ")
	var candyTitle string
	fmt.Scanln(&candyTitle)
	fmt.Printf("What is the candy flavor? ")
	var candyFlavor string
	fmt.Scanln(&candyFlavor)

	fmt.Printf("It shall be called %s %s!", candyTitle, candyFlavor)
}
