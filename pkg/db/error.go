package db

import "errors"

// Обозначение ошибок
var (
	ErrMigrate       = errors.New("migration failed")
	ErrDuplicate     = errors.New("record already exists")
	ErrNotExist      = errors.New("row does not exist")
	ErrUpdateFailed  = errors.New("update failed")
	ErrDeleteFailed  = errors.New("delete failed")
	ErrOwnerNotFound = errors.New("owner not found")
)
