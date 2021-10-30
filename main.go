package main

import (
	"astroauth-api/database"
	"astroauth-api/routers"
	"fmt"
)

func main() {
	fmt.Println("Initializing database")

	database.InitialMigration()

	fmt.Println("Database success")
	fmt.Println("Initializing router")

	routers.InitializeRouter()
}
