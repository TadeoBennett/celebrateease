package main

import (
	"fmt"
	"os"
	"tadeobennett/celebrateease/initializers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDb()
}

func main() {
	//load templates
	// Initialize standard Go html template engine
	//tells where to find all tmpl files
	engine := html.New("./views", ".tmpl")

	//setup app views engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//configure app files(css, js, images, etc...)
	app.Static("/", "./public/assets/")

	//routes
	Routes(app)

	//start app on port 3000
	port := ":" +  os.Getenv("PORT")
	app.Listen(port)
	fmt.Println("Hello, World!")
}
