package router

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"jobqueue/handler"
)

func NewRouter(l *logrus.Logger) *mux.Router {
	router := mux.NewRouter()
	router.NotFoundHandler = handler.NotFoundHandler()

	router.HandleFunc("/health", handler.Health())

	return router
}
