package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func main() {
	server := echo.New()

	server.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"isSucces": true,
			"messsage": "health",
		})
	})

	server.POST("/users", createUser)

	if err := server.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}

func createUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	uuid := uuid.New()
	user.ID = uuid.String()
	return c.JSON(http.StatusCreated, user)
}
