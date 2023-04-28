package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hilmimuharromi/go-fiber-mysql/configs"
	"github.com/hilmimuharromi/go-fiber-mysql/controllers"
)

func init() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	configs.ConnectDB(&config)
}

func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))
	app.Static("/", "./public")

	micro.Route("/files", func(router fiber.Router) {
		router.Post("/", controllers.UploadFileHandler)
	})

	micro.Route("/posts", func(router fiber.Router) {
		router.Post("/", controllers.CreatePostHandler)
		router.Get("", controllers.FindPosts)
	})
	micro.Route("/posts/:postId", func(router fiber.Router) {
		router.Delete("", controllers.DeletePost)
		router.Get("", controllers.FindPostById)
		router.Patch("", controllers.UpdatePost)
	})
	micro.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, MySQL, and GORM",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
