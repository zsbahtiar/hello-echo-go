package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zsbahtiar/hello-echo-go/internal/handler"
	"github.com/zsbahtiar/hello-echo-go/internal/usecase"
)

type Route struct {
	server *echo.Echo
}

func NewRouter(server *echo.Echo) *Route {
	return &Route{server: server}
}

func (r *Route) Register() {
	userUsecase := usecase.NewUserUsecase()

	userHandler := handler.NewUserHandler(userUsecase)

	r.server.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"isSuccess": true,
			"message":   "success",
		})
	})

	r.server.POST("/users", userHandler.CreateUser)
	r.server.DELETE("/users/:id", userHandler.DeleteUser)
	r.server.GET("/users/:id", userHandler.GetUserByID)
	r.server.GET("/users", userHandler.GetUsers)
}
