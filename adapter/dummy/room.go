package dummy

import (
	"context"
	"fmt"
	"slices"
	"sync"

	"github.com/mkaiho/go-ws-sample/entity"
	"github.com/mkaiho/go-ws-sample/usecase"
	"github.com/mkaiho/go-ws-sample/usecase/port"
)

var (
	_ port.RoomsReader = (*RoomsAccess)(nil)
	_ port.RoomsWriter = (*RoomsAccess)(nil)
)

type RoomsAccess struct {
	mux         sync.RWMutex
	idGenerator port.IDGenerator
	rooms       entity.Rooms
}

func NewRoomsAccess(idGenerator port.IDGenerator) *RoomsAccess {
	return &RoomsAccess{
		idGenerator: idGenerator,
	}
}

func (a *RoomsAccess) Find(ctx context.Context, input *port.FindRoomsInput) (*port.FindRoomsOutput, error) {
	a.mux.RLock()
	defer a.mux.RUnlock()
	return &port.FindRoomsOutput{
		Rooms: append(entity.Rooms{}, a.rooms...),
	}, nil
}

func (a *RoomsAccess) Get(ctx context.Context, input *port.GetRoomInput) (*port.GetRoomOutput, error) {
	a.mux.RLock()
	defer a.mux.RUnlock()
	for _, room := range a.rooms {
		if room.ID == input.ID {
			return &port.GetRoomOutput{
				Room: room,
			}, nil
		}
	}
	return nil, usecase.ErrNotFoundEntity
}

func (a *RoomsAccess) Create(ctx context.Context, input *port.CreateRoomInput) (*port.CreateRoomOutput, error) {
	a.mux.Lock()
	defer a.mux.Unlock()

	id, err := a.idGenerator.Generate(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to generate id: %w", err)
	}
	room := entity.Room{
		ID:          id,
		Name:        input.Name,
		Description: input.Description,
	}
	a.rooms = append(a.rooms, &room)

	return &port.CreateRoomOutput{
		Room: &room,
	}, nil
}

func (a *RoomsAccess) Delete(ctx context.Context, input *port.DeleteRoomInput) (*port.DeleteRoomOutput, error) {
	a.mux.Lock()
	defer a.mux.Unlock()

	a.rooms = slices.DeleteFunc(a.rooms, func(r *entity.Room) bool {
		return r.ID == input.ID
	})

	return &port.DeleteRoomOutput{}, nil
}
