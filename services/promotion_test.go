package services_test

import (
	"errors"
	"testing"
	errHandle "testing/errors"
	"testing/repositories"
	"testing/services"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDiscount(t *testing.T) {
	type TestCase struct {
		name            string
		purchaseMin     int
		discountPercent int
		amount          int
		expected        interface{}
	}

	cases := []TestCase{
		{name: "applied 100", purchaseMin: 100, discountPercent: 20, amount: 100, expected: 80},
		{name: "not apply 100", purchaseMin: 100, discountPercent: 20, amount: 30, expected: 30},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := repositories.NewPromotionRepositoryMock()
			repo.On("GetPromotion").Return(repositories.Promotion{
				ID:              1,
				PurchaseMin:     c.purchaseMin,
				DiscountPercent: c.discountPercent,
			}, nil)

			service := services.NewPromotionService(repo)

			discount, _ := service.CalculateDiscount(c.amount)
			assert.Equal(t, c.expected, discount)
		})
	}
}

func TestCalculateDiscountZero(t *testing.T) {
	repo := repositories.NewPromotionRepositoryMock()
	repo.On("GetPromotion").Return(repositories.Promotion{
		ID:              1,
		PurchaseMin:     100,
		DiscountPercent: 20,
	}, nil)

	service := services.NewPromotionService(repo)

	_, err := service.CalculateDiscount(-10)
	assert.ErrorIs(t, err, errHandle.ErrZeroAmount)
	repo.AssertNotCalled(t, "GetPromotion")
}

func TestCalculateDiscountRepoError(t *testing.T) {
	repo := repositories.NewPromotionRepositoryMock()
	repo.On("GetPromotion").Return(repositories.Promotion{}, errors.New(""))

	service := services.NewPromotionService(repo)

	_, err := service.CalculateDiscount(20)
	assert.ErrorIs(t, err, errHandle.ErrRepository)
}
