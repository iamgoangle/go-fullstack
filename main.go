package main

import (
	"log"
	"net/http"

	"github.com/iamgoangle/go-fullstack/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"
)

const port string = ":1323"

// BootApp Initial application
func BootApp() {
	e := echo.New()

	// ==============================
	// ==== Database connection  ====
	// ==============================
	db, err := mgo.Dial("localhost:27017")
	if err != nil {
		e.Logger.Fatal(err)
	}

	if err = db.Copy().DB("golang").C("employee").EnsureIndex(mgo.Index{
		Key:    []string{"email"},
		Unique: true,
	}); err != nil {
		log.Fatal(err)
	}

	h := &controllers.Handler{DB: db}

	// apply middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", h.GetUsers)
	e.GET("/user/:userid", h.GetUserByID)
	// e.GET("/products", product.GetProducts)

	// e.POST("/createUser", user.CreateUser)
	e.POST("/addUser", h.AddUser)

	e.DELETE("delete/:userid", h.DeleteUserByID)

	// ==============================
	// ==== Authentication route ====
	// ==============================
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "test" && password == "test" {
			return true, c.String(http.StatusOK, "Hello, Admin!")
		}

		return false, c.String(http.StatusOK, "Invalid authentication")
	}))

	e.Logger.Fatal(e.Start(port))
}

func main() {
	BootApp()
}
