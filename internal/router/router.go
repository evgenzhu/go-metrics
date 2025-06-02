package router

import (
	"github.com/evgenzhu/go-metrics-app/internal/handlers"
	"github.com/evgenzhu/go-metrics-app/internal/storage"
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	h := handlers.NewHandler(storage.NewMemoryStorage())

	mux.HandleFunc("/update/", h.HandleUpdate)

	return mux
}
