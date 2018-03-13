package product

import (
	"net/http"

	"github.com/labstack/echo"
)

type Product struct {
	Title       string  `json:"Title"`
	Description string  `json:"Description"`
	Price       float32 `json:"Price"`
}

type Products []Product

func GetProducts(c echo.Context) error {
	var Products Products
	var IPhoneX = Product{
		Title:       "iPhone X",
		Description: "This is a new generation of iphone",
		Price:       30000.00,
	}

	var SamsungS9 = Product{
		Title:       "Samsung Galaxy S9",
		Description: "Samsung Galaxy Phone",
		Price:       29000.00,
	}

	Products = append(Products, IPhoneX)
	Products = append(Products, SamsungS9)

	return c.JSON(http.StatusOK, Products)
}

// func PostProduct(c echo.Context) {
// 	var p Product

// 	return c.JSON(http.StatusOK, p)
// }
