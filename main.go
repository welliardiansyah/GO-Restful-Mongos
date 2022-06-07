package main

import (
	"restfull/configs"
	"restfull/routes"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // DATABASE
    configs.ConnectDB()

	// ROUTES
	routes.UserRoute(app)

    app.Listen(":6000")
}