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
