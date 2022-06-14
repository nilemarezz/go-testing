package handlers_test

import (
	"errors"
	"fmt"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"
	"testing/handlers"
	"testing/services"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		amount := 100
		expected := 80

		service := services.NewPromotionServiceMock()
		service.On("CalculateDiscount", amount).Return(expected, nil)

		handler := handlers.NewPromotionHandler(service)

		app := fiber.New()
		app.Get("/calculate", handler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		res, _ := app.Test(req)
		defer res.Body.Close()

		fmt.Println(res.StatusCode)

		assert.Equal(t, fiber.StatusOK, res.StatusCode)

		body, _ := io.ReadAll(res.Body)
		assert.Equal(t, strconv.Itoa(expected), string(body))
		fmt.Println("---------", string(body))
	})

	t.Run("fail amount string", func(t *testing.T) {
		amount := "test"
		expected := 80

		service := services.NewPromotionServiceMock()
		service.On("CalculateDiscount", amount).Return(expected, nil)

		handler := handlers.NewPromotionHandler(service)

		app := fiber.New()
		app.Get("/calculate", handler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		res, _ := app.Test(req)
		defer res.Body.Close()

		fmt.Println(res.StatusCode)

		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
		service.AssertNotCalled(t, "CalculateDiscount")
	})

	t.Run("fail service", func(t *testing.T) {
		amount := 100

		service := services.NewPromotionServiceMock()
		service.On("CalculateDiscount", amount).Return(0, errors.New(""))

		handler := handlers.NewPromotionHandler(service)

		app := fiber.New()
		app.Get("/calculate", handler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		res, _ := app.Test(req)
		defer res.Body.Close()

		fmt.Println(res.StatusCode)

		assert.Equal(t, fiber.StatusNotFound, res.StatusCode)
		service.AssertNotCalled(t, "CalculateDiscount")
	})
}
