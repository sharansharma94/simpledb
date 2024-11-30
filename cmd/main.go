package main

import (
	"fmt"
	"log"

	"github.com/sharansharma94/simpledb/internal/db"
)

func main() {
	database, err := db.NewDatabase("data/db.log")
	if err != nil {
		log.Fatalf("error creating database: %v", err)
	}

	err = database.Write("username", "sharan", false)
	if err != nil {
		log.Fatalf("error writing to database: %v", err)
	}

	value, err := database.Read("username")
	if err != nil {
		log.Fatalf("error reading from database: %v", err)
	}
	log.Printf("value: %s", value)

	err = database.Delete("username")
	if err != nil {
		log.Fatalf("error deleting from database: %v", err)
	}

	fmt.Println("key deleted successfully")
}
