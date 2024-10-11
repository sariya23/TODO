package auth_test

import (
	"testing"
	"todo/internal/auth"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

type MockUserInserter struct {
	mock.Mock
}

func (m *MockUserInserter) InsertUser(username, hashPassword string) (string, error) {
	args := m.Called(username, hashPassword)
	return args.String(0), args.Error(1)
}

type MockUserGetter struct {
	mock.Mock
}

func (m *MockUserGetter) GetUser(username string) (string, error) {
	args := m.Called(username)
	return args.String(0), args.Error(1)
}

func (m *MockUserGetter) GetUserPassword(userId string) (string, error) {
	args := m.Called(userId)
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
	mockRegister := new(MockUserInserter)
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
			mockRegister := new(MockUserInserter)
			mockRegister.On("InsertUser", "aboba", mock.Anything).Return(mock.Anything, nil)
			user := auth.NewUser("aboba", testPassword)
			err := auth.RegisterUser(user, mockRegister)
			require.ErrorIs(t, err, auth.ErrPasswordToShort)
		})
	}
}

// TestCanLoginuserIfExiststAndCorrectPassword проверяет,
// что пользователь может залогиниться, если он существует
// и указан верный пароль.
func TestCanLoginUserIfExiststAndCorrectPassword(t *testing.T) {
	mockGetter := new(MockUserGetter)
	bytes, err := bcrypt.GenerateFromPassword([]byte("qwerty1234"), bcrypt.DefaultCost)
	require.NoError(t, err)
	mockGetter.On("GetUser", "aboba").Return("1", nil)
	mockGetter.On("GetUserPassword", "1").Return(string(bytes), nil)

	user := auth.NewUser("aboba", "qwerty1234")
	err = auth.LoginUser(user, mockGetter)
	require.NoError(t, err)
}

// TestCannotLoginUserBecauseNotExists проверяет,
// что если пользователь не существует, то
// он не сможет залогиниться.
func TestCannotLoginUserBecauseNotExists(t *testing.T) {
	mockGetter := new(MockUserGetter)
	mockGetter.On("GetUser", "ne aboba").Return("", auth.ErrUserNotExists)
	user := auth.NewUser("ne aboba", "qwerty1234")
	err := auth.LoginUser(user, mockGetter)
	require.ErrorIs(t, err, auth.ErrUserNotExists)
}

// TestCannotLoginUserBecauseWrongPassword проверяет,
// что если пользователь существует, но он указал неверный пароль,
// то он не сможет залогиниться.
func TestCannotLoginUserBecauseWrongPassword(t *testing.T) {
	mockGetter := new(MockUserGetter)
	bytes, err := bcrypt.GenerateFromPassword([]byte("qwerty1234"), bcrypt.DefaultCost)
	require.NoError(t, err)
	mockGetter.On("GetUser", "aboba").Return("1", nil)
	mockGetter.On("GetUserPassword", "1").Return(string(bytes), nil)
	user := auth.NewUser("aboba", "ne qwerty1234")
	err = auth.LoginUser(user, mockGetter)
	require.ErrorIs(t, err, auth.ErrUserNotExists)
}
