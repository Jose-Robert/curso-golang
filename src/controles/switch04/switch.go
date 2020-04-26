package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {

	switch i.(type) {
	case int:
		return "Inteiro"
	case float32, float64:
		return "Real"
	case string:
		return "String"
	case func():
		return "Função"
	default:
		return "Tipo não reconhecido."
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(10))
	fmt.Println(tipo("Isso é um texto"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
