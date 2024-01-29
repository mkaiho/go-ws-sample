package interactor

import (
	"context"

	"github.com/mkaiho/go-ws-sample/entity"
	"github.com/mkaiho/go-ws-sample/usecase/port"
)

var _ GetRoomInteractor = (*getRoomsInteractor)(nil)

type (
	GetRoomInput struct {
		ID entity.ID
	}
	GetRoomOutput struct {
		Room *entity.Room
	}
	GetRoomInteractor interface {
		Get(ctx context.Context, input *GetRoomInput) (*GetRoomOutput, error)
	}
	getRoomsInteractor struct {
		rooms port.RoomsReader
	}
)

func NewGetRoomInteractor(rooms port.RoomsReader) *getRoomsInteractor {
	return &getRoomsInteractor{
		rooms: rooms,
	}
}

func (it *getRoomsInteractor) Get(ctx context.Context, input *GetRoomInput) (*GetRoomOutput, error) {
	out, err := it.rooms.Get(ctx, &port.GetRoomInput{
		ID: input.ID,
	})
	if err != nil {
		return nil, err
	}

	return &GetRoomOutput{
		Room: out.Room,
	}, nil
}
