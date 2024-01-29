package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkaiho/go-ws-sample/entity"
	"github.com/mkaiho/go-ws-sample/usecase"
	"github.com/mkaiho/go-ws-sample/usecase/interactor"
)

type RoomResponseDetail struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// List
type (
	ListRoomsRequest struct {
	}
	ListRoomsResponse struct {
		Rooms []*RoomResponseDetail `json:"rooms"`
	}
	ListRoomsHandler struct {
		rooms interactor.ListRoomsInteractor
	}
)

func NewListRoomsHandler(rooms interactor.ListRoomsInteractor) *ListRoomsHandler {
	return &ListRoomsHandler{
		rooms: rooms,
	}
}

func (h *ListRoomsHandler) Handle(gc *gin.Context) {
	ctx := gc.Request.Context()
	var req ListRoomsRequest
	if err := ShouldBind(gc, &req); err != nil {
		gc.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	out, err := h.rooms.List(ctx, &interactor.ListRoomsInput{})
	if err != nil {
		gErr := gc.Error(err)
		if errors.Is(err, usecase.ErrNotFoundEntity) {
			gErr.SetType(gin.ErrorTypePublic)
		}
		return
	}

	var res ListRoomsResponse
	for _, room := range out.Rooms {
		res.Rooms = append(res.Rooms, &RoomResponseDetail{
			ID:          room.ID.String(),
			Name:        room.Name,
			Description: room.Description,
		})
	}
	gc.JSON(http.StatusOK, res)
}

// Get
type (
	GetRoomRequest struct {
		ID string `json:"id" uri:"room_id" validate:"required,max=26"`
	}
	GetRoomResponse struct {
		Room *RoomResponseDetail `json:"room"`
	}
	GetRoomHandler struct {
		rooms interactor.GetRoomInteractor
	}
)

func NewGetRoomHandler(rooms interactor.GetRoomInteractor) *GetRoomHandler {
	return &GetRoomHandler{
		rooms: rooms,
	}
}

func (h *GetRoomHandler) Handle(gc *gin.Context) {
	ctx := gc.Request.Context()
	var req GetRoomRequest
	if err := ShouldBind(gc, &req); err != nil {
		gc.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	out, err := h.rooms.Get(ctx, &interactor.GetRoomInput{
		ID: entity.ID(req.ID),
	})
	if err != nil {
		gErr := gc.Error(err)
		if errors.Is(err, usecase.ErrNotFoundEntity) {
			gErr.SetType(gin.ErrorTypePublic)
		}
		return
	}

	res := GetRoomResponse{
		Room: &RoomResponseDetail{
			ID:          out.Room.ID.String(),
			Name:        out.Room.Name,
			Description: out.Room.Description,
		},
	}
	gc.JSON(http.StatusOK, res)
}

// Create
type (
	CreateRoomRequest struct {
		Name        string  `json:"name" validate:"required,max=20"`
		Description *string `json:"description,omitempty" validate:"omitempty,min=1,max=64"`
	}
	CreateRoomResponse struct {
		Room *RoomResponseDetail `json:"room"`
	}
	CreateRoomHandler struct {
		rooms interactor.CreateRoomInteractor
	}
)

func NewCreateRoomHandler(rooms interactor.CreateRoomInteractor) *CreateRoomHandler {
	return &CreateRoomHandler{
		rooms: rooms,
	}
}

func (h *CreateRoomHandler) Handle(gc *gin.Context) {
	ctx := gc.Request.Context()
	var req CreateRoomRequest
	if err := ShouldBind(gc, &req); err != nil {
		gc.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	out, err := h.rooms.Create(ctx, &interactor.CreateRoomInput{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		gErr := gc.Error(err)
		if errors.Is(err, usecase.ErrNotFoundEntity) || errors.Is(err, usecase.ErrAlreadyExistsEntity) {
			gErr.SetType(gin.ErrorTypePublic)
		}
		return
	}

	res := CreateRoomResponse{
		Room: &RoomResponseDetail{
			ID:          out.Room.ID.String(),
			Name:        out.Room.Name,
			Description: out.Room.Description,
		},
	}
	gc.JSON(http.StatusCreated, res)
}

// Delete
type (
	DeleteRoomRequest struct {
		ID string `json:"id" uri:"room_id" validate:"required,max=26"`
	}
	DeleteRoomResponse struct{}
	DeleteRoomHandler  struct {
		rooms interactor.DeleteRoomInteractor
	}
)

func NewDeleteRoomHandler(rooms interactor.DeleteRoomInteractor) *DeleteRoomHandler {
	return &DeleteRoomHandler{
		rooms: rooms,
	}
}

func (h *DeleteRoomHandler) Handle(gc *gin.Context) {
	ctx := gc.Request.Context()
	var req DeleteRoomRequest
	if err := ShouldBind(gc, &req); err != nil {
		gc.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	_, err := h.rooms.Delete(ctx, &interactor.DeleteRoomInput{
		ID: entity.ID(req.ID),
	})
	if err != nil {
		gErr := gc.Error(err)
		if errors.Is(err, usecase.ErrNotFoundEntity) {
			gErr.SetType(gin.ErrorTypePublic)
		}
		return
	}

	gc.Status(http.StatusNoContent)
}
