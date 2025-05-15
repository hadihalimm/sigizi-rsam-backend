package config

import (
	"net/http"
	"time"

	"github.com/gorilla/sessions"
)

var SessionStore = sessions.NewCookieStore([]byte("tes123"))

func InitSession() {
	SessionStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int((8 * time.Hour).Seconds()),
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
}
