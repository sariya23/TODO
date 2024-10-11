package auth

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrPasswordToShort = errors.New("password len is too short. Min: 8 symbols")
	ErrUserNotExists   = errors.New("user is not exists with this creds")
)

type UserInserter interface {
	InsertUser(username, hashPassword string) (string, error)
}

type UserGetter interface {
	GetUserPassword(userId string) (string, error)
	GetUser(username string) (string, error)
}

type User struct {
	Username string
	Password string
}

func NewUser(username string, password string) User {
	return User{
		Username: username,
		Password: password,
	}
}

func RegisterUser(u User, storage UserInserter) error {
	const op = "internal.auth.RegisterUser"
	if len(u.Password) < 8 {
		return fmt.Errorf("%v: %w", op, ErrPasswordToShort)
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%v: %w", op, err)
	}
	_, err = storage.InsertUser(u.Username, string(bytes))
	if err != nil {
		return fmt.Errorf("%v: err during register user: %w", op, err)
	}
	return nil
}

func LoginUser(u User, storage UserGetter) error {
	const op = "internal.auth.LoginUser"
	userId, err := storage.GetUser(u.Username)
	if err != nil {
		return fmt.Errorf("%v: %w", op, err)
	}
	if userId == "" {
		return fmt.Errorf("%v: %w", op, ErrUserNotExists)
	}
	password, err := storage.GetUserPassword(userId)
	if err != nil {
		return fmt.Errorf("%v: %w", op, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(u.Password))
	if err != nil {
		return fmt.Errorf("%v: wrong creds", op)
	}
	return nil
}
