package models

type Task struct {
	Id int `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`		
}


