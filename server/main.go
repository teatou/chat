package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/teatou/chat/server/initializers"
	"github.com/teatou/chat/server/storage"
)

func init() {
	initializers.LoadEnv()
	storage.ConnectToDb()
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://gofiber.io, https://gofiber.net",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))

	fmt.Println(storage.UserCollection)
}
