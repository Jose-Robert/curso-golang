package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1515
	var raio = 3.2

	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)
}
