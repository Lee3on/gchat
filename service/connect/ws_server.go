package connect

import (
	"gchat/pkg/logger"
	"gchat/pkg/util"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 65536,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// StartWSServer start WebSocket server
func StartWSServer(address string) {
	http.HandleFunc("/ws", wsHandler)
	logger.Logger.Info("websocket server start")
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Sugar.Error(err)
		return
	}

	conn := &Conn{
		CoonType: ConnTypeWS,
		WS:       wsConn,
	}
	DoConn(conn)
}

// DoConn process connection
func DoConn(conn *Conn) {
	defer util.RecoverPanic()

	for {
		err := conn.WS.SetReadDeadline(time.Now().Add(12 * time.Minute))
		if err != nil {
			HandleReadErr(conn, err)
			return
		}
		_, data, err := conn.WS.ReadMessage()
		if err != nil {
			HandleReadErr(conn, err)
			return
		}

		conn.HandleMessage(data)
	}
}

// HandleReadErr handle connection errors
func HandleReadErr(conn *Conn, err error) {
	logger.Logger.Debug("read tcp error：", zap.Int64("user_id", conn.UserId),
		zap.Int64("device_id", conn.DeviceId), zap.Error(err))
	str := err.Error()

	// use of closed network connection
	if strings.HasSuffix(str, "use of closed network connection") {
		return
	}

	conn.Close()

	// client closes connection or program exits abnormally
	if err == io.EOF {
		return
	}

	// after SetReadDeadline, the timeout returns an error
	if strings.HasSuffix(str, "i/o timeout") {
		return
	}
}
