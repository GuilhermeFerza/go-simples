package main

import "fmt"

func main(){
	var primeiroNum, segundoNum, terceiroNum float64

	fmt.Print("Digite o primeiro numero: ")
	fmt.Scan(&primeiroNum)
	fmt.Print("Digite o segundo numero: ")
	fmt.Scan(&segundoNum)
	fmt.Print("Digite o terceiro numero: ")
	fmt.Scan(&terceiroNum)

	sum := primeiroNum+segundoNum+terceiroNum
	med := sum/3

	fmt.Println("soma do total:", sum)
	fmt.Printf("media dos numeros: %.2f\n", med)

	if primeiroNum == segundoNum && primeiroNum == terceiroNum {
    fmt.Println("todos os numeros sao iguais:", primeiroNum)
	} else if primeiroNum >= segundoNum && primeiroNum >= terceiroNum {
    fmt.Println("o maior numero eh o:", primeiroNum)
	} else if segundoNum >= primeiroNum && segundoNum >= terceiroNum {
    fmt.Println("o maior numero eh o:", segundoNum)
	} else {
    fmt.Println("o maior numero eh o:", terceiroNum)
	}


}