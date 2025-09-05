package main //Todo programa executável em Go deve ter o pacote main

import (
	"encoding/json"
	"fmt"
	"os"
)

//Criamos uma struct chamada Person. As tags json:"..." são importantes para mapear os campos da struct para os campos do JSON.
type Person struct{
	ID int `json:"id"`
	Nome string `json:"nome"`
	Idade int `json:"idade"`
	Cidade string `json:"cidade"`
}

var people []Person //Declaramos uma variável global, uma slice de Person, para armazenar os dados.

//Esta função é executada automaticamente antes da main. É o local perfeito para carregar os nossos dados iniciais.
func init(){
	file, err := os.ReadFile("data.json") // Lê o conteúdo do arquivo.
	if err != nil{
		fmt.Println("Erro ao ler o arquivo JSON:", err)
		return
	}
	err = json.Unmarshal(file, &people) //Converte o conteúdo do JSON em uma slice de Person.
	if err != nil{
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}
}

func ReadAll(){
	fmt.Println("\n---Listando Pessoas---")
	for _, p := range people{
		fmt.Printf("ID: %d, Nome: %s, Idade: %d, Cidade: %s\n", p.ID, p.Nome, p.Idade, p.Cidade)
	}
	fmt.Println("-----------------")
}

func Create(p Person){
	p.ID = people[len(people)-1].ID + 1
	people = append(people, p)
	fmt.Printf("\nPessoa '%s' adicionada com sucesso", p.Nome)
}

func Update(id int, updatedPerson Person){
	for i, p := range people{
		if p.ID == id{
			people[i] = updatedPerson
			people[i].ID = id
			fmt.Printf("\nPessoa com ID %d atualizada com sucesso!\n", p.ID)
			return
		}
	}
	fmt.Printf("\nPerssoa com ID %d nao encontrada.\n", id)
}

func Delete(id int){
	for i, p := range people{
		if p.ID == id{
			people = append(people[:i], people[i+1:]...)
			fmt.Printf("\nPessoa com ID %d removida com sucesso", id)
			return
		}
	}
	fmt.Printf("\nPessoa com ID %d nao encontrada.\n", id)
}


func main(){
	fmt.Println("Dados carregados com sucesso!")


	ReadAll()

	newPerson := Person{Nome: "Jojo Toddynho", Idade: 100, Cidade:"Hell de Janeiro"}
	Create(newPerson)
	ReadAll()

	updatePerson := Person{Nome: "Alice Silva", Idade: 69, Cidade: "Sao Baulo"}
	Update(1, updatePerson)
	ReadAll()

	Delete(2)
	ReadAll()
}