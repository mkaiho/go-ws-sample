package entity

type Room struct {
	ID          ID
	Name        string
	Description *string
	Messages    PostMessages
	Users       Users
}

type Rooms []*Room
