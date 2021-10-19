package main

import (
	"astroauth-api/database"
	"astroauth-api/routers"
)

func main() {
	database.InitialMigration()
	routers.InitializeRouter()
}
