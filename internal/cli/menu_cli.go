package cli

import (
	"fmt"
	"gamestore/internal/usecase"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/schollz/progressbar/v3"
)

type AppMenu struct {
	CustomerUC      *usecase.CustomerUsecase
	GameUC          *usecase.GameUsecase
	CategoryUC      *usecase.CategoryUsecase
	OrderUC         *usecase.Orderusecase
	PaymentUC       *usecase.Paymentusecase
	ReportUC        *usecase.ReportUsecase
	LibraryUC       *usecase.Libraryusecase
	DownloadedGames map[int]bool
}

func (uc *AppMenu) Run() {
	fmt.Println(`
  ________                                __                        
 /  _____/_____    _____   ____   _______/  |_  ___________   ____  
/   \  ___\__  \  /     \_/ __ \ /  ___/\   __\/  _ \_  __ \_/ __ \ 
\    \_\  \/ __ \|  Y Y  \  ___/ \___ \  |  | (  <_> )  | \/\  ___/ 
 \______  (____  /__|_|  /\___  >____  > |__|  \____/|__|    \___  >
        \/     \/      \/     \/     \/                          \/ 
	`)

	rolesMenu := promptui.Select{
		Label: "Role",
		Items: []string{"Customer", "Admin"},
	}

	_, role, _ := rolesMenu.Run()

	switch role {
	case "Customer":
		customer := AuthMenu(uc.CustomerUC)
		if customer == nil {
			return
		}

		for {
			storeMenu := promptui.Select{
				Label: "Store Menu",
				Items: []string{
					"Store",
					"Orders",
					"Library",
					"Exit",
				},
			}

			_, menu, _ := storeMenu.Run()

			switch menu {
			case "Store":
				gameStore(customer.CustomerID, uc.GameUC, uc.CategoryUC, uc.OrderUC, uc.PaymentUC, uc.LibraryUC)
			case "Orders":
				orderGames(customer.CustomerID, uc.GameUC, uc.OrderUC, uc.PaymentUC, uc.LibraryUC)
			case "Library":
				fmt.Println("This is library")
				libraryGames, err := uc.LibraryUC.FindAllGamesInLibrary(customer.CustomerID)
				if err != nil {
					fmt.Println("Error: ", err)
					continue
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
						game, err := uc.GameUC.FindGameById(selectGame.GameID)
						if err != nil {
							fmt.Println("Error: ", err)
							continue
						}

						var isDownloaded = uc.DownloadedGames[int(selectGame.GameID)]
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
							uc.DownloadedGames[int(selectGame.GameID)] = true
						case "Uninstall":
							bar := progressbar.Default(100)
							for i := 0; i < 100; i++ {
								bar.Add(1)
								time.Sleep(40 * time.Millisecond)
							}
							fmt.Printf("%s is downloaded\n", game.Title)
							uc.DownloadedGames[int(selectGame.GameID)] = false
						case "Back":
							break
						}

						if selectMenuGame == "Back" {
							break
						}
					}

				}

			case "Exit":
				return
			}

			if menu == "Exit" {
				return
			}
		}
	case "Admin":
		adminMenu := promptui.Select{
			Label: "Admin Dashboard",
			Items: []string{"Database", "Report", "Exit"},
		}

		_, menu, _ := adminMenu.Run()
		switch menu {
		case "Database":
			databaseMenu := promptui.Select{
				Label: "Database Dashboard",
				Items: []string{"Customer", "Games", "Category", "Exit"},
			}
			_, menu, _ := databaseMenu.Run()
			switch menu {
			case "Customer":
				customerDatabase(uc.CustomerUC)
				return
			case "Games":
				gameDatabase(uc.GameUC)
				return
			case "Category":
				categoryDatabase(uc.CategoryUC)
				return
			case "Exit":
				return
			}
		case "Report":
			adminReport(uc.ReportUC)
			return
		case "Exit":
			return
		}

	case "Exit":
		return
	}

}
