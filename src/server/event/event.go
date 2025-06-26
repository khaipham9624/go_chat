package event

type EventType string

const (
	Unknown       EventType = "UNKNOWN"
	CreateGroup   EventType = "CREATEGROUP"
	DeleteGroup   EventType = "DELETEGROUP"
	JoinGroup     EventType = "JOINGROUP"
	LeaveGroup    EventType = "LEAVEGROUP"
	SendMessage   EventType = "SENDMESSAGE"
	EditMessage   EventType = "EDITMESSAGE"
	RemoveMessage EventType = "REMOVEMESSAGE"
	ReactMessage  EventType = "REACTMESSAGE"
	AddPolicy     EventType = "ADDPOLICY"
	EditPolicy    EventType = "EDITPOLICY"
	RemovePolicy  EventType = "REMOVEPOLICY"
	AddUser       EventType = "ADDUSER"
	RemoveUser    EventType = "REMOVEUSER"
	GrantAdmin    EventType = "GRANTADMIN"
)

// const (
// 	Unknown string = "UNKNOWN"
// 	CreateGroup
// 	DeleteGroup
// 	JoinGroup
// 	LeaveGroup
// 	SendMessage
// 	EditMessage
// 	RemoveMessage
// 	ReactMessage
// 	AddPolicy
// 	EditPolicy
// 	RemovePolicy
// 	AddUser
// 	RemoveUser
// 	GrantAdmin
// )
