package main

import "fmt"

func main() {

	listaCompras := []string{}
	var op int
	for op != 4{
		menu()
		fmt.Scan(&op)
		switch op{
			case 1: 
				if len(listaCompras) == 0{
					fmt.Println("nao tem nenhuma compra")
				}else{
					fmt.Println(listaCompras)
				}
			case 2: 
				listaCompras = adicionar(listaCompras)
			case 3: 
				listaCompras = retirar(listaCompras)
		}
	}
}

func menu(){
	println("selecione-----------------------")
	println("1. mostrar lista----------------")
	println("2. adicionar a lista------------")
	println("3. remover da lista-------------")
	println("4. Sair-------------------------")
}

func adicionar(listaCompras []string)[]string{
	var compra string
	op := "s"
	for op == "s"{
		fmt.Println("adicione a compra:")
		fmt.Scan(&compra)
		listaCompras = append(listaCompras, compra)	
		fmt.Println("deseja adicionar mais? s/n")
		fmt.Scan(&op)
	}
	return listaCompras
}
func retirar(listaCompras []string)[]string{
	removido := listaCompras[len(listaCompras)-1]
	listaCompras = listaCompras[:len(listaCompras)-1]
	fmt.Println("Removido:", removido)
	fmt.Println("atual:", listaCompras)
	return listaCompras
}
