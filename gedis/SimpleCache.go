package gedis

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"time"
)

type DBGetterFunc func() interface{}

const Serializer_JSON = "json"
const Serializer_GOB = "gob"

type SimpleCache struct {
	Operation  *StringOperation
	Expire     time.Duration
	DBGetter   DBGetterFunc
	Serializer string
}

func NewSimpleCache(operation *StringOperation, expire time.Duration, serializer string) *SimpleCache {
	return &SimpleCache{Operation: operation, Expire: expire, Serializer: serializer}
}

// 设置缓存
func (this *SimpleCache) SetCache(key string, val interface{}) {
	this.Operation.Set(key, val, WithExpire(this.Expire)).Unwrap()
}
func (this *SimpleCache) GetCache(key string) (ret interface{}) {
	if this.Serializer == Serializer_JSON {
		f := func() string {
			obj := this.DBGetter()
			b, err := json.Marshal(obj)
			if err != nil {
				return ""
			}
			return string(b)
		}
		ret = this.Operation.Get(key).UnwrapOrElse(f)
		this.SetCache(key, ret)
	} else if this.Serializer == Serializer_GOB {
		f := func() string {
			obj := this.DBGetter()
			buf := &bytes.Buffer{}
			enc := gob.NewEncoder(buf)
			if err := enc.Encode(obj); err != nil {
				return ""
			}
			return buf.String()
		}
		ret = this.Operation.Get(key).UnwrapOrElse(f)
		this.SetCache(key, ret)
	}

	return
}
func (this *SimpleCache) GetCacheForObject(key string, obj interface{}) interface{} {
	ret := this.GetCache(key)
	if ret == nil {
		return nil
	}
	if this.Serializer == Serializer_JSON {
		err := json.Unmarshal([]byte(ret.(string)), obj)
		if err != nil {
			return nil
		}
	} else if this.Serializer == Serializer_GOB {
		var buf = &bytes.Buffer{}
		buf.WriteString(ret.(string))
		dec := gob.NewDecoder(buf)
		if dec.Decode(obj) != nil {
			return nil
		}
	}
	return nil

}
