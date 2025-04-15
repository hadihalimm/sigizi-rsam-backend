package config

import (
	"time"

	"github.com/gorilla/sessions"
)

var SessionStore = sessions.NewCookieStore([]byte("tes123"))

func init() {
	SessionStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   int((12 * time.Hour).Seconds()),
		HttpOnly: true,
	}
}
