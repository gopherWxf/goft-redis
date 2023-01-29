package lib

import (
	"github.com/gopherWxf/goft-redis/gedis"
	"sync"
	"time"
)

var NewsCachePool *sync.Pool

func init() {
	NewsCachePool = &sync.Pool{New: func() any {
		return gedis.NewSimpleCache(gedis.NewStringOperation(), time.Second*15, gedis.Serializer_GOB)
	}}
}
func NewsCache() *gedis.SimpleCache {
	return NewsCachePool.Get().(*gedis.SimpleCache)
}
func ReleaseNewsCache(cache *gedis.SimpleCache) {
	NewsCachePool.Put(cache)
}
