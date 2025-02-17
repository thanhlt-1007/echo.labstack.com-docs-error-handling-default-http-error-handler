package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func authStatusMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			status := c.QueryParam("status")
			if status != "ok" {
				return echo.NewHTTPError(
					http.StatusUnauthorized,
					"Please set stautus query's value ist ok",
				)
			}

			return next(c)
		}
	}
}

func getPingHandler(context echo.Context) error {
	return context.String(
		http.StatusOK,
		"OK",
	)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(authStatusMiddleware())
	e.GET("/ping", getPingHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
