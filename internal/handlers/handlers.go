package handlers

import (
	"github.com/evgenzhu/go-metrics-app/internal/storage"
	"net/http"
	"strconv"
)

type Handler struct {
	storage storage.Storage
}

func NewHandler(store storage.Storage) *Handler {
	return &Handler{storage: store}
}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	parts, err := validatePath(r)
	if err != nil {
		switch err.Error() {
		case "incorrect URL structure: expected /update/{type}/{name}/{value}":
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}

	switch parts[1] {
	case "gauge":
		h.handleGauge(w, r, parts)
	case "counter":
		h.handleCounter(w, r, parts)
	default:
		http.Error(w, "unsupported metric type", http.StatusBadRequest)
	}
}

func (h *Handler) handleGauge(w http.ResponseWriter, r *http.Request, parts []string) {
	if err := validateMethod(r); err != nil {
		switch err.Error() {
		case "method not allowed":
			w.Header().Add("Allow", http.MethodPost)
			http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		case "unsupported content type":
			http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	metricValue, err := strconv.ParseFloat(parts[3], 64)
	if err != nil {
		http.Error(w, "invalid gauge metric value", http.StatusBadRequest)
		return
	}

	h.storage.SetGauge(parts[2], metricValue)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) handleCounter(w http.ResponseWriter, r *http.Request, parts []string) {
	if err := validateMethod(r); err != nil {
		switch err.Error() {
		case "method not allowed":
			w.Header().Add("Allow", http.MethodPost)
			http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		case "unsupported content type":
			http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	metricValue, err := strconv.ParseInt(parts[3], 10, 64)
	if err != nil {
		http.Error(w, "invalid counter metric value", http.StatusBadRequest)
		return
	}

	h.storage.AddCounter(parts[2], metricValue)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}
