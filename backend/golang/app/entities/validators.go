package entities

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

// InputValidation - an interface for all "input submission" structs used for
// deserialization.  We pass in the request so that we can potentially get the
// context by the request from our context manager
type InputValidation interface {
	Validate() error
}

var (
	// ErrInvalidName - error when we have an invalid name
	UserEmailReq   = errors.New("user email empty or invalid")
	UserNameReq = errors.New("user name empty or invalid")
	UserPasswordReq = errors.New("user password empty or invalid")
)

func (t UsersCreateInput) Validate() error {
	// validate the email is not empty or missing
	if !govalidator.IsEmail(t.Email) {
		return UserEmailReq
	}
	if govalidator.IsNull(t.Name) {
		return UserNameReq
	}
	if govalidator.IsNull(t.Password) {
		return UserPasswordReq
	}
	return nil
}

