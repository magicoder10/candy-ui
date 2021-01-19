package router

import (
	"fmt"
	"strings"
)

type Route struct {
	Path   string
	Object interface{}
}

type trieNode struct {
	endOfPath bool
	route     *Route
	children  map[rune]*trieNode
}

type Router struct {
	root      *trieNode
	currRoute *Route
}

func (r *Router) AddRoute(route Route) error {
	if !isPathValid(route.Path) {
		return fmt.Errorf("path is invalid: %s", route.Path)
	}

	path := preparePath(route.Path)
	root := r.root
	for index, currRune := range path {
		if child, ok := root.children[currRune]; ok {
			root = child
		} else {
			child := newTrieNode()
			root.children[currRune] = child
			root = child
		}

		if index < len(path)-1 {
			continue
		} else if root.endOfPath {
			return fmt.Errorf("path already exists: %s", route.Path)
		} else {
			root.endOfPath = true
			root.route = &route
		}
	}
	return nil
}

func (r *Router) AddRoutes(routes []Route) error {
	for _, rt := range routes {
		err := r.AddRoute(rt)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r Router) CurrentRoute() *Route {
	return r.currRoute
}

func (r *Router) Navigate(path string) error {
	rt, err := r.matchRoute(path)
	if err != nil {
		return err
	}
	r.currRoute = rt
	return nil
}

func (r Router) matchRoute(path string) (*Route, error) {
	if !isPathValid(path) {
		return nil, fmt.Errorf("path is invalid: %s", path)
	}

	path = preparePath(path)
	root := r.root
	for index, currRune := range path {
		child, ok := root.children[currRune]
		if !ok {
			break
		}
		root = child
		if index == len(path)-1 && root.endOfPath {
			return root.route, nil
		}
	}
	return nil, fmt.Errorf("path not found: %s", path)
}

func preparePath(path string) string {
	for len(path) > 1 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	return path
}

var blockedRunes = map[rune]struct{}{}

func isPathValid(path string) bool {
	if !strings.HasPrefix(path, "/") {
		return false
	}
	for _, currRune := range path {
		if _, ok := blockedRunes[currRune]; ok {
			return false
		}
	}
	return true
}

func NewRouter() *Router {
	return &Router{
		root: newTrieNode(),
	}
}

func newTrieNode() *trieNode {
	return &trieNode{children: make(map[rune]*trieNode)}
}
