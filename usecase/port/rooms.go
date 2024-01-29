package port

import (
	"context"

	"github.com/mkaiho/go-ws-sample/entity"
)

type (
	FindRoomsInput  struct{}
	FindRoomsOutput struct {
		Rooms entity.Rooms
	}
	GetRoomInput struct {
		ID entity.ID
	}
	GetRoomOutput struct {
		Room *entity.Room
	}
	RoomsReader interface {
		Find(ctx context.Context, input *FindRoomsInput) (*FindRoomsOutput, error)
		Get(ctx context.Context, input *GetRoomInput) (*GetRoomOutput, error)
	}
)

type (
	CreateRoomInput struct {
		Name        string
		Description *string
	}
	CreateRoomOutput struct {
		Room *entity.Room
	}
	DeleteRoomInput struct {
		ID entity.ID
	}
	DeleteRoomOutput struct{}
	RoomsWriter      interface {
		Create(ctx context.Context, input *CreateRoomInput) (*CreateRoomOutput, error)
		Delete(ctx context.Context, input *DeleteRoomInput) (*DeleteRoomOutput, error)
	}
)

type RoomsManager interface {
	RoomsReader
	RoomsWriter
}
