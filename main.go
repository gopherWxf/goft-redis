package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gopherWxf/goft-redis/lib"
)

func main() {

	r := gin.New()
	r.Handle("GET", "/news/:id", func(ctx *gin.Context) {
		newsCache := lib.NewsCache()
		defer lib.ReleaseNewsCache(newsCache)

		newsID := ctx.Param("id")
		newsCache.DBGetter = lib.NewsDBGetter(newsID)

		//ctx.Header("Content-type", "application/json")
		//ctx.String(200, res.(string))
		newsModel := lib.NewNewsModel()
		newsCache.GetCacheForObject("news"+newsID, newsModel)
		ctx.JSON(200, newsModel)
	})
	r.Run(":8080")
}
