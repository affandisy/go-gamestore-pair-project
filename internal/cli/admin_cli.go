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

func customerDatabase(uc *usecase.CustomerUsecase) {
	customerDatabaseMenu := promptui.Select{
		Label: "Select Action",
		Items: []string{
			"All Customers",
			"Delete Customer",
			"Exit",
		},
	}

	_, menu, _ := customerDatabaseMenu.Run()

	switch menu {
	case "All Customers":
		customers, err := uc.FindAllCustomer()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		table := tablewriter.NewTable(os.Stdout)
		table.Header("CustomerID", "Name", "Email", "Password", "CreatedAt", "UpdatedAt")

		for _, c := range customers {
			table.Append([]string{
				fmt.Sprintf("%d", c.CustomerID),
				c.Name,
				c.Email,
				"********",
				c.CreatedAt.Format("2006-01-02 15:04:05"),
				c.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}
		table.Render()
	case "Delete Customer":
		customers, err := uc.FindAllCustomer()
		if err != nil {
			fmt.Println("Error: gagal mengambil game", err)
			return
		}

		if len(customers) == 0 {
			fmt.Println("Tidak ada game yang tersedia")
			return
		}

		customerItems := []string{}

		for _, c := range customers {
			customerItems = append(customerItems, fmt.Sprintf("[%d] %s", c.CustomerID, c.Name))
		}

		selectPrompt := promptui.Select{
			Label: "Pilih customer untuk dihapus",
			Items: customerItems,
		}

		idx, _, err := selectPrompt.Run()
		if err != nil {
			log.Println("Prompt gagal", err)
			return
		}

		selectedCustomers := customers[idx]

		confirmPrompt := promptui.Prompt{
			Label:     fmt.Sprintf("Apakah anda yakin ingin menghapus Customer %s ini? (Y, N)", selectedCustomers.Name),
			IsConfirm: true,
		}

		confirm, _ := confirmPrompt.Run()
		if confirm != "y" && confirm != "Y" {
			fmt.Println("Dibatalkan")
			return
		}

		if err := uc.DeleteCustomer(selectedCustomers.CustomerID); err != nil {
			log.Println("Error delete customer: ", err)
			return
		}

		fmt.Printf("Customer %s dihapus. \n", selectedCustomers.Name)
	case "Exit":
		return
	}
}

func gameDatabase(uc *usecase.GameUsecase) {
	gameDatabaseMenu := promptui.Select{
		Label: "Select Action",
		Items: []string{
			"Publish Game",
			"All Games",
			"Update Game",
			"Delete Game",
			"Exit",
		},
	}

	_, menu, _ := gameDatabaseMenu.Run()

	switch menu {
	case "Publish Game":
		for {
			titleInput := promptui.Prompt{
				Label: "Title",
			}
			title, err := titleInput.Run()
			if err != nil {
				fmt.Println("Tolong masukkan title")
				return
			}

			categoryInput := promptui.Prompt{
				Label: "1.Action 2.Adventure 3.RPG 4.Strategy 5.Sports 6.Racing 7.Shooter 8.Puzzle  9.Horror 10.Fighting",
			}
			categoryString, _ := categoryInput.Run()
			category, err := strconv.Atoi(categoryString)
			if err != nil {
				fmt.Println("Tolong masukkan nomor id dari kategorinya saja")
				return
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
			fmt.Println("Tidak ada game yang tersedia")
			return
		}

		gameItems := []string{}
		for _, g := range games {
			gameItems = append(gameItems, fmt.Sprintf("[%d] %s", g.GameID, g.Title))
		}

		selectPrompt := promptui.Select{
			Label: "Pilih game untuk diupdate",
			Items: gameItems,
		}

		idx, _, err := selectPrompt.Run()
		if err != nil {
			log.Println("Prompt Gagal", err)
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
			Label:   fmt.Sprintf("ID Kategori Baru (current: %d)", selectedGame.CategoryID),
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
	case "Delete Game":
		games, err := uc.FindAllGame()
		if err != nil {
			fmt.Println("Error: gagal mengambil game", err)
			return
		}

		if len(games) == 0 {
			fmt.Println("Tidak ada game yang tersedia")
			return
		}

		gameItems := []string{}

		for _, g := range games {
			gameItems = append(gameItems, fmt.Sprintf("[%d] %s", g.GameID, g.Title))
		}

		selectPrompt := promptui.Select{
			Label: "Pilih game untuk dihapus",
			Items: gameItems,
		}

		idx, _, err := selectPrompt.Run()
		if err != nil {
			log.Println("Prompt gagal", err)
			return
		}

		selectedGame := games[idx]

		confirmPrompt := promptui.Prompt{
			Label:     fmt.Sprintf("Apakah anda yakin ingin menghapus game %s ini? (Y, N)", selectedGame.Title),
			IsConfirm: true,
		}

		confirm, _ := confirmPrompt.Run()
		if confirm != "y" && confirm != "Y" {
			fmt.Println("Dibatalkan")
			return
		}

		if err := uc.DeleteGame(selectedGame.GameID); err != nil {
			log.Println("Error delete game: ", err)
			return
		}

		fmt.Printf("Game %s dihapus. \n", selectedGame.Title)
	case "Exit":
		return
	}
}

// CLI Category Menu
func categoryDatabase(uc *usecase.CategoryUsecase) {
	for {
		categoryDatabaseMenu := promptui.Select{
			Label: "Select Action",
			Items: []string{
				"Tambah Category",
				"Semua Category",
				"Update Category",
				"Delete Category",
				"Exit",
			},
		}

		_, menu, _ := categoryDatabaseMenu.Run()

		switch menu {
		case "Tambah Category":
			fmt.Println("=== Tambah Kategori ===")

			namaPrompt := promptui.Prompt{
				Label: "Nama Kategori",
			}

			name, err := namaPrompt.Run()
			if err != nil || name == "" {
				fmt.Println("Tolong masukkan nama kategori yang valid")
				continue
			}

			if err := uc.CreateCategory(name); err != nil {
				fmt.Println("Error", err)
				continue
			}

			fmt.Println("Berhasil Menambahkan Kategori: ", name)

		case "Semua Category":
			categories, err := uc.FindAllCategories()
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			table := tablewriter.NewTable(os.Stdout)
			table.Header("CategoryID", "Name")
			table.Bulk(categories)
			table.Render()
		case "Update Category":
			categories, err := uc.FindAllCategories()
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			if len(categories) == 0 {
				fmt.Println("Tidak ada category yang tersedia")
				return
			}

			categoryItems := []string{}
			for _, c := range categories {
				categoryItems = append(categoryItems, fmt.Sprintf("[%d] %s", c.CategoryID, c.Name))
			}

			selectPrompt := promptui.Select{
				Label: "Pilih kategori untuk diupdate",
				Items: categoryItems,
			}

			idx, _, err := selectPrompt.Run()
			if err != nil {
				log.Println("Prompt gagal", err)
				return
			}

			selectedCategory := categories[idx]

			namePrompt := promptui.Prompt{
				Label:   fmt.Sprintf("Nama kategori Baru: (%s)", selectedCategory.Name),
				Default: selectedCategory.Name,
			}
			newName, _ := namePrompt.Run()

			selectedCategory.Name = newName

			if err := uc.UpdateCategory(&selectedCategory); err != nil {
				fmt.Println("Error update game:", err)
				return
			}
			fmt.Println("Category Update selesai")
		case "Delete Category":
			categories, err := uc.FindAllCategories()
			if err != nil {
				fmt.Println("Error: Gagal mengambil data", err)
				return
			}

			if len(categories) == 0 {
				fmt.Println("Tidak ada category yang tersedia")
				return
			}

			categoryItems := []string{}

			for _, c := range categories {
				categoryItems = append(categoryItems, fmt.Sprintf("[%d] %s", c.CategoryID, c.Name))
			}

			selectPrompt := promptui.Select{
				Label: "Pilih category untuk dihapus",
				Items: categoryItems,
			}

			idx, _, err := selectPrompt.Run()
			if err != nil {
				log.Println("Prompt gagal", err)
				return
			}

			selectedCategory := categories[idx]

			confirmPrompt := promptui.Prompt{
				Label:     fmt.Sprintf("Apakah anda ingin menghapus kategori: %s ini? (Y/N)", selectedCategory.Name),
				IsConfirm: true,
			}

			confirm, _ := confirmPrompt.Run()
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Dibatalkan")
				return
			}

			if err := uc.DeleteCategory(selectedCategory.CategoryID); err != nil {
				log.Println("Error delete game: ", err)
				return
			}

			fmt.Printf("Kategori %s telah dihapus. \n", selectedCategory.Name)
		case "Exit":
			return
		}

	}

}

func adminReport(uc *usecase.GameUsecase) {

}
