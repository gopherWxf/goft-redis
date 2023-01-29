package lib

import (
	"github.com/gopherWxf/goft-redis/gedis"
	"log"
)

func NewsDBGetter(id string) gedis.DBGetterFunc {
	return func() interface{} {
		log.Printf("Get from DB")
		newsModel := NewNewsModel()
		Gorm.Table("mynews").Where("id=?", id).Find(newsModel)
		return newsModel
	}
}
