package main

import "fmt"

func fila() {
	fila := []int{1, 2, 3, 4, 5}

	for _, numero := range fila {
		fmt.Printf("atendendo %d\n", numero)
	}

}