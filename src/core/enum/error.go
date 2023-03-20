package enum

import (
	"github.com/pkg/errors"
	"gitlab.com/golibs-starter/golib/exception"
)

var ErrResourceNotFound = errors.New("resource not found")
var ErrMissingUsername = exception.New(4000100, "missing username")
var ErrInvalidUsername = exception.New(4000101, "invalid username")
var ErrInvalidDuration = exception.New(4000102, "invalid duration")
var ErrBillingNotFound = exception.New(4040100, "billing not found")
