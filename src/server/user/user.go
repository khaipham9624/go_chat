package user

import (
	"fmt"
	"gochat/src/server/db"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserLogin struct {
	UserName string
	Password string
}

type UserRegister struct {
	UserName string
	Password string
	FullName string
	Email    string
}

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

func (u *UserRegister) Register() bool {
	HashedPassword, err := hashPassword(u.Password)
	if err != nil {
		fmt.Println("Failed to hash password")
		return false
	}
	userId := uuid.New().String()
	if (!db.CreateUserLogin(db.UserLoginInfo{
		UserId:         userId,
		UserName:       u.UserName,
		HashedPassword: HashedPassword,
	})) {
		return false
	}

	if (!db.CreateUserInfo(db.UserInfo{
		UserId:   userId,
		UserName: u.UserName,
		Email:    u.Email,
	})) {
		return false
	}
	return true
}

func (u *UserLogin) Login() bool {
	userLoginFromDb := db.ReadUserLogin(u.UserName)
	err := bcrypt.CompareHashAndPassword([]byte(userLoginFromDb.HashedPassword), []byte(u.Password))
	return err == nil
}
