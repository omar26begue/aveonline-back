package main

import (
	"aveonline-back/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

// @title Back Aveonline
// @version 1.0.0
// @description Back Aveonline
// @termsOfService http://swagger.io/terms/

// @contact.name Omar Isalgué Begue
// @contact.email omar26begue@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath /v1
func main() {
	// leyendo las variables de entorno
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {

	}

	repoInvoice, repoPromotion, repoMedicine := config.InitDatabase()

	app := fiber.New()
	app.Use(cors.New())
	config.Routers(app, repoInvoice, repoPromotion, repoMedicine)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":" + viper.GetString("PORT"))
}
