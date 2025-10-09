package cli

import (
	"fmt"
	"gamestore/internal/domain"
	"gamestore/internal/usecase"
	"strings"

	"github.com/manifoldco/promptui"
)

func AuthMenu(uc *usecase.CustomerUsecase) *domain.Customer {
	for {
		authMenuSelect := promptui.Select{
			Label: "Auth",
			Items: []string{
				"Register",
				"Login",
				"Exit",
			},
		}

		_, authMenu, _ := authMenuSelect.Run()

		switch authMenu {
		case "Register":
			for {
				nameInput := promptui.Prompt{Label: "Enter your name"}
				emailInput := promptui.Prompt{Label: "Enter your email"}
				passwordInput := promptui.Prompt{Label: "Enter your password", Mask: '*'}

				name, _ := nameInput.Run()
				name = strings.TrimSpace(name)
				email, _ := emailInput.Run()
				email = strings.TrimSpace(email)
				password, _ := passwordInput.Run()
				password = strings.TrimSpace(password)
				// Cek apakah nama email dan passwordnya tidak kosong
				if name == "" || email == "" || password == "" {
					fmt.Println("Masukkan nama, email, atau password")
					continue
				}

				// Cek apakah customer telah ada
				isCustomerExists, err := uc.FindByEmail(email)
				if err != nil {
					fmt.Println("Error: ", err)
					continue
				}
				if isCustomerExists != nil {
					fmt.Println("Customer sudah ada")
					continue
				}

				// Jika customer tidak ada, buat customer baru
				err2 := uc.CreateCustomer(name, email, password)
				if err2 != nil {
					fmt.Println("Error dalam membuat customer: ", err2)
					continue
				}

				fmt.Println("Registrasi berhasil! Silahkan login")
				break
			}
		case "Login":
			for {
				emailInput := promptui.Prompt{Label: "Enter your email"}
				passwordInput := promptui.Prompt{Label: "Enter your password", Mask: '*'}
				email, _ := emailInput.Run()
				email = strings.TrimSpace(email)
				password, _ := passwordInput.Run()
				password = strings.TrimSpace(password)
				// Cek apakah nama email dan passwordnya tidak kosong
				if email == "" || password == "" {
					fmt.Println("Masukkan email, atau password")
					continue
				}

				customer, err := uc.Login(email, password)
				if err != nil {
					fmt.Println("Error: ", err)
					continue
				}
				fmt.Println("Login berhasil, selamat datang", customer.Name)
				return customer
			}

		case "Exit":
			return nil
		}
	}
}
