package repo

import "errors"

type tansaction string

const (
	tableUser                    = "users"
	tableToken                   = "token"
	tableDocument                = "document"
	tableGrant                   = "grants"
	suffixReturningID            = "RETURNING id"
	tansactionKey     tansaction = "tansactionSQL"
)

var ErrTokenNotFound = errors.New("token not found")
