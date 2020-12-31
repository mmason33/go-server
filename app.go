package main

// reflex -r '\.go' -s -- sh -c "go run app.go"
import (
	"github.com/mmason33/go-project/config"
	"github.com/mmason33/go-project/database"
	"github.com/mmason33/go-project/handlers"

	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	dialect, dbStr := config.GetDBConfig()
	// Connected with database
	database.Connect(dialect, dbStr)
	// database.GetAll("users")
	database.Insert("users", map[string]interface{}{
		"name": "Guy",
		"age":  40,
	})

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// Bind handlers
	v1.Get("/users", handlers.UserList)
	v1.Post("/users", handlers.UserCreate)

	// Setup static files
	app.Static("/", "./static/public")

	// Handle not founds
	app.Use(handlers.NotFound)

	// Listen on port 3000
	log.Fatal(app.Listen(*port)) // go run app.go -port=:3000
}
