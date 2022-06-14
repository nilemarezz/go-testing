//go:build integration

package handlers_test

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"
	"testing/handlers"
	"testing/repositories"
	"testing/services"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscountIntegration(t *testing.T) {
	amount := 100
	expected := 80

	repo := repositories.NewPromotionRepositoryMock()
	repo.On("GetPromotion").Return(repositories.Promotion{
		ID:              1,
		PurchaseMin:     100,
		DiscountPercent: 20,
	}, nil)

	service := services.NewPromotionService(repo)
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

}
