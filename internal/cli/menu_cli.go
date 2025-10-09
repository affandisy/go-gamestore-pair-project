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
	customer := AuthMenu(uc.CustomerUC)
	if customer == nil {
		return
	}
	for {
		storeMenu := promptui.Select{
			Label: "Store Menu",
			Items: []string{
				"Store",
				"Order",
				"Library",
				"Exit",
			},
		}

		_, menu, _ := storeMenu.Run()

		switch menu {
		case "Store":
			gameStore(uc.GameUC, uc.CategoryUC)
		case "Order":
			fmt.Println("This is order")
		case "Library":
			fmt.Println("This is library")
		case "Exit":
			return
		}

		if menu == "Exit" {
			return
		}

	}
}
