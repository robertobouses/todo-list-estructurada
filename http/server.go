package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type APP interface {
	Create_task()
	//Print_task(task.PrintTaskParams) (*task.Task, error)
}

type Server struct {
	handler Handler
	engine  *gin.Engine
}

type Handler struct {
	app APP
}

func NewServer(app APP) Server {
	h := Handler{
		app: app,
	}
	return Server{
		handler: h,
		engine:  gin.Default(),
	}
}

func (s *Server) Run(port string) error {
	usr := s.engine.Group("/task")
	usr.POST("/", s.handler.RegisterTask)
	//usr.GET("/print", s.handler.app.Print_task())
	return s.engine.Run(fmt.Sprintf(":%s", port))

}
