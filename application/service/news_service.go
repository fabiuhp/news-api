package service

import (
	"fmt"
	"news-api/application/domain"
	"news-api/application/port/output"
	"news-api/configuration/logger"
	resterr "news-api/configuration/rest_err"
)

type newsService struct {
	newsPort output.GetNewsPort
}

func NewNewsService(newsPort output.GetNewsPort) *newsService {
	return &newsService{
		newsPort: newsPort,
	}
}

func (ns *newsService) GetNews(newsReq domain.NewsReqDomain) (*domain.NewsDomain, *resterr.RestErr) {
	logger.Info(
		fmt.Sprintf("Init getNewsService function, query=%s, from=%s",
			newsReq.Query, newsReq.From))

	return ns.newsPort.GetNewsPort(newsReq)
}
