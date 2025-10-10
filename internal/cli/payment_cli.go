package cli

import (
	"fmt"
	"gamestore/internal/domain"
	"gamestore/internal/usecase"
)

func payOneGame(customerID int64, ucOrder *usecase.Orderusecase, ucPay *usecase.Paymentusecase, game *domain.Game) error {
	order, err := ucOrder.CreateOrder(customerID, game.GameID)
	if err != nil {
		return err
	}
	ucPay.CreatePayment(customerID, float64(game.Price), "paid")
	ucOrder.DeleteOrder(order.OrderID)
	fmt.Println("Berhasil membayar game")
	return nil
}

func payAllGames(customerID int64, ucPay *usecase.Paymentusecase, ucOrder *usecase.Orderusecase, orders []domain.Order) error {
	err := ucPay.PayAllUserGames(customerID, "PAID")
	if err != nil {
		return err
	}

	fmt.Println("Berhasil membayar semua game!")
	for _, order := range orders {
		err := ucOrder.UpdateOrderStatus(order.OrderID, "PAID")
		if err != nil {
			return err
		}
	}

	return nil
}
