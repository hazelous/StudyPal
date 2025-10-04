package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"studypal/database"
	"studypal/routers"
)

func main() {
	// Connect DB first
	database.ConnectDatabase()

	app := fiber.New()

	// CORS: allow frontend in browser
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	routers.RouterApp(app)

	app.Get("/health", func(c *fiber.Ctx) error { return c.SendString("ok") })

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("Starting server on :%s ...", port)
	log.Fatal(app.Listen(":8000"))
}
