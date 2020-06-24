package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
	"github.com/markbates/pkger"
)

var prefix = flag.String("prefix", "/g", "String to prefix every path with. Can be useful when hosting multiple applications on the same domain")
var domain = flag.String("domain", "localhost", "Vanity domain")
var user = flag.String("user", "", "Your GitHub username")

func main() {
	flag.Parse()

	if len(*user) < 1 {
		log.Fatal("Please set -user flag")
	}

	// engine := html.NewFileSystem(pkger.Dir("views/"), ".html")
	engine := html.NewFileSystem(pkger.Dir("/views"), ".html")

	app := fiber.New(&fiber.Settings{
		Views:                 engine,
		DisableStartupMessage: true,
	})

	app.Get(*prefix+"/:package", func(c *fiber.Ctx) {
		c.Render("vanity", fiber.Map{
			"path": *domain + "/" + c.Params("package"),
			"repo": "github.com/" + *user + c.Params("package"),
		})
	})

	app.Listen(3000)
}
