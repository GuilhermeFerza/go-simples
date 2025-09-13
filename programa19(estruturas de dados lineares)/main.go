package main

import "fmt"

func main() {
	lista := []string{"teste", "teste1", "teste2"}

	novoValor := "algo legal"
	lista = lista[:len(lista)-1]
	lista = append(lista, novoValor)

	fmt.Println(lista)
}