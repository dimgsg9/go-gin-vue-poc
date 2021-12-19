package model

import "github.com/google/uuid"

type RefreshToken struct {
	ID  uuid.UUID `json:"-"`
	UID uuid.UUID `json:"-"`
	SS  string    `json:"refreshToken"`
}

type IDToken struct {
	SS string `json:"idToken"`
}

type TokenPair struct {
	IDToken
	RefreshToken
}
