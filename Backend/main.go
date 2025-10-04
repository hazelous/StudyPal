package main

import (
	"studypal/database"
	"studypal/database/migration"
	"studypal/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDatabase()
	migration.Migrate()
	migration.DBSeed(database.DB)

	app := fiber.New()
	routers.RouterApp(app)
	app.Listen(":8080")
}
