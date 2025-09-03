package main

import (
	"fmt"
	"strings"
)

func calculadora() {
	var produto float64
	var valor float64
	isRunning := true
	var totalRes float64
	i := 1
	for isRunning {
		fmt.Println("digite qtd. do", i, "produtos:")
		fmt.Scan(&produto)
		fmt.Println("digite o valor:")
		fmt.Scan(&valor)
		var res = produto * valor
		fmt.Print("\n", res, "\n\n")
		var resposta string
		fmt.Print("continuar compra? ")
		fmt.Scan(&resposta)
		fmt.Println("")
		respostaLower := strings.ToLower(resposta)
		switch respostaLower {
		case "sim", "ss", "s":
			isRunning = true
			totalRes += res
			fmt.Print("\nvalor total do carrinho: ", totalRes, "\n\n")
			fmt.Print("\n----------------------------------\n\n")
			i += 1
		case "nao", "nn", "n":
			isRunning = false
			totalRes += res
			fmt.Print("valor total do carrinho: ", totalRes, "\n\n")
			fmt.Println("encerrado...")
		default:
			isRunning = false
			fmt.Println("valor invalido!")
		}
	}
}