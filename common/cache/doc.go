package cache

import (
	"errors"
)

var (
	ErrNotFound = errors.New("cache: data not found")
	ErrSetValue = errors.New("cache: cannot set to value")
)
