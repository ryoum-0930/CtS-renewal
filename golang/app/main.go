package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func router(e *echo.Echo) {
	e.GET("/", helloWorld)
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func main() {
	e := echo.New()
	router(e)
	e.Logger.Fatal(e.Start(":8080"))
}
