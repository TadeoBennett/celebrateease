package initializers

import (
	// "flag"
	"time"

	"github.com/golangcollege/sessions"
)

func NewSession(secret string) *sessions.Session {
	//saves the session
	session := sessions.New([]byte(secret))
	session.Lifetime = 12 * time.Hour //session will expire after 12 hours
	//encrypted session keys
	session.Secure = true

	return session
}
