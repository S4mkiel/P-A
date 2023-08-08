package controller

import (
	"fmt"

	"github.com/S4mkiel/p-a/domain/entity"
	"github.com/S4mkiel/p-a/domain/service"
	"github.com/S4mkiel/p-a/infra/http/dto"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type PaController struct {
	logger    *zap.SugaredLogger
	paService *service.PaService
}

func NewPaController(logger *zap.SugaredLogger, paService *service.PaService) *PaController {
	return &PaController{logger: logger, paService: paService}
}

func (c PaController) RegisterRoutes(app fiber.Router) {
	pa := app.Group("/v1")
	pa.Post("/create", c.Create)
	pa.Post("/update", c.Update)
	pa.Post("/get", c.Get)
	pa.Get("/get-all", c.GetAll)
	pa.Delete("/delete", c.Delete)
}

func (c PaController) Create(ctx *fiber.Ctx) error {
	response := dto.Success{
		Success: true,
		Message: "",
	}

	failureResponse := dto.Failed{
		Success: false,
		Error:   "",
	}

	type Payload struct {
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var request Payload
	if err := ctx.BodyParser(&request); err != nil {
		failureResponse.Error = "Invalid request"
		return ctx.Status(fiber.StatusBadRequest).JSON(failureResponse)
	}

	u := entity.User{
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}

	error := fmt.Sprintf("Duplicate entry: '%s", u.Email)

	user, err := c.paService.Create(ctx.Context(), u)
	if err != nil {
		failureResponse.Error = error
		return ctx.Status(fiber.StatusBadRequest).JSON(failureResponse)
	}

	response.Data = user

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c PaController) Update(ctx *fiber.Ctx) error {
	response := dto.Success{
		Success: true,
		Message: "",
	}

	failureResponse := dto.Failed{
		Success: false,
		Error:   "",
	}

	type Payload struct {
		Id        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var request Payload

	if err := ctx.BodyParser(&request); err != nil {
		failureResponse.Error = "Invalid request"
		return ctx.Status(fiber.StatusBadRequest).JSON(failureResponse)
	}

	u := entity.User{
		ID:        request.Id,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}

	user, err := c.paService.Update(ctx.Context(), u)
	if err != nil {
		failureResponse.Error = "Invalid request"
		return ctx.Status(fiber.StatusBadRequest).JSON(failureResponse)
	}

	response.Data = user

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c PaController) Get(ctx *fiber.Ctx) error {
	response := dto.Success{
		Success: true,
		Message: "",
	}

	failureResponse := dto.Failed{
		Success: false,
		Error:   "",
	}

	type Payload struct {
		Id int `json:"id"`
	}

	var request Payload

	if err := ctx.BodyParser(&request); err != nil {
		failureResponse.Error = "Invalid request"
		return ctx.Status(fiber.StatusBadRequest).JSON(failureResponse)
	}

	user, err := c.paService.Get(ctx.Context(), request.Id)
	if err != nil {
		failureResponse.Error = "Invalid request"
		return ctx.Status(fiber.StatusBadRequest).JSON(failureResponse)
	}

	response.Data = user

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c PaController) GetAll(ctx *fiber.Ctx) error {
	response := dto.Success{
		Success: true,
		Message: "",
	}

	failureResponse := dto.Failed{
		Success: false,
		Error:   "",
	}

	users, err := c.paService.GetAll(ctx.Context())
	if err != nil {
		failureResponse.Error = "Invalid request"
		return ctx.Status(fiber.StatusBadRequest).JSON(failureResponse)
	}

	response.Data = users

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c PaController) Delete(ctx *fiber.Ctx) error {
	response := dto.Success{
		Success: true,
		Message: "",
	}

	failureResponse := dto.Failed{
		Success: false,
		Error:   "",
	}

	type Payload struct {
		Email string `json:"email"`
	}

	var request Payload

	if err := ctx.BodyParser(&request); err != nil {
		failureResponse.Error = "Invalid request"
		return ctx.Status(fiber.StatusBadRequest).JSON(failureResponse)
	}

	err := c.paService.Delete(ctx.Context(), request.Email)
	if err != nil {
		failureResponse.Error = "Invalid request"
		return ctx.Status(fiber.StatusBadRequest).JSON(failureResponse)
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
