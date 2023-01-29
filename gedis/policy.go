package gedis

import (
	"fmt"
	"regexp"
	"time"
)

type CachePolicy interface {
	Before(key string)
	IfNil(key string, v interface{})
	SetOperation(opt *StringOperation)
}

// 缓存穿透策略
type CrossPolicy struct {
	KeyRegx string
	opt     *StringOperation
	Expire  time.Duration
}

func (this *CrossPolicy) IfNil(key string, v interface{}) {
	this.opt.Set(key, v, WithExpire(this.Expire)).Unwrap()
}

func (this *CrossPolicy) SetOperation(opt *StringOperation) {
	this.opt = opt
}

func NewCrossPolicy(keyRegx string, Expire time.Duration) *CrossPolicy {
	return &CrossPolicy{KeyRegx: keyRegx, Expire: Expire}
}

func (this *CrossPolicy) Before(key string) {
	if !regexp.MustCompile(this.KeyRegx).MatchString(key) {
		fmt.Println(key)
		fmt.Println(this.KeyRegx)
		panic("error cache key")
	}
}
