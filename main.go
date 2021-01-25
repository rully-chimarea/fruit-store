package main

import (
	"fruit-store/controller"
	"fruit-store/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()
	storage.NewDB()
	// Routes
	e.GET("/fruit-store", controller.GetFruits)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
