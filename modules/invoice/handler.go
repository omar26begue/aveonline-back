package invoice

import (
	"aveonline-back/models"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"time"
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

// GetInvoice godoc
// @Summary Invoice
// @Description Create the about information of the application.
// @Tags Invoice
// @Accept  json
// @Produce  json
// @Param   start      path   string     true  "Start"
// @Param   start      path   string     true  "End"
// @Success 200 {array} Invoice
// @Failure 400 {object} models.HTTPResponse
// @Router /invoice/{start}/{end} [get]
func (h Handler) GetInvoice(c *fiber.Ctx) error {
	start := c.Params("start")
	end := c.Params("end")

	format := "2006-01-02"
	dateStart, err := time.Parse(format, start)
	dateEnd, err := time.Parse(format, end)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusUnprocessableEntity).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code: fiber.StatusUnprocessableEntity, Message: err.Error(),
		})
	}

	dataInvoice, err := h.repository.GetPayment(dateStart, dateEnd)
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
// @Param about body InvoiceRequest true "Invoice"
// @Success 201 {array} Invoice
// @Failure 400 {object} models.HTTPResponse
// @Router /invoice [post]
func (h Handler) CreateInvoice(c *fiber.Ctx) error {
	invoiceRequest := new(InvoiceRequest)
	if err := c.BodyParser(&invoiceRequest); err != nil {
		return json.NewEncoder(c.Status(fiber.StatusUnprocessableEntity).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code: fiber.StatusUnprocessableEntity, Message: err.Error(),
		})
	}

	var validate = validator.New()
	err := validate.Struct(invoiceRequest)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code: fiber.StatusBadRequest, Message: err.Error(),
		})
	}

	dataInvoice, err := h.repository.Create(Invoice{
		FechaCreacion: time.Now(),
		PagoTotal:     invoiceRequest.PagoTotal,
		IdPromotion:   invoiceRequest.IdPromotion,
	})
	if err != nil {
		return err
	}

	return json.NewEncoder(c.Status(fiber.StatusOK).Type("json", "utf-8").Response().BodyWriter()).Encode(dataInvoice)
}
