package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lizzy-j/dummy-api/database"
	models "github.com/lizzy-j/dummy-api/models"
)

//this is the serializer
type Vehicle struct {
	ID        uint64 `json:"id"`
	Make      string `json:"make"`
	ModelType string `json:"model_type"`
}

func CreateResponseVehicle(vehicleModel models.Vehicle) Vehicle {
	return Vehicle{ID: vehicleModel.ID, Make: vehicleModel.Make, ModelType: vehicleModel.ModelType}
}

func CreateVehicle(ctx *fiber.Ctx) error {
	var vehicle models.Vehicle

	if err := ctx.BodyParser(&vehicle); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&vehicle)
	responseVehicle := CreateResponseVehicle(vehicle)
	return ctx.Status(200).JSON(responseVehicle)
}

func UpdateVehicle(ctx *fiber.Ctx) error {
	vehicle := []models.Vehicle{}
	update := new(models.Vehicle)
	if err := ctx.BodyParser(update); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}
	database.Database.Db.Model(&vehicle).Where("id = ?", update.ID).Save(update)
	return ctx.Status(201).JSON("vehicle has been updated")
}

func DeleteVehicle(ctx *fiber.Ctx) error {
	vehicle := []models.Vehicle{}
	deletedVehicle := new(models.Vehicle)
	if err := ctx.BodyParser(deletedVehicle); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	database.Database.Db.Model(&vehicle).Where("id = ?", deletedVehicle.ID).Delete(&vehicle)
	return ctx.Status(204).JSON("vehicle has been successfully deleted")
}

func DisplayVehicles(ctx *fiber.Ctx) error {
	vehicles := []models.Vehicle{}
	database.Database.Db.Model(&vehicles).Order("ID asc").Find(&vehicles)

	return ctx.JSON(fiber.Map{
		"code": 200,
		"message": "success",
		"data": vehicles
	})
}
