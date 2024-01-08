package interfaces

import (
	"kicks/pkg/domain"
	"kicks/pkg/utils/models"
)

type OfferUseCase interface {
	AddNewOffer(model models.OfferMaking) error
	MakeOfferExpire(id int) error
	GetOffers() ([]domain.Offer, error)
}
