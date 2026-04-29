package errors

import "errors"

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLSaveFailed = errors.New("failed to save url")
	ErrURLResolveFailed = errors.New("failed to resolve url")
)