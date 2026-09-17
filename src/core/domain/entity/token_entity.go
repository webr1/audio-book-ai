package entity

import "time"

type TokenEntity struct {
	Exp     time.Time
	Subject string
	Payload map[string]interface{}
}

func NewTokenEntity(exp time.Time, subject string, payload map[string]interface{}) *TokenEntity {
	return &TokenEntity{Exp: exp, Subject: subject, Payload: payload}
}
