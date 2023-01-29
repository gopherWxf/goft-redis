package gedis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
	"time"
)

var redisClient *redis.Client
var redisClientOnce sync.Once

func Redis() *redis.Client {
	redisClientOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Network:  "tcp",
			Addr:     "127.0.0.1:55002",
			Password: "", //密码
			DB:       0,  // redis数据库

			//连接池容量及闲置连接数量
			PoolSize:     15, // 连接池数量
			MinIdleConns: 10, //好比最小连接数
			//超时
			DialTimeout:  5 * time.Second, //连接建立超时时间
			ReadTimeout:  3 * time.Second, //读超时，默认3秒， -1表示取消读超时
			WriteTimeout: 3 * time.Second, //写超时，默认等于读超时
			PoolTimeout:  4 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。

		})
		pong, err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			log.Fatal(fmt.Errorf("connect error:%s", err))
		}
		log.Println(pong)
	})
	return redisClient
}
