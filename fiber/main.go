package main

import (
	"encoding/json"
	"log"

	"github.com/Marushio/golangconference/fiber/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type user struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type (
	User struct {
		Name string `validate:"required,min=5,max=20"` // Required field, min 5 char long max 20
		Age  int    `validate:"required,teener"`       // Required field, and client needs to implement our 'teener' tag format which we'll see later
	}

	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

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

func listUsersHandler(c *fiber.Ctx) {
	//Cria a conecxao com o banco
	//db, err := sql.Open("sqlite3", "users.db")
	database.DBConn, err = gorm.Open("sqlite3", "books.db")

	//Tratamento de erro caso algum problema com o banco de dados
	if err != nil {
		panic("Error on databse connection")
	}

	//Fecha a conexao com o banco depois que todos os codigos forem executados
	defer db.Close()

	//Executa o select na base
	rows, err := db.Query("SELECT * FROM users")
	//Tratamento de erro caso algum problema no select
	if err != nil {
		panic("Error on databse select")
	}
	defer rows.Close()

	//Declaracao de um slice de users (como se fosse um vetor dinamico)
	users := []*user{}

	//Iteracao em cima do retorno do select
	for rows.Next() {
		//Cria variavel user seguindo a struct
		var u user

		//Scan da linha para pegar um user com tratamento caso algo de errado
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			panic("Error on row scan")
		}

		//Adiciona o usuario da linha ao slice de users
		users = append(users, &u)
	}
	// Marshal the users slice to JSON
	usersJSON, err := json.Marshal(users)
	if err != nil {
		// Handle error
	}

	// Return the JSON response
	return c.JSON(usersJSON)
}
