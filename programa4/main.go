package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("digite uma frase qualquer: ")
	frase, _ := reader.ReadString('\n')
	frase = strings.TrimSpace(frase)
	
	palavras := strings.Fields(frase)

	contagem := make(map[string]int)

	for _, palavra := range palavras{
		palavra = strings.ToLower(palavra)
		contagem[palavra]++
	}
	
	fmt.Println("frequencias de palavras:")
	for palavra, qtd := range contagem{
		fmt.Println(palavra, qtd)
	}
	fmt.Printf("\nTotal de palavras: %d\n", len(palavras))
}