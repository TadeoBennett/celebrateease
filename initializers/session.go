package initializers

import (
	"flag"
	"time"

	"github.com/golangcollege/sessions"
)

func NewSession()(*sessions.Session) {
	//saves the session
	secret := flag.String("secret", "8693b89c15217db6a4a90aa41cf0e6d5f31752aaf318b4e184f7c5a93a9a90c2", "Secret Key")
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour //session will expire after 12 hours
	//encrypted session keys
	session.Secure = true

	return session
}
