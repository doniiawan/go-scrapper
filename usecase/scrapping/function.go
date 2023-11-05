package scrapping

import (
	"context"
	"fmt"

	"github.com/doniiawan/tokopedia-scrapper/config"
	"github.com/sirupsen/logrus"
)

func (u *ScrappingUsecase) ProcessData(ctx context.Context, cfg config.Config) error {
	// var (
	// 	wg sync.WaitGroup
	// )

	// for i := 1; i <= 10; i++ {
	// 	wg.Add(1)
	// 	go func(page int) {
	// 		defer wg.Done()

	// 	}(i)
	// }

	res, statusCode, err := u.tokopediaAPIDomain.GetMobileProductsList(ctx, fmt.Sprintf("%d", 1))
	if err != nil {
		logrus.Error(err)
		return err
	}
	logrus.Info(statusCode, " ", res)

	return nil
}
