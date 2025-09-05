package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Person struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

var people []Person

func init() {
	file, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo JSON:", err)
		return
	}

	err = json.Unmarshal(file, &people)
	if err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/api/data", func(c *gin.Context) {
		// Substitua a mensagem estática pela variável 'people'
		c.JSON(200, people) 
	})
	r.POST("/api/data/post", func(c *gin.Context){
		c.JSON(200, people)
	})

	fmt.Println("Dados carregados com sucesso!")
	r.Run(":8080")
}