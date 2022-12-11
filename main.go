package main

import (
	"golang-fiber-gorm/controllers"
	"golang-fiber-gorm/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	models.ConnectDB()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept"}))

	api := app.Group("/api")
	book := api.Group("/books")

	book.Get("/", controllers.Index)
	book.Get("/:id", controllers.Show)
	book.Post("/", controllers.Create)
	book.Put("/:id", controllers.Update)
	book.Delete("/:id", controllers.Delete)

	app.Listen(":8000")
}
