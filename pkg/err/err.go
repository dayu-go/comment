package err

import "github.com/dayu-go/gkit/errors"

var (
	ErrAuthFail = errors.New(401, "Authentication failed", "Missing token or token incorrect")

	ErrInvalidParam = errors.New(2001, "parse params failed", "invalid params")
)
