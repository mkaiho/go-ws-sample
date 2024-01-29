package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
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
		Use:           "echo-client args...",
		Short:         "launch echo-client",
		Long:          "launch echo-client.",
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

	logger.
		WithValues("host", host).
		WithValues("port", port).
		Info("launch server")
	return exec(ctx)
}

func exec(ctx context.Context) error {
	logger := util.FromContext(ctx)
	// WebSocketサーバのURL
	u := url.URL{Scheme: "ws", Host: "localhost:3000", Path: "/rooms/1234/messages"}

	// WebSocketサーバに接続
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	// サーバからのメッセージを受信するゴルーチン
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				logger.Error(err, "failed to read message")
				return
			}
			logger.Info("recieved", "message", string(message))
		}
	}()

	// サーバにメッセージを送信
	err = c.WriteMessage(websocket.TextMessage, []byte("hello"))
	if err != nil {
		return err
	}

	// サーバからの応答を待つ
	time.Sleep(time.Second * 1)

	// WebSocket接続を閉じる
	err = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return err
	}
	select {
	case <-done:
	case <-time.After(time.Second):
	}

	return nil
}
