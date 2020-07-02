package main

import (
	"fmt"
	"net/http"
	"kee"
)

func main() {
	r := kee.New();

	r.Get("/", func (w http.ResponseWriter, req *http.Request)  {
		fmt.Fprintf(w, "URL.PATH=%s", req.URL.Path)	
	})

	r.Get("/hello", func (w http.ResponseWriter, req *http.Request)  {
		for name, value := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", name, value)
		}
	})

	r.Run(":8666");
	
}