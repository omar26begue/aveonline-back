package config

import (
	_ "aveonline-back/docs"
	"aveonline-back/modules/invoice"
	"aveonline-back/modules/medicine"
	"aveonline-back/modules/promotion"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Routers(app *fiber.App, repoInvoice *invoice.Repository, repoPromotion *promotion.Repository, repoMedicine *medicine.Repository) {
	app.Get("/docs/*", swagger.Handler)

	rInvoice := invoice.NewHandler(repoInvoice)
	rPromotion := promotion.NewHandler(repoPromotion)
	rMedicine := medicine.NewHandler(repoMedicine)

	apiV1 := app.Group("/v1", logger.New())
	{
		apiV1.Get("/promotion", rPromotion.GetAll)
		apiV1.Post("/promotion", rPromotion.Create)

		apiV1.Get("/invoice", rInvoice.ListInvoice)
		apiV1.Post("/invoice", rInvoice.CreateInvoice)
		apiV1.Get("/invoice/:start/:end", rInvoice.GetInvoice)

		apiV1.Get("/medicine", rMedicine.ListMedicine)
		apiV1.Post("/medicine", rMedicine.Create)
	}
}
