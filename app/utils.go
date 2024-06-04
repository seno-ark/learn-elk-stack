package app

import (
	"encoding/json"
	"errors"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"os"
)

type ContextKey string

const ContextRequestID ContextKey = "request_id"

var (
	Logger            *slog.Logger
	ErrBadRequest     = errors.New("invalid data")
	ErrUnauthorized   = errors.New("unaturhorized")
	ErrNotFound       = errors.New("not found")
	ErrInternalServer = errors.New("internal server error")
)

func init() {
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func JsonResp(w http.ResponseWriter, status int, resp any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func randomInt(min int, max int) int {
	// min: included
	// max: excluded
	return rand.IntN(max-min) + min
}

func ErrStatusMessage(err error) (int, string) {
	if errors.Is(err, ErrBadRequest) {
		return http.StatusBadRequest, err.Error()
	} else if errors.Is(err, ErrUnauthorized) {
		return http.StatusUnauthorized, err.Error()
	} else if errors.Is(err, ErrNotFound) {
		return http.StatusNotFound, err.Error()
	} else if errors.Is(err, ErrInternalServer) {
		return http.StatusInternalServerError, err.Error()
	}
	return http.StatusInternalServerError, "Oops"
}
