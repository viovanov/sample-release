package viewer

import (
	"net/http"
	"time"

	"github.com/pivotal-golang/lager"
)

type updateHandler struct {
	logger     lager.Logger
	flowViewer chan (string)
}

func NewUpdateHandler(logger lager.Logger, flowViewer chan (string)) http.Handler {
	return &updateHandler{
		logger:     logger.Session("hello-diego-viewer"),
		flowViewer: flowViewer,
	}
}

func (handler *updateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.WriteHeader(http.StatusOK)

	select {
	case res := <-handler.flowViewer:
		w.Write([]byte(res))
	case <-time.After(time.Millisecond * 50):
		w.Write([]byte(``))
		return
	}
}
