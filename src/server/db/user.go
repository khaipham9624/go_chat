package db

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserLoginInfo struct {
	UserName       string `bson: "username"`
	HashedPassword string `bson: "hashed_password"`
}

type UserInfo struct {
	UserId   string `bson: "_id" json: "id"`
	UserName string `bson: "username"`
	Email    string `bson: "email"`
}

func CreateUserInfo(userInfo UserInfo) bool {
	collection := "user_info"
	data := bson.M{}
	InsertOne(client, dbName, collection, data)
	return true
}

func CreateUserLogin(user UserLoginInfo) bool {
	collection := "user_login"
	data := bson.M{}
	InsertOne(client, dbName, collection, data)
	return true
}

func ReadUserLogin(username string) UserLoginInfo {
	collection := "user_login"
	filter := bson.M{}
	result := FindOne(client, dbName, collection, filter)
	userLoginInfo := UserLoginInfo{}
	err := result.Decode(&userLoginInfo)
	if err != nil {
		fmt.Println("error decode")
	}
	return userLoginInfo
}
