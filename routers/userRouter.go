package routers

import (
	"templ-test/handlers"
	"templ-test/views"

	"github.com/labstack/echo/v4"
)

func AddUsersRoutes(g *echo.Group) {
	g.GET("", GetUsers)
	g.GET("/:id", GetUser)
}

func GetUsers(c echo.Context) error {
	return handlers.Render(c, views.ShowAllUsers())
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	return handlers.Render(c, views.ShowUser(id))
}
