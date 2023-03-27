package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type userHandler struct{}
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

type UserHandler interface {
	CreateUser(ctx echo.Context) error
	GetUserByID(ctx echo.Context) error
	GetUsers(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
}

func NewUserHandler() UserHandler {
	return &userHandler{}
}

func (h *userHandler) CreateUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	uuid := uuid.New()
	user.ID = uuid.String()
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"isSuccess": true,
		"message":   "success",
		"data":      user,
	})
}

func (h *userHandler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	user := &User{
		ID:    id,
		Name:  "bob",
		Email: "bob@mail.com",
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"isSuccess": true,
		"message":   "success",
		"data":      user,
	})
}

func (h *userHandler) GetUsers(c echo.Context) error {
	users := []*User{
		{
			ID:    uuid.New().String(),
			Name:  "bob",
			Email: "bob@mail.com",
		},
		{
			ID:    uuid.New().String(),
			Name:  "alice",
			Email: "alice@mail.com",
		},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"isSuccess": true,
		"message":   "success",
		"data":      users,
	})
}

func (h *userHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"isSuccess": true,
		"message":   "success",
		"data": map[string]interface{}{
			"id": id,
		},
	})
}
