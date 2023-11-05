package model

type (
	SearchProductQueryRequest struct {
		OperationName string              `json:"operationName"`
		Query         string              `json:"query"`
		Variables     SearchProductParams `json:"variables"`
	}

	SearchProductParams struct {
		Params string `json:"params"`
	}

	SearchProductQueryResponse struct {
		Data struct {
			CategoryProducts struct {
				Count int64 `json:"count"`
				Data  []struct {
					Id       string `json:"id"`
					Name     string `json:"name"`
					Url      string `json:"url"`
					ImageUrl string `json:"imageUrl"`
					PriceInt int64  `json:"priceInt"`
					Rating   int64  `json:"rating"`
					Shop     struct {
						Name string `json:"name"`
					} `json:"shop"`
				}
			} `json:"CategoryProducts"`
		} `json:"data"`
	}
)
