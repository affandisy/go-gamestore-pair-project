package cli

import (
	"fmt"
	"gamestore/internal/usecase"

	"github.com/manifoldco/promptui"
)

func orderGames(customerID int64, uc *usecase.Orderusecase) {
	for {
		orders, err := uc.FindAllOrderByCustomerID(customerID)
		if err != nil {
			fmt.Println("Error: ", err)
		}

		var orderName = []string{}
		for _, order := range orders {
			orderName = append(orderName, fmt.Sprintf("%s - %.2f", order.GameTitle, order.GamePrice))
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

		selectedGame := orders[idx]

		for {
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
				isBought = true
			case "Remove":
				err := uc.DeleteOrder(orderID)
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
