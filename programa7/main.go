package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User

func loadUsers() {
	file, err := os.ReadFile("user.json")
	if err != nil{
		fmt.Println("erro ao ler arquivo JSON.", err)
		return
	}
	err = json.Unmarshal(file, &users)
	if err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}

func main() {

	loadUsers()

	http.HandleFunc("/users", getUsers)
	fmt.Println("servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

