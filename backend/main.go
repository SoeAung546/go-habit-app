package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Habit struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var habits = []Habit{
	{ID: 1, Name: "Drink Water", Done: false},
	{ID: 2, Name: "Exercise", Done: false},
	{ID: 3, Name: "Read a Book", Done: false},
	{ID: 4, Name: "Meditate", Done: false},
}

func main() {
	r := gin.Default()

	// Allow CORS (for Vue)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
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
		newHabit.Done = false
		habits = append(habits, newHabit)
		c.JSON(http.StatusCreated, newHabit)
	})

	r.DELETE("/habits/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		for i, h := range habits {
			if h.ID == id {
				habits = append(habits[:i], habits[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Habit deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found"})
	})

	r.PATCH("/habits/:id/done", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID",
			})
		}

		for i, h := range habits {
			if h.ID == id {
				habits[i].Done = !habits[i].Done
				c.JSON(http.StatusOK, habits[i])
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Habit ot found"})
	})

	r.Run(":8080") // listen on port 8080
}
