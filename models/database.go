package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const Driver string = "mysql"
const Username string = "golang"
const Password string = "gopherito"
const Host string = "localhost"
const Port int = 3306
const Database string = "tasks"

func CreateConection() (db *sql.DB) {
	conection, err := sql.Open(Driver, generateURL())
	db = conection
	if err != nil {
		panic(err)
	}
	fmt.Println("conectada a la base de datos...")
	return db

}

var db = CreateConection()

func Ping() {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
}

func CloseConection() {
	defer db.Close()
}

func generateURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", Username, Password, Host, Port, Database)
}
