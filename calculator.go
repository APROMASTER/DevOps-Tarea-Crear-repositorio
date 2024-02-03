package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var result int

	fmt.Println("Suma de numeros - Hecho en GoLang por Ariel Nuñez 2022-0551")
	fmt.Print("\nInserte un numero: ")
	fmt.Scan(&a)
	fmt.Print("\nInserte otro numero: ")
	fmt.Scan(&b)

	result = a + b
	fmt.Println("\nEntonces tenemos que:\n", a, "+", b, "=", result)
}
