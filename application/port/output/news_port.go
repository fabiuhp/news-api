package output

import (
	"news-api/application/domain"
	resterr "news-api/configuration/rest_err"
)

type GetNewsPort interface {
	GetNewsPort(domain.NewsReqDomain) (*domain.NewsDomain, *resterr.RestErr)
}
