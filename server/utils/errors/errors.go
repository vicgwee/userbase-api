package utilsErrors

import (
	"net/http"

	"github.com/pkg/errors"
)

type InternalError struct{}
type MongoDbError struct{}
type UserInvalidError struct{}
type UserNotFoundError struct{}
type UserDuplicateError struct{}

func (e *InternalError) Error() string {
	return "Internal Error"
}
func NewInternalError(fmt string, args ...interface{}) error {
	return errors.Wrapf(&InternalError{}, fmt, args...)
}

func (e *MongoDbError) Error() string {
	return "MongoDB Error"
}
func NewMongoDbError(fmt string, args ...interface{}) error {
	return errors.Wrapf(&MongoDbError{}, fmt, args...)
}

func (e *UserInvalidError) Error() string {
	return "UserInvalid Error"
}
func NewUserInvalidError(fmt string, args ...interface{}) error {
	return errors.Wrapf(&UserInvalidError{}, fmt, args...)
}

func (e *UserNotFoundError) Error() string {
	return "UserNotFound Error"
}
func NewUserNotFoundError(fmt string, args ...interface{}) error {
	return errors.Wrapf(&UserNotFoundError{}, fmt, args...)
}

func (e *UserDuplicateError) Error() string {
	return "UserDuplicate Error"
}
func NewUserDuplicateError(fmt string, args ...interface{}) error {
	return errors.Wrapf(&UserDuplicateError{}, fmt, args...)
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	//TODO: Implement other status codes
	return http.StatusInternalServerError
}
