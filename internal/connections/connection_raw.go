package connections

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// type DB struct {
// 	Conn *sql.DB
// }

func NewConnection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil
	}

	connectionString := os.Getenv("CONNECTION_STRING")

	dsn := connectionString

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
