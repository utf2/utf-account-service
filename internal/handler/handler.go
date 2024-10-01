package handler

import (
	"encoding/json"
	"net/http"

	httperror "github.com/utf2/utf-account-service/internal/error"
)

type HttpHandlerWithError func(w http.ResponseWriter, r *http.Request) *httperror.HttpError

func WrapHttpHandlerFunc(f HttpHandlerWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpError := f(w, r)
		if httpError == nil {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpError.StatusCode)
		json.NewEncoder(w).Encode(struct {
			ErrorMessage string
		}{
			ErrorMessage: httpError.ErrorMessage,
		})
	}
}
