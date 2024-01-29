package interactor

import (
	"context"

	"github.com/mkaiho/go-ws-sample/entity"
	"github.com/mkaiho/go-ws-sample/usecase/port"
)

var _ ListRoomsInteractor = (*listRoomsInteractor)(nil)

type (
	ListRoomsInput struct {
	}
	ListRoomsOutput struct {
		Rooms entity.Rooms
	}
	ListRoomsInteractor interface {
		List(ctx context.Context, input *ListRoomsInput) (*ListRoomsOutput, error)
	}
	listRoomsInteractor struct {
		rooms port.RoomsReader
	}
)

func NewListRoomsInteractor(rooms port.RoomsReader) *listRoomsInteractor {
	return &listRoomsInteractor{
		rooms: rooms,
	}
}

func (it *listRoomsInteractor) List(ctx context.Context, input *ListRoomsInput) (*ListRoomsOutput, error) {
	out, err := it.rooms.Find(ctx, &port.FindRoomsInput{})
	if err != nil {
		return nil, err
	}

	return &ListRoomsOutput{
		Rooms: out.Rooms,
	}, nil
}
