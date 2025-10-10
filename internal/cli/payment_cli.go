package cli

import (
	"fmt"
	"gamestore/internal/domain"
	"gamestore/internal/usecase"
)

func payOneGame(customerID int64, ucOrder *usecase.Orderusecase, ucPay *usecase.Paymentusecase, ucLib *usecase.Libraryusecase, game *domain.Game) error {
	order, err := ucOrder.CreateOrder(customerID, game.GameID)
	if err != nil {
		return err
	}
	ucPay.CreatePayment(customerID, float64(game.Price), "PAID")
	err2 := ucLib.CreateGameInLibrary(customerID, game.GameID)
	if err2 != nil {

		return err
	}

	err3 := ucOrder.UpdateOrderStatus(order.OrderID, "PAID")
	if err3 != nil {
		return err3
	}

	fmt.Println("Berhasil membayar game")
	return nil
}

func payAllGames(customerID int64, ucPay *usecase.Paymentusecase, ucOrder *usecase.Orderusecase, ucLib *usecase.Libraryusecase, orders []domain.Order) error {
	// Bayar semua game sekaligus
	err := ucPay.PayAllUserGames(customerID, "PAID")
	if err != nil {
		return err
	}

	// Update semua order status
	for _, order := range orders {
		if err := ucOrder.UpdateOrderStatus(order.OrderID, "PAID"); err != nil {
			return err
		}
	}

	// Tambahkan semua game ke library
	for _, order := range orders {
		if err := ucLib.CreateGameInLibrary(customerID, order.GameID); err != nil {
			return err
		}
	}

	fmt.Println("Berhasil membayar semua game!")
	return nil
}
