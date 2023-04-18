package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tcarreira/api-server/api"
)

var (
	Version = "dev"
	commit  = ""
	date    = ""
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e = api.New(e)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	e.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintln(Version, commit, date))
	})

	// Start server
	port := "8888"
	if p := os.Getenv("API_PORT"); p != "" {
		port = p
	}
	e.Logger.Fatal(e.Start(":" + port))
}
