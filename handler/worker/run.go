package worker

import (
	"context"

	"github.com/doniiawan/tokopedia-scrapper/config"
	"github.com/sirupsen/logrus"
)

func (h *AppHandler) Run(ctx context.Context) {
	config, error := config.InitConfig()
	if error != nil {
		logrus.Error(error)
		return
	}

	result := h.usecase.ProcessData(ctx, config)
	if result != nil {
		logrus.Error(result)
		return
	}
	logrus.Info("success")
}
