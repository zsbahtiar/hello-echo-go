package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/zsbahtiar/hello-echo-go/internal/core/entity"
	"github.com/zsbahtiar/hello-echo-go/internal/core/usecase"
)

type userHandler struct {
	userUc usecase.UserUsecase
}

type UserHandler interface {
	CreateUser(ctx echo.Context) error
	GetUserByID(ctx echo.Context) error
	GetUsers(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
}

func NewUserHandler(userUc usecase.UserUsecase) UserHandler {
	return &userHandler{userUc: userUc}
}

func (h *userHandler) CreateUser(c echo.Context) error {
	user := new(entity.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	resp, err := h.userUc.CreateUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"isSuccess": true,
			"message":   "failed to create user",
			"error":     err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"isSuccess": true,
		"message":   "success create user",
		"data":      resp,
	})

}

func (h *userHandler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	user := &entity.User{
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
	users := []*entity.User{
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
