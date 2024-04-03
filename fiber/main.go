package main

import (
	"encoding/json"
	"log"
	"fmt"

	"github.com/Marushio/golangconference/fiber/database"
	"github.com/Marushio/golangconference/fiber/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}


func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/api/v1/book", getUsers())
}

func main() {
	app := fiber.New()

	app.Use(recover.New())

	//hello world!
	app.Get("/hello/:name?", func(c *fiber.Ctx) error {
		return helloHandler(c)
	})

	// Example API
	app.Get("/api/posts", func(c *fiber.Ctx) error {
		posts := getPosts() // your logic
		if len(posts) == 0 {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   "There are no posts!",
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"posts":   posts,
		})
	})

	//hello world!
	app.Get("/hello/:name?", func(c *fiber.Ctx) error {
		return helloHandler(c)
	})

	// Example validate input
	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		if c.Is("json") {
			return c.Next()
		}
		return c.SendString("Only JSON allowed!")
	})

	app.Get("/validateJson", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello World",
		})
	})

	//Template example
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	//test with users db
	initDatabase()

	app.Get("/listUsersHandler", func(c *fiber.Ctx) error {
		return listUsersHandler(c)
	})

	log.Fatal(app.Listen(":3000"))

}

func helloHandler(c *fiber.Ctx) error {
	if c.Params("name") != "" {
		return c.SendString("Hello " + c.Params("name"))
		// => Hello john
	}
	return c.SendString("Where is john?")
}

func getPosts() []string {
	return []string{
		"Post 1",
		"Post 2",
		"Post 3",
	}
}



func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}


func getUsers(c *fiber.Ctx) {
	db := database.DBConn
	var users []User
	db.Find(&users)
	c.JSON(users)
}
