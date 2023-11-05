package tokopedia_api

import (
	"net/http"
	"sync"
	"time"

	"github.com/doniiawan/tokopedia-scrapper/config"
)

var (
	s sync.Mutex
	r *TokopediaAPIResource
)

func New(config config.Config) (*TokopediaAPIResource, error) {
	s.Lock()
	defer s.Unlock()

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	if r == nil {
		r = &TokopediaAPIResource{
			config: config,
			client: client,
		}
	}

	return r, nil
}
