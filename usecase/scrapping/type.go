package scrapping

import (
	"context"

	"github.com/doniiawan/tokopedia-scrapper/config"
	"github.com/doniiawan/tokopedia-scrapper/model"
)

type ScrappingUsecase struct {
	tokopediaAPIDomain tokopediaAPIDomain
	config             config.Config
}

type tokopediaAPIDomain interface {
	GetMobileProductsList(ctx context.Context, page string) (res *model.SearchProductQueryResponse, statusCode int, err error)
}
