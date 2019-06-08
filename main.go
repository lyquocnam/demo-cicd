package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
	"log"
)

func main() {
	e := echo.New()
	e.AutoTLSManager.Cache = autocert.DirCache(".cache")
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, echo.Map{
			"message": "ok",
		})
	})
	log.Fatalln(e.StartTLS(":443", "localhost.pem", "localhost-key.pem"))
}


