package medicine

import (
	"aveonline-back/models"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	repository *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repository: repo}
}

// ListMedicine godoc
// @Summary Medicine
// @Description Create the about information of the application.
// @Tags Medicine
// @Accept  json
// @Produce  json
// @Success 200 {array} Medicine
// @Failure 400 {object} models.HTTPResponse
// @Router /medicine [get]
func (h Handler) ListMedicine(c *fiber.Ctx) error {
	dataMedicine, err := h.repository.GetAll()
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    400,
			Message: err.Error(),
		})
	}

	return json.NewEncoder(c.Status(fiber.StatusOK).Type("json", "utf-8").Response().BodyWriter()).Encode(dataMedicine)
}

// Create godoc
// @Summary medicine
// @Description Create the about information of the application.
// @Tags Medicine
// @Accept  json
// @Produce  json
// @Param medicine body MedicineRequest true "Medicine"
// @Success 201 {object} Medicine
// @Failure 400 {object} models.HTTPResponse
// @Router /medicine [post]
func (h Handler) Create(c *fiber.Ctx) error {
	medicineRequest := new(MedicineRequest)
	if err := c.BodyParser(&medicineRequest); err != nil {
		return json.NewEncoder(c.Status(fiber.StatusUnprocessableEntity).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code: fiber.StatusUnprocessableEntity, Message: err.Error(),
		})
	}

	var validate = validator.New()
	err := validate.Struct(medicineRequest)
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code: fiber.StatusBadRequest, Message: err.Error(),
		})
	}

	dataMedicine, err := h.repository.Create(Medicine{
		Name:     medicineRequest.Name,
		Price:    medicineRequest.Price,
		Location: medicineRequest.Location,
	})
	if err != nil {
		return json.NewEncoder(c.Status(fiber.StatusBadRequest).Type("json", "utf-8").Response().BodyWriter()).Encode(models.HTTPResponse{
			Code:    400,
			Message: err.Error(),
		})
	}

	return json.NewEncoder(c.Status(fiber.StatusOK).Type("json", "utf-8").Response().BodyWriter()).Encode(dataMedicine)
}
