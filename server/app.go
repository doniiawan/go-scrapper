package server

import (
	"context"

	"github.com/doniiawan/tokopedia-scrapper/config"
	tokopediaDomain "github.com/doniiawan/tokopedia-scrapper/domain/tokopedia-api"
	"github.com/doniiawan/tokopedia-scrapper/handler/worker"
	tokopediaResource "github.com/doniiawan/tokopedia-scrapper/resource/tokopedia-api"
	"github.com/doniiawan/tokopedia-scrapper/usecase/scrapping"
	"github.com/sirupsen/logrus"
)

type AppHandlers struct {
	AppHandlers *worker.AppHandler
}

func InitApp(ctx context.Context) (cfg config.Config, handler *AppHandlers, err error) {
	cfg, err = config.InitConfig()
	if err != nil {
		logrus.Error(err)
		return
	}

	tokpedResource, err := tokopediaResource.New(cfg)
	if err != nil {
		logrus.Error(err)
		return
	}

	tokpedDomain, err := tokopediaDomain.New(cfg, tokpedResource)
	if err != nil {
		logrus.Error(err)
		return
	}

	scrappingUsecase, err := scrapping.New(tokpedDomain, cfg)
	if err != nil {
		logrus.Error(err)
		return
	}

	scrappingHandler, err := worker.New(cfg, scrappingUsecase)
	if err != nil {
		logrus.Error(err)
		return
	}

	handler = &AppHandlers{
		AppHandlers: scrappingHandler,
	}

	return
}
