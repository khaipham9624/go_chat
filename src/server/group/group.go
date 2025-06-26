package group

import (
	"gochat/src/server/db"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var GroupCollection string = "group"

type GroupType int

const (
	Private GroupType = iota
	Protected
	Public
)

type Group struct {
	Id        uuid.UUID
	Name      string
	GroupType GroupType
	Users     []uuid.UUID
}

func (g *Group) WriteToDb() {
	client := db.GetClient()
	if client == nil {
		return
	}
	var bsonUsers []string
	for _, user := range g.Users {
		bsonUsers = append(bsonUsers, user.String())
	}

	db.InsertOne(client, db.DbName, GroupCollection, bson.M{"_id": g.Id.String(), "name": g.Name, "grouptype": g.GroupType, "members": bsonUsers})
}
