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
	for {
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
				continue
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
				continue
			}

			if len(customers) == 0 {
				fmt.Println("Tidak ada game yang tersedia")
				continue
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
				continue
			}

			selectedCustomers := customers[idx]

			confirmPrompt := promptui.Prompt{
				Label:     fmt.Sprintf("Apakah anda yakin ingin menghapus Customer %s ini? (Y, N)", selectedCustomers.Name),
				IsConfirm: true,
			}

			confirm, _ := confirmPrompt.Run()
			if confirm != "y" && confirm != "Y" {
				fmt.Println("Dibatalkan")
				continue
			}

			if err := uc.DeleteCustomer(selectedCustomers.CustomerID); err != nil {
				log.Println("Error delete customer: ", err)
				continue
			}

			fmt.Printf("Customer %s dihapus. \n", selectedCustomers.Name)
		case "Exit":
			return
		}

	}
}

func gameDatabase(uc *usecase.GameUsecase) {
	for {
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

func adminReport(reportUC *usecase.ReportUsecase) {
	for {
		prompt := promptui.Select{
			Label: "Pilih Report",
			Items: []string{
				"Histori Pembelian Customer",
				"Game Terlaris",
				"Total Pendapatan",
				"Summary",
				"Exit",
			},
		}

		_, choice, err := prompt.Run()
		if err != nil {
			fmt.Println("Prompt gagal:", err)
			return
		}

		switch choice {
		case "Histori Pembelian Customer":
			history, err := reportUC.GetPurchaseHistory()
			if err != nil {
				fmt.Println("Error", err)
				return
			}

			table := tablewriter.NewTable(os.Stdout)
			table.Header("OrderID", "Customer", "Email", "Game", "Harga", "Status", "Tanggal")

			for _, h := range history {
				table.Append([]string{
					fmt.Sprintf("%d", h.OrderID),
					h.CustomerName,
					h.CustomerEmail,
					h.GameTitle,
					fmt.Sprintf("Rp %.2f", h.Price),
					h.PaymentStatus,
					h.OrderDate.Format("2006-01-02"),
				})
			}
			table.Render()
		case "Game Terlaris":
			best, err := reportUC.GetBestSellingGames()
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			table := tablewriter.NewTable(os.Stdout)
			table.Header("Game", "Total Terjual", "Total Pendapatan")

			for _, b := range best {
				table.Append([]string{
					b.GameName,
					fmt.Sprintf("%d", b.TotalTerjual),
					fmt.Sprintf("Rp %.2f", b.TotalPendapatan),
				})
			}
			table.Render()
		case "Total Pendapatan":
			total, err := reportUC.GetRevenueSummary()
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			table := tablewriter.NewTable(os.Stdout)
			table.Header("Kategori", "Nilai")
			table.Append([]string{"Total Revenue", fmt.Sprintf("Rp %.2f", total.TotalRevenue)})
			table.Append([]string{"Outstanding Bills", fmt.Sprintf("Rp %.2f", total.OutstandingBills)})
			table.Append([]string{"Daily Income", fmt.Sprintf("Rp %.2f", total.DailyIncome)})
			table.Render()
		case "Summary":
			summary, err := reportUC.GetAdminSummary()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			table := tablewriter.NewTable(os.Stdout)
			table.Header("Metrix", "Jumlah")
			table.Append([]string{"Total Customer", fmt.Sprintf("%d", summary.TotalCustomers)})
			table.Append([]string{"Total Game", fmt.Sprintf("%d", summary.TotalGames)})
			table.Append([]string{"Total Order", fmt.Sprintf("%d", summary.TotalOrders)})
			table.Append([]string{"Total Pembayaran", fmt.Sprintf("%d", summary.TotalPayments)})
			table.Append([]string{"Total Pendapatan", fmt.Sprintf("Rp %.2f", summary.TotalRevenue)})
			table.Render()
		case "Exit":
			return
		}
	}
}
