package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type UserLoginInfo struct {
	UserId         string `bson:"_id" json: "id"`
	UserName       string `bson:"username"`
	HashedPassword string `bson:"hashed_password"`
}

type UserInfo struct {
	UserId   string `bson:"_id" json: "id"`
	UserName string `bson:"username"`
	Email    string `bson:"email"`
}

func CreateUserIndex() {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1, "username": 1},
		Options: options.Index().SetUnique(true),
	}

	collection := client.Database(dbName).Collection("user_info")
	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(dbName).Collection("user_login")
	_, err = collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateUserInfo(userInfo UserInfo) bool {
	collection := "user_info"
	binData, err := bson.Marshal(userInfo)
	if err != nil {
		fmt.Println("failed to marshal")
		return false
	}
	data := bson.M{}
	err = bson.Unmarshal(binData, &data)
	if err != nil {
		fmt.Println("failed to unmarshal")
		return false
	}
	fmt.Println(data)
	InsertOne(client, dbName, collection, data)
	return err == nil
}

func CreateUserLogin(user UserLoginInfo) bool {
	collection := "user_login"
	binData, err := bson.Marshal(user)
	if err != nil {
		fmt.Println("failed to marshal")
		return false
	}

	data := bson.M{}
	err = bson.Unmarshal(binData, &data)
	if err != nil {
		fmt.Println("failed to unmarshal")
		return false
	}
	fmt.Println(data)
	InsertOne(client, dbName, collection, data)
	return err == nil
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
