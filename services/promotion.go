package services

import (
	"testing/errors"
	"testing/repositories"
)

type PromotionService interface {
	CalculateDiscount(amount int) (int, error)
}

type promotionService struct {
	repo repositories.PromotionRepository
}

func (s promotionService) CalculateDiscount(amount int) (int, error) {
	if amount <= 0 {
		return 0, errors.ErrZeroAmount
	}

	promotion, err := s.repo.GetPromotion()
	if err != nil {
		return 0, errors.ErrRepository
	}

	if amount >= promotion.PurchaseMin {
		return amount - (promotion.DiscountPercent * amount / 100), nil
	}
	return amount, nil
}

func NewPromotionService(repo repositories.PromotionRepository) PromotionService {
	return promotionService{repo: repo}
}
