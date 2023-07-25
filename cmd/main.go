package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1325"))
}

// Handler
func hello(c echo.Context) error {
	name := c.Param("name")
	fmt.Println(name)
	text := "Hello, " + name

	return c.String(http.StatusOK, text)
}
