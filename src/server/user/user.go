package user

import (
	"gochat/src/server/db"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserLogin struct {
	UserName       string
	HashedPassword string
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

func (u *UserRegister) Register() {
	HashedPassword := hashPassword(u.Password)
	db.CreateUserLogin(db.UserLoginInfo{
		UserName:       u.UserName,
		HashedPassword: u.Password,
	})

	db.CreateUserInfo(db.UserInfo{
		UserId:   uuid.New().String(),
		UserName: u.UserName,
		Email:    u.Email,
	})
}

func Login(u *UserLogin) bool {
	userLoginFromDb := db.ReadUserLogin(u.UserName)
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(userLoginFromDb.HashedPassword))
	return err == nil
}
