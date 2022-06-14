package handlers

import (
	"strconv"
	"testing/services"

	"github.com/gofiber/fiber/v2"
)

type PromotionHandler interface {
	CalculateDiscount(c *fiber.Ctx) error
}

type promotionHandler struct {
	service services.PromotionService
}

func NewPromotionHandler(service services.PromotionService) promotionHandler {
	return promotionHandler{service}
}

func (h promotionHandler) CalculateDiscount(c *fiber.Ctx) error {
	// http://localhost:8000/calculate?amount=100
	amountStr := c.Query("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	discount, err := h.service.CalculateDiscount(amount)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.SendString(strconv.Itoa(discount))
}
