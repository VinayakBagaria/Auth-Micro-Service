package validators

import (
	"errors"
	"strings"

	"github.com/VinayakBagaria/auth-micro-service/pb"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrInvalidUserId      = errors.New("invalid user id")
	ErrEmptyName          = errors.New("name can't be empty")
	ErrEmptyEmail         = errors.New("email can't be empty")
	ErrEmptyPassword      = errors.New("password can't be empty")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrSignInFailed       = errors.New("signin failed")
)

func ValidateSignUp(user *pb.User) error {
	if !bson.IsObjectIdHex(user.Id) {
		return ErrInvalidUserId
	}
	if user.Name == "" {
		return ErrEmptyName
	}
	if user.Email == "" {
		return ErrEmptyEmail
	}
	if user.Password == "" {
		return ErrEmptyPassword
	}
	return nil
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
