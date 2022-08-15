package token

import "time"

// Maker interface is a general interface for different (JWT and Paseto implementations),
// which allows us to use different authorizations techniques without the need to change the whole code
type Maker interface {
	CreateToken(username string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}