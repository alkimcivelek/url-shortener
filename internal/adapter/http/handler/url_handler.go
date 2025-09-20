package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"url-shortener/internal/adapter/dto"
	"url-shortener/internal/application/port"
	"url-shortener/internal/domain/repository"
	"url-shortener/internal/domain/valueobject"
)

type URLHandler struct {
	urlService port.URLService
	baseURL    string
}

func NewURLHandler(urlService port.URLService, baseURL string) *URLHandler {
	return &URLHandler{
		urlService: urlService,
		baseURL:    strings.TrimSuffix(baseURL, "/"),
	}
}

func (h *URLHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Only POST method is allowed")
		return
	}

	var req dto.ShortenURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid_request", "Invalid JSON format")
		return
	}

	if strings.TrimSpace(req.URL) == "" {
		h.writeError(w, http.StatusBadRequest, "missing_url", "URL field is required")
		return
	}

	url, err := h.urlService.ShortenURL(req.URL)
	if err != nil {
		h.handleDomainError(w, err)
		return
	}

	response := dto.URLEntityToShortenResponse(url, h.baseURL)
	h.writeJSON(w, http.StatusOK, response)
}

func (h *URLHandler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Only GET method is allowed")
		return
	}

	shortCode := strings.TrimPrefix(r.URL.Path, "/")
	if shortCode == "" || strings.HasPrefix(shortCode, "api/") || shortCode == "health" {
		http.NotFound(w, r)
		return
	}

	url, err := h.urlService.RedirectURL(shortCode)
	if err != nil {
		h.handleDomainError(w, err)
		return
	}

	http.Redirect(w, r, url.OriginalURL().Value(), http.StatusMovedPermanently)
}

func (h *URLHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Only GET method is allowed")
		return
	}

	shortCode := strings.TrimPrefix(r.URL.Path, "/api/v1/stats/")
	if shortCode == "" {
		h.writeError(w, http.StatusBadRequest, "missing_code", "Short code is required")
		return
	}

	url, err := h.urlService.GetURLStats(shortCode)
	if err != nil {
		h.handleDomainError(w, err)
		return
	}

	response := dto.URLEntityToStatsResponse(url)
	h.writeJSON(w, http.StatusOK, response)
}

func (h *URLHandler) Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "Only GET method is allowed")
		return
	}

	response := dto.HealthResponse{
		Status:    "healthy",
		Service:   "url-shortener",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	h.writeJSON(w, http.StatusOK, response)
}

func (h *URLHandler) handleDomainError(w http.ResponseWriter, err error) {
	switch err {
	case valueobject.ErrEmptyURL, valueobject.ErrInvalidURL, valueobject.ErrUnsupportedScheme:
		h.writeError(w, http.StatusBadRequest, "invalid_url", err.Error())
	case valueobject.ErrEmptyShortCode:
		h.writeError(w, http.StatusBadRequest, "invalid_short_code", err.Error())
	case repository.ErrURLNotFound:
		h.writeError(w, http.StatusNotFound, "not_found", "Short URL not found")
	case repository.ErrDuplicateURL:
		h.writeError(w, http.StatusConflict, "duplicate_url", "URL already exists")
	default:
		log.Printf("Unexpected error: %v", err)
		h.writeError(w, http.StatusInternalServerError, "internal_error", "Internal server error")
	}
}

func (h *URLHandler) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

func (h *URLHandler) writeError(w http.ResponseWriter, status int, error, message string) {
	response := dto.ErrorResponse{
		Error:   error,
		Message: message,
	}
	h.writeJSON(w, status, response)
}
