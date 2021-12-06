package main

import (
	"fmt"
	"log"

	"github.com/nabhajit-ray/projects/http_server_golang/internal/database"
)

func main() {
	c := database.NewClient("db.json")
	err := c.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database ensured!")
}
