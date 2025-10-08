package connections

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewConnection() *sql.DB {
	dsn := "postgres://postgres:@localhost:5432/gamestore?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln("Error: ", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		log.Fatalln("Error: ", err)
		return nil
	}

	fmt.Println("Connected to the database")
	return db
}
