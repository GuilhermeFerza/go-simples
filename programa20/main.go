package main

import "fmt"

func main() {

	slice := make([]int, 1500000)
	for i := 0; i < 1500000; i++ {
		slice[i] = i + 1
	}

	alvo := 18924

	resultado := bs(slice, alvo)

	if resultado != -1 {
		fmt.Printf("O número %d foi encontrado no índice %d\n", alvo, resultado)
	} else {
		fmt.Printf("O número %d não foi encontrado\n", alvo)
	}
}