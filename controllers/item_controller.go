package controllers

import (
	"fmt"
	"procurement-fiber-gorm/config"
	"procurement-fiber-gorm/models"

	"github.com/gofiber/fiber/v2"
)

func GetItems(ctx *fiber.Ctx) error {

	var items []models.Item

	query := config.DB

	query.Select("id", "name", "stock", "price").Find(&items)
	fmt.Println(query.Statement.Vars...)

	if query.Error != nil {
		return ctx.Status(500).JSON(query.Error)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Items retrieved successfully",
		"data":    items,
	})
}

func CreateItem(ctx *fiber.Ctx) error {

	var item models.ItemRequestData

	if err := ctx.BodyParser(&item); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}

	newItem := models.Item{
		Name:  item.Name,
		Stock: item.Stock,
		Price: item.Price,
	}

	query := config.DB.Create(&newItem)

	if query.Error != nil {
		return ctx.Status(500).JSON(query.Error)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Items created successfully",
		"data":    item,
	})
}
