package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Person struct{
	ID int `json:"id"`
	Nome string `json:"nome"`
	Email string `json:"email"`
}

var people []Person

func init(){
	file, err := os.ReadFile("data.json")
	if err != nil{
		fmt.Println("erro ao ler arquivo JSON")
		return
	}
	err = json.Unmarshal(file, &people)
	if err != nil{
		fmt.Println("erro ao decodificar arquivo JSON")
		return
	}
}

func main(){
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/api/data/", func(c *gin.Context){
		c.JSON(200, people)
	})
	r.POST("/api/data/", func(c *gin.Context){
		var newPerson Person
		if err := c.ShouldBindJSON(&newPerson); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newPerson.ID=len(people) + 1
		people = append(people, newPerson)
		c.JSON(http.StatusCreated, newPerson)
	})
	r.DELETE("/api/data/:id", func(c *gin.Context){
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		var index int = -1

		for i, p := range people{
			if p.ID == id{
				index = i
				break
			}
		}
		if index == -1{
			c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
			return
		}
		people = append(people[:index], people[index+1:]...)
		 c.JSON(http.StatusOK, gin.H{"message": "Person deleted"})
	})
	r.PUT("/api/data/:id", func(c *gin.Context){
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id format"})
			return
		}

		var updatedPerson Person
		if err := c.ShouldBindJSON(&updatedPerson); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var found bool = false
		
		for i, p := range people{
			if p.ID == id{
				people[i].Nome = updatedPerson.Nome
				people[i].Email= updatedPerson.Email
				found = true
				break
			}
		}
		if !found{
			c.JSON(http.StatusNotFound, gin.H{"error": "person not found"})
			return
		}
			c.JSON(http.StatusOK, updatedPerson)
	})
	fmt.Println("dados carregados com sucesso!")
	r.Run(":8080")
}