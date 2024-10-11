package auth_test

import (
	"testing"
	"todo/internal/auth"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockStorager struct {
	mock.Mock
}

func (m *MockStorager) InsertUser(username, hashPassword string) (string, error) {
	args := m.Called(username, hashPassword)
	return args.String(0), args.Error(1)
}

// TestCanRegisterUserWithValidPassword - проверяет, что
// пользователь успешно регистриурется при использовании
// безопасного пароля.
// Безопасный пароль:
//
// - Больше или равно 8ми символам
func TestCanRegisterUserWithValidPassword(t *testing.T) {
	testPassword := "qwerty1234"
	mockRegister := new(MockStorager)
	mockRegister.On("InsertUser", "aboba", mock.Anything).Return(mock.Anything, nil)
	user := auth.NewUser("aboba", testPassword)
	err := auth.RegisterUser(user, mockRegister)
	require.NoError(t, err)
}

// TestRegisterUserCannotRegisterUserWithShortPassword проверяет,
// что если длина пароля меньше 8 символов, то пользователь не сможет
// зарегестрироваться.
func TestRegisterUserCannotRegisterUserWithShortPassword(t *testing.T) {
	testCases := []struct {
		name     string
		password string
	}{
		{"one symbol", "1"},
		{"7 symbols", "1111111"},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			testPassword := tt.password
			mockRegister := new(MockStorager)
			mockRegister.On("InsertUser", "aboba", mock.Anything).Return(mock.Anything, nil)
			user := auth.NewUser("aboba", testPassword)
			err := auth.RegisterUser(user, mockRegister)
			require.ErrorIs(t, err, auth.ErrPasswordToShort)
		})
	}
}
