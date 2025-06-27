package user

import (
	"gochat/src/server/db"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type User struct {
	Id       uuid.UUID
	UserName string
	FullName string
}

var UserCollection string = "user"
var client *mongo.Client

func (u *User) WriteToDb() {
	client = db.GetClient()
	if client == nil {
		return
	}
	db.InsertOne(client, db.DbName, UserCollection, bson.M{"_id": u.Id.String(), "username": u.UserName, "fullname": u.FullName})
}

func CreateUser()
