package models

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Logger struct {
	handler http.Handler
}

func NewLogger(handler http.Handler) *Logger {
	return &Logger{handler: handler}
}


func (lg *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log := logrus.New()

	contextWithLogger := context.WithValue(r.Context(), "log", log)
	requestWithLogger := r.WithContext(contextWithLogger)

	lg.handler.ServeHTTP(w, requestWithLogger)
}
