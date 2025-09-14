package main

import (
	"fmt"
)


//nao completamente entendido


// Hash simples: soma dos valores ASCII dos caracteres, módulo do tamanho
func hash(chave string, tamanho int) int {
	soma := 0
	for _, char := range chave {
		soma += int(char)
	}
	return soma % tamanho
}

func main() {
	// Criamos uma "tabela hash" com 10 gavetas
	tamanho := 10
	tabela := make([][][2]string, tamanho) // cada gaveta guarda pares [chave, valor]

	// Função para inserir no "map"
	inserir := func(chave, valor string) {
		indice := hash(chave, tamanho)
		tabela[indice] = append(tabela[indice], [2]string{chave, valor})
	}

	// Função para buscar no "map"
	buscar := func(chave string) string {
		indice := hash(chave, tamanho)
		for _, par := range tabela[indice] {
			if par[0] == chave {
				return par[1]
			}
		}
		return "não encontrado"
	}

	// Inserindo valores
	inserir("João", "25")
	inserir("Maria", "30")
	inserir("Ana", "20")
	inserir("Pedro", "40")

	// Testando busca
	fmt.Println("Idade João:", buscar("João"))
	fmt.Println("Idade Maria:", buscar("Maria"))
	fmt.Println("Idade Ana:", buscar("Ana"))
	fmt.Println("Idade Pedro:", buscar("Pedro"))

	// Mostrando a tabela
	fmt.Println("\nTabela Hash interna:")
	for i, bucket := range tabela {
		fmt.Printf("Gaveta %d: %v\n", i, bucket)
	}
}
