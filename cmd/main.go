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
	gameRepo := &repository.GameRepository{DB: db}
	categoryRepo := &repository.CategoryRepository{DB: db}
	orderRepo := &repository.OrderRepository{DB: db}
	paymentRepo := &repository.PaymentRepository{DB: db}

	customerUc := usecase.NewCustomerUsecase(customerRepo)
	gameUc := usecase.NewGameUsecase(gameRepo)
	categoryUc := usecase.NewCategoryUsecase(categoryRepo)
	orderUc := usecase.NewOrderUsecase(orderRepo)
	paymentUc := usecase.NewPaymentUsecase(paymentRepo)

	app := cli.AppMenu{
		CustomerUC: customerUc,
		GameUC:     gameUc,
		CategoryUC: categoryUc,
		OrderUC:    orderUc,
		PaymentUC:  paymentUc,
	}

	app.Run()
}

// KATEGORI PENILAIN
// debuggin
// problem solving
// testify, database deployment
