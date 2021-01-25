package main

import (
	"fruit-store/controller"
	"fruit-store/storage"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	e.Validator = &CustomValidator{Validator: validator.New()}

	storage.NewDB()
	// Routes
	e.GET("/fruit-store", controller.GetFruits)
	e.POST("/fruit-store", controller.AddFruit)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// CutstomValidator :
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate : Validate Data
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
