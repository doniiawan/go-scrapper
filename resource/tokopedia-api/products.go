package tokopedia_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/doniiawan/tokopedia-scrapper/model"
	"github.com/sirupsen/logrus"
)

func (r *TokopediaAPIResource) GetMobileProducts(body *model.SearchProductQueryRequest) (*model.SearchProductQueryResponse, int, error) {
	logrus.Info("GetMobileProducts Resource")
	logrus.Info(r.config.Env.Urls.TokopediaUrl)
	url := fmt.Sprintf("%v/graphql/SearchProductQuery", r.config.Env.Urls.TokopediaUrl)
	reqPayload, err := json.Marshal(body)
	logrus.Info(string(reqPayload))
	if err != nil {
		return nil, 0, err
	}

	response, _, err := r.PostAPI(http.MethodPost, url, reqPayload)
	if err != nil {
		logrus.Error("[debug] error :", err)
		return nil, 0, err
	}

	payload, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}

	mappedResponse := new(model.SearchProductQueryResponse)
	err = json.Unmarshal(payload, mappedResponse)
	if err != nil {
		return nil, 0, err
	}

	return mappedResponse, response.StatusCode, nil
}

func (r *TokopediaAPIResource) PostAPI(method, url string, body []byte) (res *http.Response, statusCode int, err error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")

	res, err = r.client.Do(req)
	if err != nil {
		logrus.Error("[debug] error :", err)
		return nil, res.StatusCode, err
	}

	return res, res.StatusCode, nil
}

func (r *TokopediaAPIResource) GetAPI(method, url string) (res *http.Response, statusCode int, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, 0, err
	}

	res, err = r.client.Do(req)
	if err != nil {
		logrus.Error("[debug] error :", err)
		return nil, res.StatusCode, err
	}

	return res, res.StatusCode, nil
}
