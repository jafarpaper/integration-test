package exception

import "errors"

var NotFoundError = errors.New("user not found")
var InvalidStatePrecondition = errors.New("invalid state precondition ")
var DuplicateReferenceIdError = errors.New("the provided id has been used before")
