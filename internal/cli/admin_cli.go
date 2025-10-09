package cli

import (
	"fmt"
	"gamestore/internal/usecase"
	"log"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/olekukonko/tablewriter"
)

func adminDatabase(uc *usecase.GameUsecase) {
	adminDatabaseMenu := promptui.Select{
		Label: "Select Action",
		Items: []string{
			"Publish Game",
			"All Games",
			"Update Game",
			"Delete Game",
			"Exit",
		},
	}

	_, menu, _ := adminDatabaseMenu.Run()

	switch menu {
	case "Publish Game":
		for {
			titleInput := promptui.Prompt{
				Label: "Title",
			}
			title, _ := titleInput.Run()

			categoryInput := promptui.Prompt{
				Label: "1.Action 2.Adventure 3.RPG 4.Strategy 5.Sports 6.Racing 7.Shooter 8.Puzzle  9.Horror 10.Fighting",
			}
			categoryString, _ := categoryInput.Run()
			category, err := strconv.Atoi(categoryString)
			if err != nil {
				fmt.Println("Tolong masukkan nomor id dari kategorinya saja")
				continue
			}

			priceInput := promptui.Prompt{
				Label: "Price",
			}
			priceString, _ := priceInput.Run()
			price, err := strconv.ParseFloat(priceString, 64)
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}

			if categoryString == "" || title == "" || priceString == "" {
				fmt.Println("Masukkan title, category, dan input")
				continue
			}

			err2 := uc.CreateGame(int64(category), title, price)
			if err2 != nil {
				fmt.Println("Error: ", err2)
				continue
			}

			fmt.Println("Berhasil publish game!")
			break
		}

	case "All Games":
		games, err := uc.FindAllGame()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		table := tablewriter.NewTable(os.Stdout)
		table.Header("GameID", "Category", "Title", "Price", "Created_At", "Updated_At")
		table.Bulk(games)
		table.Render()
	case "Update Game":
		games, err := uc.FindAllGame()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		if len(games) == 0 {
			fmt.Println("No games available")
			return
		}

		gameItems := []string{}
		for _, g := range games {
			gameItems = append(gameItems, fmt.Sprintf("[%d] %s", g.GameID, g.Title))
		}

		selectPrompt := promptui.Select{
			Label: "Select Game to Update",
			Items: gameItems,
		}

		idx, _, err := selectPrompt.Run()
		if err != nil {
			log.Println("Prompt Failed", err)
			return
		}

		selectedGame := games[idx]

		titledPrompt := promptui.Prompt{
			Label:   fmt.Sprintf("Judul Baru: (%s)", selectedGame.Title),
			Default: selectedGame.Title,
		}
		newTitle, _ := titledPrompt.Run()

		pricePrompt := promptui.Prompt{
			Label:   fmt.Sprintf("Harga Baru: (%.2f)", selectedGame.Price),
			Default: fmt.Sprintf("%.2f", selectedGame.Price),
		}
		newPriceStr, _ := pricePrompt.Run()
		newPrice, _ := strconv.ParseFloat(newPriceStr, 64)

		categoryPrompt := promptui.Prompt{
			Label:   fmt.Sprintf("New Category ID (current: %d)", selectedGame.CategoryID),
			Default: fmt.Sprintf("%d", selectedGame.CategoryID),
		}
		newCategoryStr, _ := categoryPrompt.Run()
		newCategoryID, _ := strconv.ParseInt(newCategoryStr, 10, 64)

		selectedGame.Title = newTitle
		selectedGame.Price = newPrice
		selectedGame.CategoryID = newCategoryID

		if err := uc.UpdateGame(&selectedGame); err != nil {
			fmt.Println("Error update game:", err)
			return
		}
		fmt.Println("Game Update selesai")
	case "Exit":
		return
	}
}

func adminReport(uc *usecase.GameUsecase) {

}
