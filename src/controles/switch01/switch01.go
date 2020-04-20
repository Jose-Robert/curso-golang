package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)

	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota Inválida."
	}
}

func main() {
	fmt.Println("A nota é de categoria", notaParaConceito(9.1))
	fmt.Println("A nota é de categoria", notaParaConceito(7.2))
	fmt.Println("A nota é de categoria", notaParaConceito(1.55))
	fmt.Println("A nota é de categoria", notaParaConceito(11.90))
}
