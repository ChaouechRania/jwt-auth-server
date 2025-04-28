package errors

import "errors"

// ErrEmailTaken is returned when trying to create a user the a taken email address
var ErrEmailTaken = errors.New("Email address already taken")

// ErrEmail is returned when the email address is not valid
var ErrEmail = errors.New("Invalid email format")

// ErrNoSuchEntity is returned when the entity does not exist
var ErrNoSuchEntity = errors.New("No such entity")

// Email not found
var ErrEmailNotFound = errors.New("Email not found")

// ErrNoSuchEntity is returned when the entity does not exist
var ErrNoSuchPassword = errors.New("Invalid password")
