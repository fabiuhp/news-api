package input

import (
	"news-api/application/domain"
	resterr "news-api/configuration/rest_err"
)

type NewsUseCase interface {
	GetNews(domain.NewsReqDomain) (*domain.NewsDomain, *resterr.RestErr)
}
