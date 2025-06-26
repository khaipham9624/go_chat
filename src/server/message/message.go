package message

import (
	"time"

	"github.com/google/uuid"
)

type MessageType string

const (
	TEXT  MessageType = "TEXT"
	IMG   MessageType = "IMG"
	VIDEO MessageType = "VIDEO"
)

type Message struct {
	Id         int
	GroupId    uuid.UUID
	FromUserId string
	Type       MessageType
	Data       string
	When       time.Time
}
