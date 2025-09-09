package main

import "fmt"

// Função para gerar todas as permutações de um slice de inteiros
func permute(arr []int, l int, r int) {
	if l == r {
		fmt.Println(arr) // imprime a permutação
		return
	}
	for i := l; i <= r; i++ {
		arr[l], arr[i] = arr[i], arr[l]       // troca elementos
		permute(arr, l+1, r)                  // chamada recursiva
		arr[l], arr[i] = arr[i], arr[l]       // desfaz a troca (backtrack)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	permute(arr, 0, len(arr)-1)
}
