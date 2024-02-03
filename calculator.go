package main

import (
	"fmt"
)

func main() {
	var operation int

	fmt.Println("Calculadora de 2 numeros - Hecho en GoLang por Ariel Nuñez 2022-0551")
	fmt.Println("\nTipos de operaciones:")
	fmt.Println("	1- Suma")
	fmt.Println("	2- Resta")
	fmt.Println("	3- Multiplicacion")
	fmt.Println("	4- Division")

	for operation < 1 || operation > 4 {
		fmt.Print("\nInserte el tipo de operacion para hacer: ")
		fmt.Scan(&operation)
	}

	calculate(operation)

}

func calculate(operation int) {
	var a int
	var b int
	var result int
	var symbol string

	fmt.Print("\nInserte un numero: ")
	fmt.Scan(&a)
	fmt.Print("\nInserte otro numero: ")
	fmt.Scan(&b)

	switch operation {
	case 1:
		symbol = "+"
		result = sum(a, b)

	case 2:
		symbol = "-"
		result = subtract(a, b)

	case 3:
		symbol = "*"
		result = multiply(a, b)

	case 4:
		symbol = "/"
		result = divide(a, b)

	}

	fmt.Println("\nEntonces tenemos que:", a, symbol, b, "=", result, "✔")
}

func sum(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	return a / b
}
