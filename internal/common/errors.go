package common

import "errors"

var (
	ErrUserAlreadyExist = errors.New("user already exist")
	ErrUserNotExist     = errors.New("user does not exist")
)

type Err5xx struct {
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
	Code      int    `json:"code"`
}
