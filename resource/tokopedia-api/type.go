package tokopedia_api

import (
	"net/http"

	"github.com/doniiawan/tokopedia-scrapper/config"
)

type (
	TokopediaAPIResource struct {
		config config.Config
		client *http.Client
	}
)
