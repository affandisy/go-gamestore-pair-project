package cli

import (
	"fmt"
	"gamestore/internal/usecase"

	"github.com/manifoldco/promptui"
)

func gameStore(customerID int64, ucGame *usecase.GameUsecase, ucCat *usecase.CategoryUsecase, ucOrder *usecase.Orderusecase, ucPay *usecase.Paymentusecase, ucLib *usecase.Libraryusecase) {
	for {
		categories, err := ucCat.FindAllCategories()
		if err != nil {
			fmt.Println("Error: ", err)
		}

		items := []string{}
		for _, c := range categories {
			items = append(items, c.Name)
		}
		items = append(items, "Back")

		menu := promptui.Select{
			Label: "Pilih kategori",
			Items: items,
		}

		_, result, err := menu.Run()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		if result == "Back" {
			break
		}

		var categoryID int64
		for _, c := range categories {
			if c.Name == result {
				categoryID = c.CategoryID
			}
		}

		for {
			// menu game berdasarkan category
			games, err := ucGame.FindGameByCategoryID(categoryID)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			gameItems := []string{}
			for _, g := range games {
				gameItems = append(gameItems, fmt.Sprintf("%s - Rp%.2f", g.Title, g.Price))
			}
			gameItems = append(gameItems, "Back")

			menuGames := promptui.Select{
				Label: "Pick a game you want",
				Items: gameItems,
			}
			idx, selected, err := menuGames.Run()
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}
			if selected == "Back" {
				break
			}

			selectedGame := games[idx]

			for {
				game, err := ucGame.FindGameById(selectedGame.GameID)
				if err != nil {
					fmt.Println("Error", err)
				}

				menuGame := promptui.Select{
					Label: game.Title,
					Items: []string{
						"Buy now",
						"Add To Orders Cart",
						"Back",
					},
				}

				var isPaid = false
				var isAdded = false

				_, selectedMenuGame, err := menuGame.Run()
				if err != nil {
					fmt.Println("Error: ", err)
					continue
				}
				switch selectedMenuGame {
				case "Buy now":
					err := payOneGame(customerID, ucOrder, ucPay, ucLib, game)
					if err != nil {
						fmt.Println("Error: ", err)
						continue
					}
					isPaid = true
				case "Add To Orders Cart":
					order, err := ucOrder.CreateOrder(customerID, game.GameID)
					if err != nil {
						fmt.Println("Error: ", err)
						continue
					}

					fmt.Printf("%s berhasil dimasukkan ke orders dengan id %d\n", game.Title, order.OrderID)
					isAdded = true
				}

				if isPaid || isAdded || selectedMenuGame == "Back" {
					break
				}
			}
		}
	}

}
