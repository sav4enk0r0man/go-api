package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sav4enk0r0man/go-api/database"
	"github.com/sav4enk0r0man/go-api/task"
	"log"
)

const (
	apiVersion = 1
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.ConnectDB()
	defer database.DB.Close()

	apiGroup := fmt.Sprintf("/api/v%d", apiVersion)
	api := app.Group(apiGroup)
	task.Register(api, database.DB)

	log.Fatal(app.Listen(":5000"))
}
