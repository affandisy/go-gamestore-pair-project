package main

import "gamestore/internal/connections"

func main() {
	db := connections.NewConnection()
	defer db.Close()
}

// KATEGORI PENILAIN
// debuggin
// problem solving
// testify, database deployment
