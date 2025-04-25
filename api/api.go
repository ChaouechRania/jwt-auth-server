package api

import (
	"jwt-auth-server/errors"

	"github.com/asaskevich/govalidator"
)

// SignUpRequest signup request
type SignUpRequest struct {
	Password  string `json:"password,omitempty" valid:"required"`
	Email     string `json:"email,omitempty" valid:"email,required"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

func (req *SignUpRequest) Validate() error {
	// Validate the struct using govalidator rules
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return err
	}
	// Validate email format explicitly if needed
	if req.Email != "" {
		if !govalidator.IsEmail(req.Email) {
			return errors.ErrEmail
		}
	}

	return nil
}

// Login represents auth data
type Login struct {
	Email    string `json:"email" valid:"email~Invalid email"`
	Password string `json:"password" valid:"required~The password is required"`
}
