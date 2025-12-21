package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "hello",
		})
	})

	e.GET("/request", func(c echo.Context) error {
		name := c.QueryParam("name")

		if name == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "missing required parameter 'name'",
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "received request",
			"name":    name,
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
