package router

import (
	"myapp/handlers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
    e.POST("/users", handlers.CreateUser)
    e.GET("/users/:id", handlers.GetUser)
    e.PUT("/users/:id", handlers.UpdateUser)
    e.DELETE("/users/:id", handlers.DeleteUser)
}