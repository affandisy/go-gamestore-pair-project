package cli

import (
	"fmt"
	"gamestore/internal/usecase"
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
	case "Exit":
		return
	}
}

func adminReport(uc *usecase.GameUsecase) {

}
