package invoice

import (
	"aveonline-back/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	repository *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repository: repo}
}

// ListInvoice godoc
// @Summary Invoice
// @Description Create the about information of the application.
// @Tags Invoice
// @Accept  json
// @Produce  json
// @Success 200 {array} Invoice
// @Failure 400 {object} models.HTTPResponse
// @Router /invoice [get]
func (h Handler) ListInvoice(c *fiber.Ctx) error {
	dataInvoice, err := h.repository.GetAll()
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    400,
			Message: err.Error(),
		})
	}

	return json.NewEncoder(c.Status(fiber.StatusOK).Type("json", "utf-8").Response().BodyWriter()).Encode(dataInvoice)
}

// CreateInvoice godoc
// @Summary Create
// @Description Create the about information of the application.
// @Tags Invoice
// @Accept  json
// @Produce  json
// @Param about body Invoice true "Invoice"
// @Success 201 {array} Invoice
// @Failure 400 {object} models.HTTPResponse
// @Router /invoice [post]
func (h Handler) CreateInvoice(c *fiber.Ctx) error {
	return nil
}
