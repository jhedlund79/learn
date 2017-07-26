package myserver

import (
	"fmt"
	"net/http"
	"sync"
)

type MyHandler struct {
	sync.Mutex
	count int
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var count int
	h.Lock()
	h.count++
	count = h.count
	h.Unlock()
	fmt.Fprintf(w, "request uri is %s\n\n", r.RequestURI)

	if r.RequestURI == "/products" {
		fmt.Fprintf(w, "products count: %d.", count)
	}

	if r.RequestURI == "/contacts" {
		fmt.Fprintf(w, "contacts count: %d.", count)
	}

}
