package main

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