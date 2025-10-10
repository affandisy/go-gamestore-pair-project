package repository_test

import (
	"database/sql"
	"fmt"
	"gamestore/internal/domain"
	"gamestore/internal/repository"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var testRepo *repository.CustomerRepository

func TestMain(m *testing.M) {
	_ = godotenv.Load("../../.env")

	connStr := os.Getenv("TEST_DB_URL")
	if connStr == "" {
		log.Fatal("TEST_DB_URL not found — make sure .env is loaded")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Ping failed: %v", err)
	}

	testRepo = repository.NewCustomerRepository(db)
	code := m.Run()
	os.Exit(code)
}

func TestCreateAndFindCustomer(t *testing.T) {
	cust := domain.Customer{
		Name:      "Test User",
		Email:     fmt.Sprintf("test_%d@mail.com", time.Now().UnixNano()),
		Password:  "password123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := testRepo.Create(&cust)
	assert.NoError(t, err, "Expected no error when inserting customer")

	customers, err := testRepo.FindAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, customers, "Expected at least one customer in DB")

	found := false
	for _, c := range customers {
		if c.Email == cust.Email {
			found = true
			break
		}
	}
	assert.True(t, found, "Expected created customer to exist in database")
}
