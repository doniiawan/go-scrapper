package main

import (
	"context"

	"github.com/doniiawan/tokopedia-scrapper/server"
	"github.com/sirupsen/logrus"
)

func runApp() {
	ctx := context.TODO()
	_, runner, err := server.InitApp(ctx)
	if err != nil {
		logrus.Error(err)
		return
	}

	runner.AppHandlers.Run(ctx)

}
