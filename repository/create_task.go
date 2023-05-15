package repository

import "github.com/gin-gonic/gin"

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Completed   bool   `json:"completed"`
}

func (r *repository) Create_task(c *gin.Context) {
	return
}
