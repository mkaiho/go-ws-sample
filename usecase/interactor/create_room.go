package interactor

import (
	"context"

	"github.com/mkaiho/go-ws-sample/entity"
	"github.com/mkaiho/go-ws-sample/usecase/port"
)

var _ CreateRoomInteractor = (*createRoomInteractor)(nil)

type (
	CreateRoomInput struct {
		Name        string
		Description *string
	}
	CreateRoomOutput struct {
		Room *entity.Room
	}
	CreateRoomInteractor interface {
		Create(ctx context.Context, input *CreateRoomInput) (*CreateRoomOutput, error)
	}
	createRoomInteractor struct {
		rooms port.RoomsManager
	}
)

func NewCreateRoomInteractor(rooms port.RoomsManager) *createRoomInteractor {
	return &createRoomInteractor{
		rooms: rooms,
	}
}

func (it *createRoomInteractor) Create(ctx context.Context, input *CreateRoomInput) (*CreateRoomOutput, error) {
	out, err := it.rooms.Create(ctx, &port.CreateRoomInput{
		Name:        input.Name,
		Description: input.Description,
	})
	if err != nil {
		return nil, err
	}

	return &CreateRoomOutput{
		Room: out.Room,
	}, nil
}
