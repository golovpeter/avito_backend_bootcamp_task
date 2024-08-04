package common

import "errors"

var (
	ErrUserAlreadyExist     = errors.New("user already exist")
	ErrUserNotExist         = errors.New("user does not exist")
	ErrIncorrectCredentials = errors.New("invalid credentials")
	ErrCompileRegexp        = errors.New("failed to compile regexp")
)

type Err5xx struct {
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Code      int    `json:"code"`
}
