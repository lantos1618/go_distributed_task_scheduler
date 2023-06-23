package main

import (
	"database/sql"
	"log"
)

type Task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Schedule    string `json:"schedule"`
}

func main() {
	connStr := "user=myuser password=mypassword dbname=mydatabase sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
