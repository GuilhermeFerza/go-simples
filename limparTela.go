package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func limparTela(){
	var cmd *exec.Cmd

	switch runtime.GOOS{
	case "windows":
		cmd=exec.Command("cmd", "/c", "cls")
	default:
		cmd=exec.Command("clear")
	}

	cmd.Stdout = os.Stdout 

	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao limpar a tela: ", err)
	}
}