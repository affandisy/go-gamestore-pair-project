package cli

import (
	"fmt"
	"gamestore/internal/usecase"

	"github.com/manifoldco/promptui"
)

func orderGames(customerID int64, ucGame *usecase.GameUsecase, ucOrder *usecase.Orderusecase, ucPay *usecase.Paymentusecase, ucLib *usecase.Libraryusecase) {
	for {
		orders, err := ucOrder.FindAllOrderByCustomerID(customerID)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		var orderName = []string{}
		for _, order := range orders {
			orderName = append(orderName, fmt.Sprintf("%s - %.2f", order.GameTitle, order.GamePrice))

		}
		if len(orderName) > 0 {
			orderName = append(orderName, "Bayar semua")
		}
		orderName = append(orderName, "Back")
		orderMenu := promptui.Select{
			Label: "Orders",
			Items: orderName,
		}
		idx, selectedOrderMenu, _ := orderMenu.Run()
		if selectedOrderMenu == "Back" {
			break
		}
		if selectedOrderMenu == "Bayar semua" {
			err := payAllGames(customerID, ucPay, ucOrder, ucLib, orders)
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}
			continue
		}
		selectedGame := orders[idx]

		for {
			game, err := ucGame.FindGameById(selectedGame.GameID)
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}
			orderID := selectedGame.OrderID
			price := selectedGame.GamePrice

			menuGameOrder := promptui.Select{
				Label: "Actions",
				Items: []string{"Buy", "Remove", "Back"},
			}

			_, selectedMenuGameOrder, _ := menuGameOrder.Run()

			var isBought = false
			var isRemoved = false
			switch selectedMenuGameOrder {
			case "Buy":
				fmt.Println("Buy", orderID, price)
				err := payOneGame(customerID, ucOrder, ucPay, ucLib, game)
				if err != nil {
					fmt.Println("Error: ", err)
					continue
				}
				err2 := ucOrder.UpdateOrderStatus(orderID, "PAID")
				if err2 != nil {
					fmt.Println("Error: ", err2)
					continue
				}

				isBought = true
			case "Remove":
				err := ucOrder.DeleteOrder(orderID)
				if err != nil {
					fmt.Println("Error: ", err)
					continue
				}
				fmt.Printf("%s is removed\n", selectedGame.GameTitle)
				isRemoved = true
			case "Back":
				// break
			}
			if isBought || isRemoved || selectedMenuGameOrder == "Back" {
				break
			}
		}

	}

}
