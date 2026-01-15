package controllers

import (
	"fmt"
	"procurement-fiber-gorm/config"
	"procurement-fiber-gorm/models"
	"procurement-fiber-gorm/validations"

	"github.com/go-playground/validator/v10"
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

	var item validations.ItemRequestValidation

	if err := ctx.BodyParser(&item); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}

	validate := *validator.New()

	errValidate := validate.Struct(item)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   errValidate.(validator.ValidationErrors).Error(),
		})
	}

	newItem := models.ItemRequestData{
		Name:        item.Name,
		Description: item.Description,
		Stock:       item.Stock,
		Price:       item.Price,
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

func UpdateItem(ctx *fiber.Ctx) error {

	var item validations.ItemRequestValidation

	if err := ctx.BodyParser(&item); err != nil {

		return ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}

	validate := *validator.New()

	errValidate := validate.Struct(item)

	if errValidate != nil {

		return ctx.Status(400).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   errValidate.(validator.ValidationErrors).Error(),
		})
	}

	updateItem := models.ItemRequestData{
		Name:        item.Name,
		Description: item.Description,
		Stock:       item.Stock,
		Price:       item.Price,
	}

	query := config.DB.Model(&models.ItemRequestData{}).Where("id = ?", ctx.Params("id")).Updates(updateItem)

	if query.Error != nil {
		return ctx.Status(500).JSON(query.Error)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Item updated successfully",
		"data":    updateItem,
	})
}

func DeleteItem(ctx *fiber.Ctx) error {

	var params map[string]string = ctx.AllParams()

	queryFindItem := config.DB.Select("id").Find()

	fmt.Println(queryFindItem.Statement.Vars...)

	if queryFindItem.Error != nil {
		return ctx.Status(500).JSON(queryFindItem.Error)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Item deleted successfully",
		"data":    queryFindItem,
	})

}
