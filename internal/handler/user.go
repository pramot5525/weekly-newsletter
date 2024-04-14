package handler

import (
	"fmt"
	"weekly-newsletter/internal/model"
	"weekly-newsletter/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sourcegraph/conc"
)

type Handler struct {
	service service.UserService
}

func New(service service.UserService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Subscribe(c *fiber.Ctx) error {
	req := model.UserRequest{}
	err := c.BodyParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	validate := validator.New()

	err = validate.Struct(req)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.Error()})

	}

	user, err := h.service.CreateUser(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": user})
}

func (h *Handler) Unsubscribe(c *fiber.Ctx) error {
	req := model.UserRequest{}
	err := c.QueryParser(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	validate := validator.New()

	err = validate.Struct(req)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errors.Error()})

	}

	err = h.service.DeleteUser(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}

func (h *Handler) Publish(c *fiber.Ctx) error {
	users, err := h.service.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var wg conc.WaitGroup
	for _, user := range users {
		wg.Go(func() {
			fmt.Println("sending email to", user.Email)
			err := h.service.SendEmail(user.Email)
			if err != nil {
				log.Errorf("error sending email to %s: %v", user.Email, err)
			}

		})
	}
	wg.Wait()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}
