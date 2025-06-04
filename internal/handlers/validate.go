package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func validateMethod(r *http.Request) error {
	if r.Method != http.MethodPost {
		return fmt.Errorf("only POST method is allowed")
	}

	if !strings.HasPrefix(r.Header.Get("Content-Type"), "text/plain") {
		return fmt.Errorf("unsupported content type, expected text/plain")
	}
	return nil
}

func validatePath(r *http.Request) ([]string, error) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	if len(parts) < 4 {
		return nil, fmt.Errorf("incorrect URL structure: expected /update/{type}/{name}/{value}")
	}
	return parts, nil
}
