package trie

import (
	"fmt"
	"testing"
)

func TestEngine(t *testing.T) {
	r := New()
	r.Use(Logger())
	v1 := r.Group("/v1")
	{
		v1.GET("/hello", func(c *Context) {
			fmt.Println("hello world")
		})
		v2 := v1.Group("/v2")
		v2.Use(OnlyForV2())
		v2.GET("/test", func(c *Context) {
			fmt.Println("hello world")
		})
	}
}
