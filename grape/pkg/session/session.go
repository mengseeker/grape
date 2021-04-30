package session

import (
	"encoding/json"

	"github.com/gofrs/uuid"
)

type Session struct {
	ID     string                 `json:"id"`
	UserID int                    `json:"user_id"`
	Data   map[string]interface{} `json:"data"`
}

func NewSession(userid int) *Session {
	uid, _ := uuid.NewV4()
	s := Session{
		ID:     uid.String(),
		UserID: userid,
		Data:   map[string]interface{}{},
	}
	return &s
}

func (s *Session) Marshal() []byte {
	ms, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return ms
}

func UnMarshal(raw []byte) *Session {
	s := &Session{}
	err := json.Unmarshal(raw, s)
	if err != nil {
		panic(err)
	}
	return s
}

func (s *Session) UnSaveSet(key string, val interface{}) {
	if s.Data == nil {
		s.Data = map[string]interface{}{}
	}
	s.Data[key] = val
}

func (s *Session) Set(key string, val interface{}) {
	s.UnSaveSet(key, val)
	Save(s)
}

func (s *Session) Get(key string) interface{} {
	return s.Data[key]
}
