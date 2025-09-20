package router

import (
	"net/http"
	"strings"
	"url-shortener/internal/adapter/http/handler"
	"url-shortener/internal/adapter/http/middleware"
)

type Router struct {
	urlHandler *handler.URLHandler
}

func NewRouter(urlHandler *handler.URLHandler) *Router {
	return &Router{
		urlHandler: urlHandler,
	}
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	middleware.CORS(http.HandlerFunc(rt.route)).ServeHTTP(w, r)
}

func (rt *Router) route(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch {
	case path == "/health":
		rt.urlHandler.Health(w, r)
	case path == "/api/v1/shorten":
		rt.urlHandler.ShortenURL(w, r)
	case strings.HasPrefix(path, "/api/v1/stats/"):
		rt.urlHandler.GetStats(w, r)
	case path != "/" && !strings.HasPrefix(path, "/api/"):
		rt.urlHandler.RedirectURL(w, r)
	default:
		http.NotFound(w, r)
	}
}
