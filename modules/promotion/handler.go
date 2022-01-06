package promotion

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

// GetAll godoc
// @Summary Promotion
// @Description Create the about information of the application.
// @Tags Promotion
// @Accept  json
// @Produce  json
// @Success 200 {array} Promocion
// @Failure 400 {object} models.HTTPResponse
// @Router /promotion [get]
func (h Handler) GetAll(c *fiber.Ctx) error {
	dataPromotion, err := h.repository.GetAll()
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    400,
			Message: err.Error(),
		})
	}

	return json.NewEncoder(c.Status(fiber.StatusOK).Type("json", "utf-8").Response().BodyWriter()).Encode(dataPromotion)
}

// Create godoc
// @Summary Promotion
// @Description Create the about information of the application.
// @Tags Promotion
// @Accept  json
// @Produce  json
// @Param promotion body PromocionRequest true "Promocion"
// @Success 201 {object} Promocion
// @Failure 400 {object} models.HTTPResponse
// @Router /promotion [post]
func (h Handler) Create(c *fiber.Ctx) error {
	promotionRequest := new(PromocionRequest)
	if err := c.BodyParser(&promotionRequest); err != nil {
		return json.NewEncoder(c.Status(fiber.StatusUnprocessableEntity).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code: fiber.StatusUnprocessableEntity, Message: err.Error(),
		})
	}

	var validate = validator.New()
	err := validate.Struct(promotionRequest)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code: fiber.StatusBadRequest, Message: err.Error(),
		})
	}

	parse := "2006-01-02"
	inicio, err := time.Parse(parse, promotionRequest.FechaInicio)
	fin, err := time.Parse(parse, promotionRequest.FechaFin)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusUnprocessableEntity).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code: fiber.StatusUnprocessableEntity, Message: err.Error(),
		})
	}

	resultPromotion, err := h.repository.SearchDate(inicio, fin)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusUnprocessableEntity).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code: fiber.StatusUnprocessableEntity, Message: err.Error(),
		})
	}
	if len(*resultPromotion) > 0 {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code: fiber.StatusBadRequest, Message: "Ya existen promociones para esta rango de fechas",
		})
	}

	dataPromotion, err := h.repository.Create(Promocion{
		Descripcion: promotionRequest.Descripcion,
		Porcentaje:  promotionRequest.Porcentaje,
		FechaInicio: inicio,
		FechaFin:    fin,
	})
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    400,
			Message: err.Error(),
		})
	}

	return json.NewEncoder(c.Status(fiber.StatusOK).Type("json", "utf-8").Response().BodyWriter()).Encode(dataPromotion)
}


