package routes

import (
	"news-api/adapter/input/controller"
	"news-api/adapter/output/news_http"
	"news-api/application/service"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	newsClient := news_http.NewNewsClient()
	newService := service.NewNewsService(newsClient)
	newsController := controller.NewNewsController(newService)

	r.GET("/news", newsController.GetNews)
}
