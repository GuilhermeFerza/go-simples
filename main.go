package main

import "fmt"

func main() {
	var op int
	var ultimoValor float64


	for{
		menu()
	fmt.Scan(&op)
	
	switch op {
			case 1:
				ultimoValor = calculadora()
			case 2:
				fmt.Println(ultimoValor)
			case 3:
				fmt.Println("saindo...")
				return
			default:
				fmt.Println("resposta invalida")
		}
	}
	
}