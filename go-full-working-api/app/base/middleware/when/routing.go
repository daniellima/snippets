package when

import (
	"net/http"

	"golang.org/x/exp/maps"
)

type MethodToHandlerMap map[string]http.HandlerFunc

func (m MethodToHandlerMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler

	handler, ok := m[r.Method]
	if !ok {
		handler = http.NotFoundHandler()
	}

	handler.ServeHTTP(w, r)
}

func (m MethodToHandlerMap) And(handlerMap MethodToHandlerMap) MethodToHandlerMap {

	maps.Copy(m, handlerMap)

	return m
}

func Get(handler http.HandlerFunc) MethodToHandlerMap {
	return MethodToHandlerMap{
		"GET": handler,
	}
}

func Post(handler http.HandlerFunc) MethodToHandlerMap {
	return MethodToHandlerMap{
		"POST": handler,
	}
}
