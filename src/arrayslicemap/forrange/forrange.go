package main

import (
	"fmt"
)

func main() {
	numeros := [...]int{2, 4, 6, 8, 10} //compilador conta

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}
}
