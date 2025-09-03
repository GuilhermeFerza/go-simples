package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	var idade int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("me informe seu nome")
	nome, _ := reader.ReadString('\n')
	nome = strings.TrimSpace(nome)
	
	fmt.Println("me informe sua idade")
	idadeStr, _ := reader.ReadString('\n')
	idadeStr = strings.TrimSpace(idadeStr)
	idade, _ = strconv.Atoi(idadeStr)
	
	dezAnos := idade+10
	fmt.Printf("%s daqui a 10 anos voce tera %d\n", nome, dezAnos)

	if idade >= 18{
		fmt.Println("voce eh maior de idade")
	}else{
		fmt.Println("voce eh menor de idade")
	}
	
}