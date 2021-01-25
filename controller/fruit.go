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

func GetRepoFruits() ([]model.Fruits, error) {
	db := db.GetDBInstance()
	fruits := []model.Fruits{}

	if err := db.Find(&fruits).Error; err != nil {
		return nil, err
	}

	return fruits, nil
}
