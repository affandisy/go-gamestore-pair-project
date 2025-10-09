package cli

import (
	"fmt"
	"gamestore/internal/usecase"

	"github.com/manifoldco/promptui"
)

func gameStore(ucGame *usecase.GameUsecase, ucCat *usecase.CategoryUsecase) {
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

		_, result, _ := menu.Run()

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
			idx, selected, _ := menuGames.Run()
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
						"Buy",
						"Wishlist",
						"Back",
					},
				}

				_, selectedMenuGame, _ := menuGame.Run()

				switch selectedMenuGame {
				case "Buy":
					fmt.Println("Buy")
				case "Wishlist":
					fmt.Println("Wishlist")
				}

				if selectedMenuGame == "Back" {
					break
				}
			}
		}
	}

}
