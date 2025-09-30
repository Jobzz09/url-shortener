package storage

import (
	"database/sql"
	"errors"
)

var (
	ErrUrlNotFound = errors.New("url not found")
	ErrUrlExists   = errors.New("url exists")
)

type Storage struct {
	Db *sql.DB
	
}


