package tokopedia_api

import (
	"sync"

	"github.com/doniiawan/tokopedia-scrapper/config"
)

var (
	s sync.Mutex
	d *TokopediaAPIDomain
)

func New(cfg config.Config, tokopediaRes tokopediaAPIResource) (*TokopediaAPIDomain, error) {
	s.Lock()
	defer s.Unlock()

	if d == nil {
		d = &TokopediaAPIDomain{
			config:               cfg,
			TokopediaAPIResource: tokopediaRes,
		}
	}
	return d, nil
}
