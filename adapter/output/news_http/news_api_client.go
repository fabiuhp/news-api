package news_http

import (
	"news-api/application/domain"
	resterr "news-api/configuration/rest_err"
)

type newsClient struct {
}

func NewNewsClient() *newsClient {
	return &newsClient{}
}

func (nc *newsClient) GetNewsPort(newsDomain domain.NewsReqDomain) (*domain.NewsDomain, *resterr.RestErr) {
	return nil, nil
}
