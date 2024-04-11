package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	databaseName = "mydatabase"
	username = "myuser"
	password = "mypassword"
)
var db *sql.DB

type Product struct {
	ID int	`json:"id"`
	Name string `json:"name"`
	Price int 	`json:"price"`
}



func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	sdb, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	db = sdb

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("product" , getAllProductHandler)
	app.Get("/product/:id" , getProductHandler)
	app.Post("/product" , createProductHandler)
	app.Put("/product/:id" , updateProductHandler)
	app.Delete("/product/:id" ,deleteProductHandler)

	app.Listen(":8080")
}

func getProductHandler(c *fiber.Ctx) error {
	productId , err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	} 
	product , err := getProduct(productId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(product)
}

func getAllProductHandler(c *fiber.Ctx) error {
	products , err := getProducts()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(products)
}

func createProductHandler(c *fiber.Ctx) error {
	 p := new(Product)
	if err := c.BodyParser(p) ; 
	err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err := createProduct(p)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(p)
}

func updateProductHandler(c *fiber.Ctx) error {
	productId , err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	p := new(Product) 
	if err := c.BodyParser(p) ; err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	product , err := updateproduct(productId,p)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(product)
}


func deleteProductHandler(c *fiber.Ctx ) error  {
	 productId , err := strconv.Atoi(c.Params("Id"))
	 if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	 }
	 err = deleteProduct(productId)
	 if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	 }
	 return c.SendString("Delete Sucessesfully !! ")
}


