package kee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandleFunc
}

func New() *Engine {
	return &Engine{
		router: make(map[string]HandleFunc),
	}
}

func (e *Engine) addRouter(method string, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

func (e *Engine) Get(pattern string, handler HandleFunc) {
	e.addRouter("GET", pattern, handler)
}

func (e *Engine) Post(pattern string, handler HandleFunc) {
	e.addRouter("POST", pattern, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	fmt.Println(key, e.router)
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprint(w, "404")
	}
}

func (e *Engine) Run(port string) {
	http.ListenAndServe(port, e)
}
