package errs

import "github.com/dayu-go/gkit/errors"

var (
	ErrAuthFail = errors.New(401, "Authentication failed", "Missing token or token incorrect")

	ErrInternal = errors.New(1001, "操作失败，请稍后再试", "internal error")

	ErrInvalidParam = errors.New(2001, "parse params failed", "invalid params")
)
