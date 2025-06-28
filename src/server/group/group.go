package group

import (
	"gochat/src/server/db"

	"github.com/google/uuid"
)

type GroupType string

const (
	Private GroupType = "private"
	Public  GroupType = "public"
)

type Group struct {
	Id        uuid.UUID
	Name      string
	GroupType GroupType
	Users     []uuid.UUID
}

func (g *Group) WriteToDb() {
	db.CreateGroup(db.Group{})
}
