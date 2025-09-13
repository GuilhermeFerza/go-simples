package main

import "fmt"

func pilha() {
	var pilha []string

	pilha = append(pilha, "guilherme")
	pilha = append(pilha, "danilo")
	pilha = append(pilha, "fernando")

	fmt.Println(pilha)

	ultimo := pilha[len(pilha)-1]
	pilha = pilha[:len(pilha)-1]
	fmt.Println("Removido:", ultimo)
	fmt.Println("Pilha atual:", pilha)
}