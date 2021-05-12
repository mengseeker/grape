package auth

import (
	"encoding/json"
	"errors"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
)

type Token struct {
	ID        string `json:"id"`
	AppID     int    `json:"app_id"`
	ExpiredAt int64  `json:"expired_at"`
}

var (
	errUnauthorizedToken = errors.New("unauthorized token")
	errExpiredToken      = errors.New("token out of date")
	errAppRemoved        = errors.New("unauthorized user")

	tokens    = map[string]*Token{}
	tokenKeys = map[string]string{}
)

func (t *Token) Check() error {
	if t.ExpiredAt > 0 && t.ExpiredAt < time.Now().Unix() {
		return errExpiredToken
	}
	return nil
}

// check token then return app
func GetAppByToken(tokenid string) (*App, error) {
	token, ok := tokens[tokenid]
	if !ok {
		return nil, errUnauthorizedToken
	}
	if err := token.Check(); err != nil {
		return nil, err
	}
	app, ok := apps[token.AppID]
	if !ok {
		return nil, errAppRemoved
	}
	return app, nil
}

func (t *Token) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

func UnmarshalToken(raw []byte) (*Token, error) {
	var a Token
	err := json.Unmarshal(raw, &a)
	return &a, err
}

func SetupToken(kv *mvccpb.KeyValue) error {
	t, err := UnmarshalToken(kv.Value)
	if err != nil {
		return err
	}
	tokens[t.ID] = t
	tokenKeys[string(kv.Key)] = t.ID
	log.Infof("token %s added", t.ID)
	return nil
}

func RemoveToken(kv *mvccpb.KeyValue) error {
	tid := tokenKeys[string(kv.Key)]
	delete(tokenKeys, string(kv.Key))
	delete(tokens, tid)
	log.Infof("token %s removed", tid)
	return nil
}

func UpdateToken(kv *mvccpb.KeyValue) error {
	// no way
	return nil
}
