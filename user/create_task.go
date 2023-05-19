package user

import (
	"github.com/sirupsen/logrus"
)

type TaskParams struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Completed   bool   `json:"completed"`
}

func (u userAppService) Create_task() {

	logrus.Info("hola estoy en create task del user app service")
	//fmt.Println("hola estoy en create task del user app service")
}
