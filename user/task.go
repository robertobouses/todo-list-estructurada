package user

import (
	"errors"
)

var (
	IdIsRequired          = errors.New("id is required")
	TitleIsRequired       = errors.New("title is required")
	DescriptionIsRequired = errors.New("description is required")
	DueDateIsRequired     = errors.New("due date is required")
	CompletedIsTrue       = errors.New("completed is already true")
)

type Task struct {
	ID          int
	Title       string
	Description string
	DueDate     string
	Completed   bool
}
