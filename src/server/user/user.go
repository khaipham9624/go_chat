package user

import (
	"gochat/src/server/db"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

type User struct {
	Id       uuid.UUID
	UserName string
	FullName string
}

func (u *User) WriteToDb() {
	db.CreateUserLogin(db.UserLoginInfo{})
	db.CreateUserInfo(db.UserInfo{})
}
