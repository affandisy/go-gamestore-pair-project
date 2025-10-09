package cli

import (
	"fmt"
	"gamestore/internal/usecase"

	"github.com/manifoldco/promptui"
)

type AppMenu struct {
	CustomerUC *usecase.CustomerUsecase
	GameUC     *usecase.GameUsecase
	OrderUC    *usecase.Orderusecase
	PaymentUC  *usecase.Paymentusecase
}

func (uc *AppMenu) Run() {
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
		fmt.Println("This is store")
	case "Order":
		fmt.Println("This is order")
	case "Library":
		fmt.Println("This is library")
	case "Exit":
		break
	}
}
