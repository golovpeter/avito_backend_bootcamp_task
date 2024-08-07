package common

import "errors"

var (
	ErrUserAlreadyExist     = errors.New("user already exist")
	ErrUserNotExist         = errors.New("user does not exist")
	ErrIncorrectCredentials = errors.New("invalid credentials")
	ErrCompileRegexp        = errors.New("failed to compile regexp")
	ErrInvalidAuthHeader    = errors.New("invalid authorization header")
	ErrAccessDenied         = errors.New("access denied")
	ErrHouseAlreadyExist    = errors.New("house already exist")
	ErrFlatAlreadyExist     = errors.New("flat already exist")
)

type Err5xx struct {
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Code      int    `json:"code"`
}
