package news_http

import (
	"fmt"
	"news-api/adapter/output/model/response"
	"news-api/application/domain"
	"news-api/configuration/env"
	"news-api/configuration/logger"
	resterr "news-api/configuration/rest_err"

	"github.com/go-resty/resty/v2"
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

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"q":      newsDomain.Query,
			"from":   newsDomain.From,
			"apiKey": env.GetNewsTokenAPI(),
		}).SetResult(newsResponse).
		Get("/everything")

	if err != nil {
		logger.Info(fmt.Sprintf("Error in HTTP request: %v", err))
		return nil, resterr.NewInternalServerError("error when trying to get news")
	}

	logger.Info(fmt.Sprintf("API Response Status: %d, Body length: %d", resp.StatusCode(), len(resp.String())))

	if resp.StatusCode() != 200 {
		logger.Info(fmt.Sprintf("API Error Response: %s", resp.String()))
		return nil, resterr.NewInternalServerError("error from news api: " + resp.String())
	}

	logger.Info(fmt.Sprintf("Parsed Response: Status=%s, TotalResults=%d, Articles=%d", newsResponse.Status, newsResponse.TotalResults, len(newsResponse.Articles)))

	// Manually convert response to domain
	newsResponseDomain := &domain.NewsDomain{
		Status:       newsResponse.Status,
		TotalResults: newsResponse.TotalResults,
		Articles:     make([]domain.Article, len(newsResponse.Articles)),
	}

	logger.Info(fmt.Sprintf("Domain Response being created: Status=%s, TotalResults=%d, Articles=%d", newsResponseDomain.Status, newsResponseDomain.TotalResults, len(newsResponseDomain.Articles)))

	for i, article := range newsResponse.Articles {
		source := ""
		if article.Source.Name != "" {
			source = article.Source.Name
		}
		newsResponseDomain.Articles[i] = domain.Article{
			Source:      source,
			Id:          article.Id,
			Name:        article.Name,
			Author:      article.Author,
			Title:       article.Title,
			Description: article.Description,
			Url:         article.Url,
			PublishedAt: article.PublishedAt,
			UrlToImage:  article.UrlToImage,
			Content:     article.Content,
		}
	}

	return newsResponseDomain, nil
}
