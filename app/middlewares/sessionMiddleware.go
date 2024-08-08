package middlewares

import (
	"github.com/alexedwards/scs/v2"
	"net/http"
)

var session *scs.SessionManager

func SetMainSession(manager *scs.SessionManager) {
	session = manager
}
func SessionMiddleware(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
