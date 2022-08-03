package main

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/lizzy-j/dummy-api/controllers"
	"github.com/lizzy-j/dummy-api/database"
)

func welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to the vehicle api")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", welcome)
	app.Get("api/vehicles", controllers.DisplayVehicles)
	app.Post("/api/vehicle", controllers.CreateVehicle)
	app.Put("/api/vehicle/:id", controllers.UpdateVehicle)
	app.Delete("api/vehicle", controllers.DeleteVehicle)
}

func main() {
	app := fiber.New()
	database.ConnectDb()
	setupRoutes(app)

	app.Listen(":3001")
}
