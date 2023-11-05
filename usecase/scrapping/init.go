package scrapping

import (
	"sync"

	"github.com/doniiawan/tokopedia-scrapper/config"
)

var (
	s sync.Mutex
	u *ScrappingUsecase
)

func New(
	tokopediaAPIDomain tokopediaAPIDomain,
	config config.Config,
) (*ScrappingUsecase, error) {
	s.Lock()
	defer s.Unlock()

	if u == nil {
		u = &ScrappingUsecase{
			tokopediaAPIDomain: tokopediaAPIDomain,
			config:             config,
		}
	}

	return u, nil
}
