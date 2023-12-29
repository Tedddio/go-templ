package main

import (
	"templ-test/routers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
    e.Use(middleware.Logger())
    routers.InitRoutes(e);
	e.GET("/", handleHomePage)

	e.Logger.Fatal(e.Start(":1323"))
}

func handleHomePage(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to the templ-test app")
}

