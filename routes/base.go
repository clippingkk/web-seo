package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func indexHandler(c *gin.Context) {
	// 首页
	notfoundHandler(c)
}

func notfoundHandler(c *gin.Context) {
	c.Writer.Header().Set("content-type", "text/html")
	c.Writer.Write(pageCache)
}

func InitRoutes(indexPageTemplate string) *gin.Engine {
	// cronjob to refresh html template
	if err := refreshPageCache(indexPageTemplate); err != nil {
		logrus.Panicln(err)
	}
	go func() {
		loopRefreshPageCache(indexPageTemplate)
	}()

	h := gin.Default()

	h.GET("/", indexHandler)
	h.GET("*", notfoundHandler)
	return h
}
