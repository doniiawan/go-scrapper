package handler

import "github.com/doniiawan/tokopedia-scrapper/config"

type Handler struct {
	config config.Config
}

func New(config config.Config) *Handler {
	return &Handler{
		config: config,
	}
}
