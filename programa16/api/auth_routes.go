package api

import (
	"encoding/json"
	"fmt"
	"os"
	"programa16/internal/domain"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var people []domain.Person

func init() {
	file, err := os.ReadFile("data/data.json")
	if err != nil {
		fmt.Println("Erro ao ler arquivo json")
		return
	}
	err = json.Unmarshal(file, &people)
	if err != nil {
		fmt.Println("erro ao decodificar arquivo json")
		return
	}
}

func Api(){
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/api/data/", func(c *gin.Context) {
		c.JSON(200, people)
	})
	r.Run(":8080")
}