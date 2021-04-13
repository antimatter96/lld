package custom_errors

import "errors"

var (
	EmptyString       error = errors.New("enter a name")
	NameTaken         error = errors.New("name taken")
	TooFew            error = errors.New("too few characters")
	InvalidCharacters error = errors.New("invalid characters")
	AlreadyPicked     error = errors.New("this has already been marked")
)
