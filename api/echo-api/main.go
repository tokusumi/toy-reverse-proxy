package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	ec := echo.New()

	ec.Use(middleware.Logger())
	ec.Use(middleware.Recover())

	ec.GET("/api", hello)

	ec.Logger.Fatal(ec.Start(":1387"))
}

func hello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Hello world, Echo Golang API server"})
}
