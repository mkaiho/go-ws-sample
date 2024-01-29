package routes

import (
	"net/http"

	"github.com/mkaiho/go-ws-sample/controller/web/handlers"
)

func NewRoomsRoutes(
	roomsList *handlers.ListRoomsHandler,
	roomsGet *handlers.GetRoomHandler,
	roomsCreate *handlers.CreateRoomHandler,
	roomsDelete *handlers.DeleteRoomHandler,
) Routes {
	return Routes{
		{
			method:   http.MethodGet,
			path:     "/rooms",
			handlers: handlers.Handlers{roomsList.Handle},
		},
		{
			method:   http.MethodPost,
			path:     "/rooms",
			handlers: handlers.Handlers{roomsCreate.Handle},
		},
		{
			method:   http.MethodGet,
			path:     "/rooms/:room_id",
			handlers: handlers.Handlers{roomsGet.Handle},
		},
		{
			method:   http.MethodDelete,
			path:     "/rooms/:room_id",
			handlers: handlers.Handlers{roomsDelete.Handle},
		},
	}
}
