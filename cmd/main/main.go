package main

import (
	"os"

	"github.com/tudemaha/data-collector/pkg"
	"github.com/tudemaha/data-collector/routes"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	routes.Router()

	pkg.DatabaseConnection()
	pkg.StartServer(os.Getenv("PORT"))
}
