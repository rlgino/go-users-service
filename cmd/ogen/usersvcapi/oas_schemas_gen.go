// Code generated by ogen, DO NOT EDIT.

package usersvcapi

import (
	"fmt"

	"github.com/google/uuid"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// AddUserCreated is response for AddUser operation.
type AddUserCreated struct{}

// DeleteUserOK is response for DeleteUser operation.
type DeleteUserOK struct{}

// Represents error object.
// Ref: #/components/schemas/Error
type Error struct {
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// Ref: #/components/schemas/User
type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
}

// GetID returns the value of ID.
func (s *User) GetID() uuid.UUID {
	return s.ID
}

// GetName returns the value of Name.
func (s *User) GetName() string {
	return s.Name
}

// GetPassword returns the value of Password.
func (s *User) GetPassword() string {
	return s.Password
}

// SetID sets the value of ID.
func (s *User) SetID(val uuid.UUID) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *User) SetName(val string) {
	s.Name = val
}

// SetPassword sets the value of Password.
func (s *User) SetPassword(val string) {
	s.Password = val
}

type Users []User
