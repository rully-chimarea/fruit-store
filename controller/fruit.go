package controller

import (
	"fruit-store/model"

	db "fruit-store/storage"

	"fruit-store/utils"

	"net/http"

	"github.com/labstack/echo/v4"
)

func GetFruits(c echo.Context) error {
	fruit, err := GetRepoFruits()
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, fruit)
}

// AddFruit : Create Fruit
func AddFruit(c echo.Context) error {
	type RequestBody struct {
		// Id        uint    `json:"id" validate:"required"`
		FruitName string  `json:"fruit-name" validate:"required"`
		FruitType string  `json:"fruit-type" validate:"required"`
		Price     float64 `json:"price" validate:"required"`
		Quantity  uint    `json:"quantity"`
	}

	var body RequestBody

	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	db := db.GetDBInstance()
	fruit := model.Fruits{
		// Id:        body.Id,
		FruitName: body.FruitName,
		FruitType: body.FruitType,
		Price:     body.Price,
		Quantity:  body.Quantity,
	}
	db.Create(&fruit)

	return c.JSON(http.StatusCreated, fruit)
}

func GetRepoFruits() ([]model.Fruits, error) {
	db := db.GetDBInstance()
	fruits := []model.Fruits{}

	if err := db.Find(&fruits).Error; err != nil {
		return nil, err
	}

	return fruits, nil
}
