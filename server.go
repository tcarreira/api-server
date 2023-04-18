package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tcarreira/api-server/api"
)

var (
	Version = ""
)

func init() {
	if Version == "" {
		if bi, ok := debug.ReadBuildInfo(); ok && bi.Main.Version != "(devel)" {
			Version = bi.Main.Version
		}
	}
	if Version == "" {
		Version = "dev"
	}
}

func main() {
	verInfo := fmt.Sprintf("api-server version: %s", Version)
	flagVersion := flag.Bool("version", false, "Print version information and quit")
	flag.Parse()
	if *flagVersion {
		fmt.Println(verInfo)
		os.Exit(0)
	}

	// Echo instance
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e = api.New(e)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	e.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, verInfo)
	})

	port := "8888"
	if p := os.Getenv("API_PORT"); p != "" {
		port = p
	}

	// Start server
	fmt.Println(verInfo)
	e.Logger.Fatal(e.Start(":" + port))
}
