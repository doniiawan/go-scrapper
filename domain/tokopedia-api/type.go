package tokopedia_api

import (
	"net/http"

	"github.com/doniiawan/tokopedia-scrapper/config"
	"github.com/doniiawan/tokopedia-scrapper/model"
)

type TokopediaAPIDomain struct {
	TokopediaAPIResource tokopediaAPIResource
	config               config.Config
}

type tokopediaAPIResource interface {
	GetMobileProducts(body *model.SearchProductQueryRequest) (*model.SearchProductQueryResponse, int, error)
	PostAPI(method, url string, body []byte) (res *http.Response, statusCode int, err error)
	GetAPI(method, url string) (res *http.Response, statusCode int, err error)
}
