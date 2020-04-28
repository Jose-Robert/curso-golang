package main

import "fmt"

func main() {
	funcsESalario := map[string]float64{
		"José Robert":   11325.45,
		"Ruth França":   15456.78,
		"Bryan Calleri": 1200.00,
	}

	funcsESalario["Pedro Junior"] = 1350.51
	funcsESalario["João Pessoa"] = 937.00
	delete(funcsESalario, "Pedro Junior")

	for nome, salario := range funcsESalario {
		fmt.Println(nome, salario)
	}
}
