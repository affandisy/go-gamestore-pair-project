package cli

import (
	"fmt"
	"gamestore/internal/usecase"

	"github.com/manifoldco/promptui"
)

type AppMenu struct {
	CustomerUC *usecase.CustomerUsecase
	GameUC     *usecase.GameUsecase
	CategoryUC *usecase.CategoryUsecase
	OrderUC    *usecase.Orderusecase
	PaymentUC  *usecase.Paymentusecase
}

func (uc *AppMenu) Run() {

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
			gameStore(customer.CustomerID, uc.GameUC, uc.CategoryUC, uc.OrderUC)
		case "Orders":
			orderGames(customer.CustomerID, uc.OrderUC)
			fmt.Println("This is order")
		case "Library":
			fmt.Println("This is library")
		case "Exit":
			return
		}

		if menu == "Exit" {
			return
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
				Items: []string{"Games", "Category", "Exit"},
			}
			_, menu, _ := databaseMenu.Run()
			switch menu {
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
			// adminReport(uc.GameUC)
		case "Exit":
			return
		}

	case "Exit":
		return
	}

}
