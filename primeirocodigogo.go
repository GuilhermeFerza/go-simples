package main

import (
	"fmt"
	"strings"
)
func main(){
	var produto float64
	var valor float64
	isRunning := true
	var totalRes float64
	for isRunning{
		fmt.Println("digite qtd. de produtos:")
		fmt.Scan(&produto)
		fmt.Println("digite o valor:")
		fmt.Scan(&valor)
		var res = produto*valor
		fmt.Println(res)
		var resposta string
		fmt.Println("cont?")
		fmt.Scan(&resposta)
		respostaLower := strings.ToLower(resposta)
		switch respostaLower {
			case "sim", "ss", "s":
				isRunning = true
				totalRes += res
				fmt.Println(totalRes)
			case "nao", "nn", "n":
				isRunning = false
				fmt.Println("encerrado")
			default:
				fmt.Println("valor invalido")
		}
		
	}	
}