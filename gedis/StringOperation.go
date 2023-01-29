package gedis

import "context"

type StringOperation struct {
	ctx context.Context
}

func NewStringOperation() *StringOperation {
	return &StringOperation{ctx: context.Background()}
}

func (this *StringOperation) Set() {

}

func (this *StringOperation) Get(key string) *StringResult {
	return NewStringResult(Redis().Get(this.ctx, key).Result())
}
func (this *StringOperation) MGet(key ...string) *SliceResult {
	return NewSliceResult(Redis().MGet(this.ctx, key...).Result())
}
