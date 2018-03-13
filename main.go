package main

import (
	"net/http"

	"github.com/iamgoangle/go-fullstack/controllers/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const port string = "localhost:3000"

// BootApp Initial application
func BootApp() {
	e := echo.New()

	// apply middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/createUser", user.CreateUser)
	e.GET("/users/:id", user.GetUser)
	e.GET("users", user.GetUsers)
	e.GET("/show", user.QueryUser)

	// e.PUT("/users/:id", updateUser)

	// e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(port))
}

func main() {
	BootApp()
}
