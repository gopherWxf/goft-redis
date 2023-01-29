package main

import (
	"fmt"
	"github.com/gopherWxf/goft-redis/gedis"
	"time"
)

func main() {
	//iter := gedis.
	//	NewStringOperation().
	//	MGet("name", "age", "11").
	//	Iter()
	//for iter.HasNext() {
	//	fmt.Println(iter.Next())
	//}
	fmt.Println(gedis.NewStringOperation().
		Set("bb", 213123,
			gedis.WithExpire(20*time.Second),
			gedis.WithXX(),
		))
}
