package groupmanager

import (
	"gochat/src/server/group"

	"github.com/google/uuid"
)

type AvailableGroupId int

type GroupManager interface {
	GetGroup(groupId uuid.UUID) *group.Group
	CreateGroup(groupName string, createdBy uuid.UUID, userIds []uuid.UUID) uuid.UUID
	DeleteGroup(groupId uuid.UUID) bool
}
