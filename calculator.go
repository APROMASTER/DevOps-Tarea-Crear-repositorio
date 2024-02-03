package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var result int

	fmt.Println("Suma de numeros - Hecho en GoLang!")
	fmt.Print("\nInserte un numero: ")
	fmt.Scan(&a)
	fmt.Print("\nInserte otro numero: ")
	fmt.Scan(&b)

	result = a + b
	fmt.Println("Entonces tenemos que\n", a, "+", b, "=", result)
}
