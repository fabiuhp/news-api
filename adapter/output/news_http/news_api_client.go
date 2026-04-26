package news_http

import (
	"news-api/adapter/output/model/response"
	"news-api/application/domain"
	"news-api/configuration/env"
	resterr "news-api/configuration/rest_err"

	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
)

var (
	client *resty.Client
)

type newsClient struct {
}

func NewNewsClient() *newsClient {
	client = resty.New().SetBaseURL("https://newsapi.org/v2")
	return &newsClient{}
}

func (nc *newsClient) GetNewsPort(newsDomain domain.NewsReqDomain) (*domain.NewsDomain, *resterr.RestErr) {
	newsResponse := &response.NewsClientResponse{}

	_, err := client.R().
		SetQueryParams(map[string]string{
			"q":      newsDomain.Query,
			"from":   newsDomain.From,
			"apiKey": env.GetNewsTokenAPI(),
		}).SetResult(newsResponse).
		Get("/everything")

	if err != nil {
		return nil, resterr.NewInternalServerError("error when trying to get news")
	}

	newsResponseDomain := &domain.NewsDomain{}
	copier.Copy(newsResponseDomain, newsResponse)

	return newsResponseDomain, nil
}
