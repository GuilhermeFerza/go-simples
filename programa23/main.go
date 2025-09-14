package main

import "fmt"

func main() {

	slice := []int{
		10, 5, 9, 1, 2, 8, 4, 7, 3, 6,
}

	lista := quickSort(slice)
	fmt.Println(lista)
}

func bubbleSort(slice []int) []int {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
		fmt.Printf("Passo %d: %v\n", i+1, slice)
	}
	return slice
}

func quickSort(slice []int)[]int{
	if len(slice) <= 1{
		return slice
	}
	pivo := slice[len(slice)/2]
	var menores, maiores, iguais []int

	for _ ,v := range slice{
		if v < pivo{
			menores = append(menores, v)
		}else if v > pivo{
			maiores = append(maiores, v)
		}else{
			iguais = append(iguais, v)
		}
	}
	
	fmt.Printf("Pivô: %d | Menores: %v | Iguais: %v | Maiores: %v\n", pivo, menores, iguais, maiores)

	menores = quickSort(menores)
	maiores = quickSort(maiores)

	return append(append(menores,iguais...),maiores...)
}