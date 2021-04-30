package session

import (
	"errors"
	"grape/pkg/rediskv"
)

const (
	TTL = 3600 * 3
)

var (
	ErrSessionNotFound = errors.New("session not found")
)

func Save(s *Session) {
	_, err := rediskv.GetClient().Do("SETEX", s.ID, TTL, s.Marshal())
	if err != nil {
		panic(err)
	}
}

func Find(id string) (*Session, error) {
	reply, err := rediskv.GetClient().Do("GET", id)
	if err != nil {
		panic(err)
	}
	if reply == nil {
		return nil, ErrSessionNotFound
	}
	return UnMarshal([]byte(reply.([]byte))), nil
}
