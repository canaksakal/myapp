package handlers

import (
	"net/http"
	"myapp/models"
	"github.com/labstack/echo/v4"
)

var users = map[string]*model.User{}

func CreateUser(c echo.Context) error {
    user := new(model.User)
    if err := c.Bind(user); err != nil {
        return err
    }
    users[user.ID] = user
    return c.JSON(http.StatusCreated, user)
}

func GetUser(c echo.Context) error {
    id := c.Param("id")
    user, exists := users[id]
    if !exists {
        return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
    }
    return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
    id := c.Param("id")
    user := users[id]
    if user == nil {
        return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
    }

    if err := c.Bind(user); err != nil {
        return err
    }
    users[id] = user
    return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
    id := c.Param("id")
    if _, exists := users[id]; !exists {
        return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
    }
    delete(users, id)
    return c.NoContent(http.StatusNoContent)
}

