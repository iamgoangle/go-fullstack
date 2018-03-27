package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/iamgoangle/go-fullstack/config"
	"github.com/iamgoangle/go-fullstack/models"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

type User models.User
type Users []User

// GetUserByID
// Query user by userid
func (h *Handler) GetUserByID(c echo.Context) (err error) {
	userID := c.Param("userid")

	db := h.DB.Clone()
	defer db.Close()

	var user User
	if err = db.Copy().
		DB(config.DBCONFIG.Dbname).
		C(config.DBCONFIG.Collection).
		Find(bson.M{"userid": bson.ObjectIdHex(userID)}).
		One(&user); err != nil {
		return
	}

	return c.JSON(http.StatusOK, user)
}

// GetUsers
// Get User of Users
func (h *Handler) GetUsers(c echo.Context) (err error) {
	log.Println("GetUsers")

	db := h.DB.Clone()
	defer db.Close()

	var users Users
	if err = db.Copy().
		DB(config.DBCONFIG.Dbname).
		C(config.DBCONFIG.Collection).
		Find(bson.M{}).
		All(&users); err != nil {
		return
	}

	return c.JSON(http.StatusOK, users)
}

// AddUser
// Adding a new user
func (h *Handler) AddUser(c echo.Context) (err error) {
	log.Println("AddUser")

	u := &models.User{
		ID:         bson.NewObjectId(),
		UserID:     bson.NewObjectId(),
		CreateDate: time.Now().Format("2006-01-02"),
	}

	if err = c.Bind(u); err != nil {
		return
	}

	db := h.DB.Clone()
	defer db.Close()

	if err = db.Copy().
		DB(config.DBCONFIG.Dbname).
		C(config.DBCONFIG.Collection).
		Insert(u); err != nil {
		return
	}

	return c.JSON(http.StatusCreated, u)
}

// DeleteUserByID
// Delete a user by userid
func (h *Handler) DeleteUserByID(c echo.Context) (err error) {
	log.Println("DeleteUserByID")

	userID := c.Param("userid")

	db := h.DB.Clone()
	db.Close()

	if err = db.Copy().
		DB(config.DBCONFIG.Dbname).
		C(config.DBCONFIG.Collection).
		Remove(bson.M{"userid": bson.ObjectIdHex(userID)}); err != nil {
		return
	}

	return c.NoContent(http.StatusNoContent)
}
