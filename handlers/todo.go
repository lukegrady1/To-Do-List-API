package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/lukegrady1/to-do-list-api/models"
)

var DB *gorm.DB

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

func GetTodoByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo models.Todo
	if err := DB.First(&todo, id).Error; err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo models.Todo
	if err := DB.First(&todo, id).Error; err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DB.Delete(&models.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
