package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	root := &Trie{
		root: &Node{},
	}

	root.Insert("tomato")
	root.Insert("toba")
	root.Insert("abcd")
	root.Insert("hello")

	fmt.Printf("search: %s, result:%t \n", "to", root.Search("to"))
	fmt.Printf("search: %s, result:%t \n", "cd", root.Search("cd"))
	fmt.Printf("search: %s, result:%t \n", "abc", root.Search("abc"))
}
