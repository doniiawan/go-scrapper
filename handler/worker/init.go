package worker

import "github.com/doniiawan/tokopedia-scrapper/config"

func New(cfg config.Config, usecase scrappingUsecase) (*AppHandler, error) {
	return &AppHandler{
		cfg:     cfg,
		usecase: usecase,
	}, nil
}
