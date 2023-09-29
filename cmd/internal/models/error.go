package models

import "errors"

var (
	ErrNoRecond           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("authentication: invalid credentials")
	ErrDuplicateEmail     = errors.New("authentication: email already in use")
	ErrRestrictedAcc      = errors.New("authentication: account is restricted")
)
