package models

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
)

func MiddlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logrus.New()

		contextWithLogger := context.WithValue(r.Context(), "log", log)
		requestWithLogger := r.WithContext(contextWithLogger)

		next.ServeHTTP(w,requestWithLogger)
	})
}