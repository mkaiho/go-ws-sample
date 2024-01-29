package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mkaiho/go-ws-sample/adapter/dummy"
	idAdapter "github.com/mkaiho/go-ws-sample/adapter/id"
	"github.com/mkaiho/go-ws-sample/controller/web"
	"github.com/mkaiho/go-ws-sample/controller/web/handlers"
	"github.com/mkaiho/go-ws-sample/controller/web/routes"
	"github.com/mkaiho/go-ws-sample/usecase/interactor"
	"github.com/mkaiho/go-ws-sample/usecase/port"
	"github.com/mkaiho/go-ws-sample/util"
	"github.com/spf13/cobra"
)

var (
	initErr error
	command *cobra.Command
)

func init() {
	util.InitGLogger(
		util.OptionLoggerLevel(util.LoggerLevelDebug),
		util.OptionLoggerFormat(util.LoggerFormatJSON),
	)
	command = newCommand()
}

func main() {
	var err error
	logger := util.GLogger()
	defer func() {
		if p := recover(); p != nil {
			msg := "panic has occured"
			if pErr, ok := p.(error); ok {
				logger.Error(pErr, msg)
			} else {
				logger.Error(fmt.Errorf("%v", p), msg)
			}
			os.Exit(1)
		}
		if err != nil {
			logger.Error(err, "error has occured")
			os.Exit(1)
		}
		logger.Info("completed")
	}()
	if err = command.Execute(); err != nil {
		return
	}
}

func newCommand() *cobra.Command {
	command := cobra.Command{
		Use:           "echo-server args...",
		Short:         "launch echo-server",
		Long:          "launch echo-server.",
		RunE:          handle,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	command.Flags().IntP("port", "", 3000, "listening port")
	command.Flags().StringP("host", "", "", "host name")

	return &command
}

func handle(cmd *cobra.Command, args []string) (err error) {
	var (
		host string
		port int
	)
	ctx := util.NewContextWithLogger(context.Background(), util.GLogger())
	logger := util.FromContext(ctx)
	if initErr != nil {
		return initErr
	}

	host, err = cmd.Flags().GetString("host")
	if err != nil {
		return err
	}
	port, err = cmd.Flags().GetInt("port")
	if err != nil {
		return err
	}

	server, err := server()
	if err != nil {
		return err
	}

	logger.
		WithValues("host", host).
		WithValues("port", port).
		Info("launch server")
	return server.Run(fmt.Sprintf("%s:%d", "", port))
}

func server() (*web.Server, error) {
	// ports
	var (
		ulidGenerator port.IDGenerator
		roomsManager  port.RoomsManager
	)
	{
		ulidGenerator = idAdapter.NewULIDGenerator()
		roomsManager = dummy.NewRoomsAccess(ulidGenerator)
	}

	// interactors
	var (
		listRoomsInteractor  interactor.ListRoomsInteractor
		getRoomInteractor    interactor.GetRoomInteractor
		createRoomInteractor interactor.CreateRoomInteractor
		deleteRoomInteractor interactor.DeleteRoomInteractor
	)
	{
		listRoomsInteractor = interactor.NewListRoomsInteractor(roomsManager)
		getRoomInteractor = interactor.NewGetRoomInteractor(roomsManager)
		createRoomInteractor = interactor.NewCreateRoomInteractor(roomsManager)
		deleteRoomInteractor = interactor.NewDeleteRoomInteractor(roomsManager)
	}

	// routes
	var r routes.Routes
	health := routes.NewHealthRoutes(
		handlers.NewHealthGetHandler(),
	)
	r = append(r, health...)
	rooms := routes.NewRoomsRoutes(
		handlers.NewListRoomsHandler(listRoomsInteractor),
		handlers.NewGetRoomHandler(getRoomInteractor),
		handlers.NewCreateRoomHandler(createRoomInteractor),
		handlers.NewDeleteRoomHandler(deleteRoomInteractor),
	)
	r = append(r, rooms...)

	return web.NewGinServer(r...), nil
}
