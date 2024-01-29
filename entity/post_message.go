package entity

import "time"

type PostMessage struct {
	ID             ID
	PostedDatetime *time.Time
	PostedBy       *User
}

type PostMessages []*PostMessage
