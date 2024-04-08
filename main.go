package main

import (
	"log"
	"os"
	"time"

	_ "github.com/boon/fiber-test/docs"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)
  

type Book struct {
	ID		int `json:"id"`
	Title	string `json:"title"`
	Author  string `json:"author"`
}

var books []Book

func checkMiddleware(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["role"] != "admin" {
		return fiber.ErrUnauthorized
	}
	return c.Next()
}


  // @title Book API
  // @description This is a sample server for a book API.
  // @version 1.0
  // @host localhost:8080
  // @BasePath /
  // @schemes http
  // @securityDefinitions.apikey ApiKeyAuth
  // @in header
  // @name Authorization


func main() {
	if err := godotenv.Load() ; err != nil {
		log.Fatal("load .env error ")
	}
	
	engine := html.New("./views/" ,".html")
	
	app := fiber.New(fiber.Config{
		Views: engine,
	  })

	app.Get("/swagger/*", swagger.HandlerDefault) // default


	books = append(books, Book{ID :1, Title: "Nebula" , Author: "Boon"})
	books = append(books, Book{ID :2, Title: "Mm" , Author: "Boon"})

	app.Post("/login" ,login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	app.Use(checkMiddleware)


	app.Get("/books", getBooks)
	app.Get("/books/:id",getBook)
	app.Post("/books",createBook)
	app.Put("/books/:id",updateBook)
	app.Delete("books/:id",deleteBook)

	app.Post("/upload",uploadFile)
	app.Get("/test-html",testhtml)

	app.Get("/config",getEnv)

	app.Listen(":8080")
}


func uploadFile(c *fiber.Ctx) error {
	file ,err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err = c.SaveFile(file,"./uploads/" + file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("File Upload Complete! ")
}

func testhtml(c * fiber.Ctx) error {
	return c.Render("index",fiber.Map{
		"Title":"Hello, World!",
		"Name":"Boon",
	})
}

func getEnv(c *fiber.Ctx) error {
	secret := os.Getenv("SECRET")

	if secret == "" {
		secret = "defaultsecret"
	}
	return c.JSON(fiber.Map{

		"SECRET":secret,
	})
}

type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

var member =  User {
	Email: "user@ex.com",
	Password: "12345",
}

func login(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.Email != member.Email || user.Password != member.Password {
		return fiber.ErrUnauthorized
	}

	// Create token
    token := jwt.New(jwt.SigningMethodHS256)

    // Set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["email"] = user.Email
    claims["role"] = "admin"
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

    // Generate encoded token
    t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
      return c.SendStatus(fiber.StatusInternalServerError)
    }
	return c.JSON(fiber.Map {
		"message":"Login Sucess",
		"tokens":t, 
	})
}

