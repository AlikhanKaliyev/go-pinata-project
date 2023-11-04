package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Pinatas PinataModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Pinatas: PinataModel{DB: db},
	}
}
