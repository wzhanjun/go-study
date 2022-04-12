package trie

import (
	"fmt"
	"testing"
)

func TestRouter(t *testing.T) {

	router := &router{
		root:  make(map[string]*node),
		route: make(map[string]HandleFunc),
	}

	router.addRoute("GET", "/tomato/wiki-classify/migrate", nil)
	router.addRoute("GET", "/tomato/wiki-classify/:id", nil)
	router.addRoute("GET", "/tomato/wiki-classify/tree", nil)
	router.addRoute("GET", "/tomato/short-code", nil)

	node, params := router.getRoute("GET", "/tomato/wiki-classify/tree")
	fmt.Printf("node:%#v, params:%#v \n", node, params)

	node, params = router.getRoute("GET", "/tomato/wiki-classify/50")
	fmt.Printf("node:%#v, params:%#v \n", node, params)
}
