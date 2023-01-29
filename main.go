package main

import (
	"fmt"
	"github.com/gopherWxf/goft-redis/gedis"
)

func main() {
	iter := gedis.
		NewStringOperation().
		MGet("name", "age", "11").
		Iter()
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}
