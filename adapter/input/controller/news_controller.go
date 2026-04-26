package controller

import (
	"news-api/adapter/input/model/request"
	"news-api/application/domain"
	"news-api/application/port/input"
	"news-api/configuration/logger"
	"news-api/configuration/validation"

	"github.com/gin-gonic/gin"
)

type newsController struct {
	newsUseCase input.NewsUseCase
}

func NewNewsController(newsUsecase input.NewsUseCase) *newsController {
	return &newsController{
		newsUseCase: newsUsecase,
	}
}

func (nc *newsController) GetNews(c *gin.Context) {
	//q=tesla&from=2026-03-25&apiKey=dad0524823ff406da7182d0059940e80
	logger.Info("Received request for news")
	request := request.NewsRequest{}

	if err := c.ShouldBindQuery(&request); err != nil {
		logger.Error("Error trying to validate fields from request", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(400, gin.H{"error": errRest})
		return
	}

	newsDomain := domain.NewsReqDomain{
		Query: request.Query,
		From:  request.From,
	}

	_, _ = nc.newsUseCase.GetNews(newsDomain)

	c.JSON(200, newsDomain)
}
