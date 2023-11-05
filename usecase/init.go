package usecase

import (
	"github.com/doniiawan/tokopedia-scrapper/domain"
)

type Usecase struct {
	domain *domain.Domain
}

func New(domain *domain.Domain) *Usecase {
	return &Usecase{
		domain: domain,
	}
}
