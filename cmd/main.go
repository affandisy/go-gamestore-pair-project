package main

import (
	"gamestore/internal/cli"
	"gamestore/internal/connections"
	"gamestore/internal/repository"
	"gamestore/internal/usecase"
)

func main() {
	db := connections.NewConnection()
	defer db.Close()

	customerRepo := &repository.CustomerRepository{DB: db}
	customerUc := usecase.NewCustomerUsecase(customerRepo)

	app := cli.AppMenu{
		CustomerUC: customerUc,
	}

	app.Run()

}

// KATEGORI PENILAIN
// debuggin
// problem solving
// testify, database deployment
