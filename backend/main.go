package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Habit struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var habits = []Habit{
	{ID: 1, Name: "Drink Water"},
	{ID: 2, Name: "Exercise"},
}

func main() {
	r := gin.Default()

	// Allow CORS (for Vue)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	// Routes
	r.GET("/habits", func(c *gin.Context) {
		c.JSON(http.StatusOK, habits)
	})

	r.POST("/habits", func(c *gin.Context) {
		var newHabit Habit
		if err := c.BindJSON(&newHabit); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newHabit.ID = len(habits) + 1
		habits = append(habits, newHabit)
		c.JSON(http.StatusCreated, newHabit)
	})

	r.Run(":8080") // listen on port 8080
}
