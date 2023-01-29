package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gopherWxf/goft-redis/lib"
)

func main() {

	r := gin.New()
	r.Handle("GET", "/news/:id", func(ctx *gin.Context) {
		//1. 从对象池 获取新闻缓存 对象
		newsCache := lib.NewsCache()
		defer lib.ReleaseNewsCache(newsCache)

		//2. 获取参数，设置DBGetter
		newsID := ctx.Param("id")
		newsCache.DBGetter = lib.NewsDBGetter(newsID)

		//3. 取缓存输出，如果没有则调用上面的DBGetter
		//ctx.Header("Content-type", "application/json")
		//ctx.String(200, res.(string))
		newsModel := lib.NewNewsModel()
		newsCache.GetCacheForObject("news"+newsID, newsModel)
		ctx.JSON(200, newsModel)
	})
	r.Run(":8080")
}
