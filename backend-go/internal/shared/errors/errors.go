package errors

import "errors"

var ErrPostNotFound = errors.New("post not found")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrUnauthorized = errors.New("unauthorized")