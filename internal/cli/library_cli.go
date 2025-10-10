package cli

import (
	"fmt"
	"gamestore/internal/domain"
	"gamestore/internal/usecase"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/schollz/progressbar/v3"
)

func libraryGames(customer *domain.Customer, ucLib *usecase.Libraryusecase, ucGame *usecase.GameUsecase, ucIsDownloaded map[int]bool) error {
	fmt.Println("Library")
	libraryGames, err := ucLib.FindAllGamesInLibrary(customer.CustomerID)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	gameNames := []string{}
	for _, game := range libraryGames {
		gameNames = append(gameNames, fmt.Sprintf("%s", game.GameTitle))
	}
	gameNames = append(gameNames, "Back")

	gameMenu := promptui.Select{
		Label: "Library",
		Items: gameNames,
	}

	for {
		idx, selectedGameMenu, _ := gameMenu.Run()
		if selectedGameMenu == "Back" {
			break
		}
		selectGame := libraryGames[idx]

		for {
			game, err := ucGame.FindGameById(selectGame.GameID)
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}

			var isDownloaded = ucIsDownloaded[int(selectGame.GameID)]
			items := []string{}
			if !isDownloaded {
				items = append(items, "Download")
			} else if isDownloaded {
				items = append(items, "Uninstall")
			}

			items = append(items, "Back")

			menuGame := promptui.Select{
				Label: game.Title,
				Items: items,
			}

			_, selectMenuGame, _ := menuGame.Run()

			switch selectMenuGame {
			case "Download":
				bar := progressbar.Default(100)
				for i := 0; i < 100; i++ {
					bar.Add(1)
					time.Sleep(40 * time.Millisecond)
				}
				fmt.Printf("%s is downloaded\n", game.Title)
				ucIsDownloaded[int(selectGame.GameID)] = true
			case "Uninstall":
				bar := progressbar.Default(100)
				for i := 0; i < 100; i++ {
					bar.Add(1)
					time.Sleep(40 * time.Millisecond)
				}
				fmt.Printf("%s is downloaded\n", game.Title)
				ucIsDownloaded[int(selectGame.GameID)] = false
			case "Back":
				break
			}

			if selectMenuGame == "Back" {
				break
			}
		}

	}

	return nil
}
