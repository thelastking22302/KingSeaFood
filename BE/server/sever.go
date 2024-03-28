package server

import (
	"fmt"
	"log"
	"thelastking/kingseafood/pkg/db"
)

func Run() {
	config := &db.Config{
		Host:     "localhost",
		Port:     5432,
		Password: "220302",
		User:     "postgres",
		DbName:   "seafood",
	}
	db, err := config.NewConnection()
	if err != nil {
		log.Fatalf("Fails to connect to database: %v", err)
	}
	fmt.Printf("Connect suscess to database: %v", db)
}
