package gedis

import (
	"fmt"
	"time"
)

const (
	ATTR_EXPIRE = "expr"
	ATTR_NX     = "nx"
	ATTR_XX     = "xx"
)

type empty struct{}

type OperationAttr struct {
	Name  string
	Value interface{}
}
type OperationAttrs []*OperationAttr

func (this OperationAttrs) Find(name string) *InterfaceResult {
	for _, attr := range this {
		if attr.Name == name {
			return NewInterfaceResult(attr.Value, nil)
		}
	}
	return NewInterfaceResult(nil, fmt.Errorf("OperationAttrs found error:%s", name))
}
func WithExpire(t time.Duration) *OperationAttr {
	return &OperationAttr{
		Name:  ATTR_EXPIRE,
		Value: t,
	}
}
func WithNX() *OperationAttr {
	return &OperationAttr{
		Name:  ATTR_NX,
		Value: empty{},
	}
}
func WithXX() *OperationAttr {
	return &OperationAttr{
		Name:  ATTR_XX,
		Value: empty{},
	}
}
