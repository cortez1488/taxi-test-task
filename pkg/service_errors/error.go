package service_errors

import "errors"

var (
	ErrNoData       = errors.New("data isn't exists")
	ErrUnableToSave = errors.New("unable to save")
)
