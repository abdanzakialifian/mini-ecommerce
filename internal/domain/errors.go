package domain

import "errors"

var ErrProductAlreadyExists = errors.New("A product with the same identifier already exists")
var ErrProductNotFound = errors.New("Product with the given identifier was not found")
var ErrCategoryAlreadyExists = errors.New("A category with the same identifier already exists")
var ErrCategoryNotFound = errors.New("Category with the given identifier was not found")
var ErrUserAlreadyExists = errors.New("User with this email already exists")
var ErrUserNotFound = errors.New("User not found")
var ErrUserInvalid = errors.New("Invalid email or password")
