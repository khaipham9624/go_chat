package inmemorygroupmanager

import (
	"gochat/src/server/group"

	"github.com/google/uuid"
)

type InMemoryGroupManager struct {
	groups map[uuid.UUID]*group.Group
}

func NewInMemoryGroupManager() *InMemoryGroupManager {
	return &InMemoryGroupManager{
		groups: make(map[uuid.UUID]*group.Group),
	}
}

func (gm *InMemoryGroupManager) GetGroup(groupId uuid.UUID) *group.Group {
	return gm.groups[groupId]
}

func (gm *InMemoryGroupManager) CreateGroup(groupName string, createdBy uuid.UUID, userIds []uuid.UUID) uuid.UUID {
	users := []uuid.UUID{}
	users = append(users, createdBy)
	users = append(users, userIds...)
	groupId := uuid.New()
	newGroup := &group.Group{
		Id:        groupId,
		Name:      groupName,
		GroupType: group.Public,
		Users:     users,
	}
	newGroup.WriteToDb()
	gm.groups[groupId] = newGroup
	return groupId
}

func (gm *InMemoryGroupManager) DeleteGroup(groupId uuid.UUID) bool {
	delete(gm.groups, groupId)
	return true
}
