
package kee

import (
	"net/http"
	"fmt"
)

type Handler func(http.ResponseWriter, *http.Request);

type Engine struct {
	routes map[string]Handler
}

func New() *Engine {
	return &Engine{
		routes: make(map[string]Handler),
	}
}

func (e *Engine) addRoute(method string, addr string, handler Handler) {
	key := method + "-" + addr;
	e.routes[key] = handler;
}

func (e *Engine) Get(addr string, handler Handler) {
	e.addRoute("GET", addr, handler)	
}

func (e *Engine) ServeHTTP (res http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path;
	if handler, ok := e.routes[key]; ok {
		handler(res, req)
	} else {
		fmt.Fprint(res, "404")
	}
}


func (e *Engine) Run(addr string){
	http.ListenAndServe(addr, e)
}

