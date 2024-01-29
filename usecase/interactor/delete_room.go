package interactor

import (
	"context"

	"github.com/mkaiho/go-ws-sample/entity"
	"github.com/mkaiho/go-ws-sample/usecase/port"
)

var _ DeleteRoomInteractor = (*deleteRoomInteractor)(nil)

type (
	DeleteRoomInput struct {
		ID entity.ID
	}
	DeleteRoomOutput struct {
		Room *entity.Room
	}
	DeleteRoomInteractor interface {
		Delete(ctx context.Context, input *DeleteRoomInput) (*DeleteRoomOutput, error)
	}
	deleteRoomInteractor struct {
		rooms port.RoomsManager
	}
)

func NewDeleteRoomInteractor(rooms port.RoomsManager) *deleteRoomInteractor {
	return &deleteRoomInteractor{
		rooms: rooms,
	}
}

func (it *deleteRoomInteractor) Delete(ctx context.Context, input *DeleteRoomInput) (*DeleteRoomOutput, error) {
	_, err := it.rooms.Delete(ctx, &port.DeleteRoomInput{
		ID: input.ID,
	})
	if err != nil {
		return nil, err
	}

	return &DeleteRoomOutput{}, nil
}
