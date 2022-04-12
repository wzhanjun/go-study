package trie

import (
	"fmt"
	"net/http"
	"strings"
)

type HandleFunc func(*Context)

type node struct {
	path     string
	part     string
	children map[string]*node
	isWild   bool
}

type router struct {
	root  map[string]*node
	route map[string]HandleFunc
}

func newRouter() *router {
	return &router{
		root:  make(map[string]*node),
		route: make(map[string]HandleFunc),
	}
}

func (r *router) addRoute(method string, path string, handle HandleFunc) {
	parts := parsePath(path)
	if _, ok := r.root[method]; !ok {
		r.root[method] = &node{children: make(map[string]*node)}
	}

	root := r.root[method]
	key := method + "-" + path

	for _, part := range parts {
		if root.children[part] == nil {
			root.children[part] = &node{
				part:     part,
				children: make(map[string]*node),
				isWild:   part[0] == ':' || part[0] == '*',
			}
		}
		root = root.children[part]
	}

	root.path = path
	r.route[key] = handle
}

func (r *router) getRoute(method string, path string) (node *node, params map[string]string) {
	params = map[string]string{}
	searchParts := parsePath(path)

	var ok bool
	if node, ok = r.root[method]; !ok {
		return nil, nil
	}

	for i, part := range searchParts {
		var temp string

		for _, child := range node.children {
			fmt.Println(child.part)
			if child.part == part || child.isWild {
				if child.part[0] == '*' {
					params[child.part[1:]] = strings.Join(searchParts[i:], "/")
				}
				if child.part[0] == ':' {
					params[child.part[1:]] = part
				}
				temp = child.part
			}
		}

		if len(temp) > 0 && temp[0] == '*' {
			return node.children[temp], params
		}
		node = node.children[temp]
	}
	return
}

func parsePath(path string) (parts []string) {
	par := strings.Split(path, "/")
	for _, p := range par {
		if p != "" {
			parts = append(parts, p)
			if p[0] == '*' {
				break
			}
		}
	}
	return
}

func hello(c *Context) error {
	fmt.Println("this is hello")
	return nil
}

func fuzzy(c *Context) error {
	fmt.Println("this is fuzzy")
	return nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		key := c.Method + "-" + n.path
		c.Params = params
		c.handlers = append(c.handlers, r.route[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}

	c.Next()
}
