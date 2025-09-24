package main

import "fmt"

func main() {
	slice := make([]int, 15000)
	for i := 0; i < 15000; i++ {
		slice[i] = i + 1
	}
	alvo := 145

	resultado := bs(slice, alvo)

	if resultado != -1{
		fmt.Printf("resultado = %d, indice = %d", alvo, resultado)
	}else{
		fmt.Println("resultado nao encontrado")
	}

}

func bs(slice []int, alvo int) int {
	inicio := 0
	fim := len(slice) - 1
	for inicio <= fim {
		meio := (inicio + fim) / 2
		if slice[meio] == alvo {
			return meio
		} else if slice[meio] < alvo {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}
	return -1
}