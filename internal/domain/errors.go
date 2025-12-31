package domain

import "errors"

var ErrProductAlreadyExists = errors.New("A product with the same identifier already exists")
var ErrProductNotFound = errors.New("Product with the given identifier was not found")
