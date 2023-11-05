package worker

import (
	"context"

	"github.com/doniiawan/tokopedia-scrapper/config"
)

type AppHandler struct {
	cfg     config.Config
	usecase scrappingUsecase
}

type scrappingUsecase interface {
	ProcessData(ctx context.Context, config config.Config) error
}
