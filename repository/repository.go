package repository

import (
	"database/sql"
	_ "embed"
)

//go:embed sql/create_task.sql
var create_taskQuery string

var print_taskQuery string

type repository struct {
	db          *sql.DB
	create_task *sql.Stmt
	//print_task  *sql.Stmt
}

func NewUserRepo(db *sql.DB) (*repository, error) {

	create_taskStmt, err := db.Prepare(create_taskQuery)
	if err != nil {
		return nil, err
	}

	/*print_taskStmt, err := db.Prepare(print_taskQuery)
	if err != nil {
		return nil, err
	}*/

	return &repository{
		db:          db,
		create_task: create_taskStmt,
		//print_task:  print_taskStmt,
	}, nil
}
