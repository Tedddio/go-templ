package routers

import "github.com/labstack/echo/v4"

func InitRoutes(e *echo.Echo) {
    AddUsersRoutes(e.Group("/users"))
}
