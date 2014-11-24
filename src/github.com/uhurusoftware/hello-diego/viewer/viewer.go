package viewer

import (
	"net/http"

	"github.com/pivotal-golang/lager"
	"github.com/tedsuo/rata"
)

func NewServer(logger lager.Logger, flowViewer chan (string)) (http.Handler, error) {
	handlers := map[string]http.Handler{
		"simple_get": NewRootHandler(logger),
		"update_get": NewUpdateHandler(logger, flowViewer),
	}

	routes := rata.Routes{
		{Name: "simple_get", Method: "GET", Path: "/"},
		{Name: "update_get", Method: "GET", Path: "/update"},
	}

	return rata.NewRouter(routes, handlers)
}
