package middlewares

import (
	"github.com/justinas/nosurf"
	"net/http"
)

func CSRFMiddleware(handler http.Handler) http.Handler {
	return nosurf.New(handler)
}
