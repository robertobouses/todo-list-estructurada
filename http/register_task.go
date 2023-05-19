package http

import "github.com/gin-gonic/gin"

//github.com/robertobouses/

type RegisterTaskRequest struct {
	ID          int    `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `description:"lastName" binding:"required"`
	DueDate     string `duedate:"password" binding:"required"`
	Completed   bool   `completed:"password" binding:"required"`
}

func (h Handler) RegisterTask(c *gin.Context) {

}
