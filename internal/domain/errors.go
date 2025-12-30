package domain

import "errors"

var ErrProductAlreadyExists = errors.New("A product with the same identifier already exists")
