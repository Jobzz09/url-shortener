package storage

import "errors"

var (
	ErrUrlNotFound = errors.New("Url not found!")
	ErrUrlExists   = errors.New("Url exists")
)
