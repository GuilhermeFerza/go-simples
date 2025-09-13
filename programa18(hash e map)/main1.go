package main

import "fmt"

func main1(){

	palavras := []string{"go", "java", "go", "python", "go", "java", "rust"}

	//criamos um mapa para contar
	contador := make(map[string]int)

	for _, palavra := range palavras {
		contador[palavra] += 1
	}

	fmt.Println("contagem de palavras:")
	for palavra, qtd := range contador{
		fmt.Printf("%s: %d\n", palavra, qtd)
	}
}