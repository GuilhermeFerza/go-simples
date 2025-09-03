package main

import (
	"fmt"
	"time"
)

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
				if(ultimoValor == 0){
					fmt.Println("ainda nao ha nenhum valor.")
				}else{
					fmt.Print("\n\nultimo valor registrado: ", ultimoValor,"\n\n")
				}
				time.Sleep(2 * time.Second)
				limparTela()
				
			case 3:
				fmt.Println("saindo...")
				time.Sleep(2 * time.Second)
				limparTela()
				return
			default:
				fmt.Println("resposta invalida")
		}
	}
}