package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"server/internal/domain"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var task []domain.Tasks

func init() {
	file, err := os.ReadFile("data/data.json")
	if err != nil{
		fmt.Println("Erro ao ler arquivo JSON")
		return
	}
	err = json.Unmarshal(file, &task)
	if err != nil{
		fmt.Println("erro ao decodificar arquivo json")
		return
	}
}

func Api(){
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin","Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Lenght"},
		AllowCredentials: true,
	}))

	r.GET("/api/data/", func(c *gin.Context) {
		c.JSON(200, task)
	})

	r.POST("/api/data", func(c *gin.Context) {
		var newTask domain.Tasks
		if err := c.ShouldBindJSON(&newTask); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newTask.ID = len(task) + 1
		task = append(task, newTask)
		c.JSON(http.StatusCreated, newTask)
	})
	r.DELETE("/api/data/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
			return
		}

		var index int = -1
		for i, p := range task{
			if p.ID == id{
				index = i
				break
			}
		}
		if index == -1{
			c.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
			return
		}
		task = append(task[:index],task[index+1:]...)
		c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
	})
	r.PUT("/api/data/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
			return
		}

		var updatedTask domain.Tasks
		if err := c.ShouldBindJSON(&updatedTask); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var found bool = false
		for i, p := range task{
			if p.ID == id{
				task[i].Task = updatedTask.Task
				found = true
				break
			}
		}
		if !found{
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}
		c.JSON(http.StatusOK, updatedTask)
	})
	r.Run(":8080")
}