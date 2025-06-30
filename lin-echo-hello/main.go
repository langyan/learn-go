package main

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			elapsed := time.Since(start)
			log.Printf("Echo: %s %s took %s\n", c.Request().Method, c.Request().URL.Path, elapsed)
			return err
		}
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Hello, World!"})

	})

	e.Logger.Fatal(e.Start(":8080"))
}
