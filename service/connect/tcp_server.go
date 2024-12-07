package connect

import (
	"bufio"
	"context"
	"gchat/pkg/logger"
	"gchat/pkg/protocol/pb"
	"gchat/pkg/rpc"
	"net"
	"sync"
	"time"

	"go.uber.org/zap"
)

var (
	listener net.Listener
	wg       sync.WaitGroup
)

// StartTCPServer starts a TCP server.
func StartTCPServer(addr string) {
	var err error
	listener, err = net.Listen("tcp", addr)
	if err != nil {
		logger.Sugar.Error("Error starting server:", err)
		panic(err)
	}
	logger.Sugar.Info("TCP server started on", addr)

	wg.Add(1)
	go acceptConnections()
	wg.Wait()
}

func acceptConnections() {
	defer wg.Done()
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Sugar.Error("Error accepting connection:", err)
			continue
		}
		logger.Logger.Debug("connect:", zap.String("addr", conn.RemoteAddr().String()))
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		logger.Logger.Debug("Connection closed:", zap.String("addr", conn.RemoteAddr().String()))
	}()

	// Initialize connection data
	c := &Conn{
		CoonType: CoonTypeTCP,
		TCP:      conn,
	}
	connCtx := context.WithValue(context.Background(), "conn", c)

	// Read messages from the connection
	reader := bufio.NewReader(conn)
	for {
		// Set a timeout for reading
		conn.SetReadDeadline(time.Now().Add(11 * time.Minute))
		message, err := reader.ReadBytes('\n') // Assuming messages end with newline
		if err != nil {
			logger.Logger.Error("Error reading from connection:", zap.Error(err))
			handleDisconnection(connCtx, c, err)
			return
		}
		c.HandleMessage(message)
	}
}

func handleDisconnection(ctx context.Context, conn *Conn, err error) {
	logger.Logger.Debug("close", zap.String("addr", conn.TCP.RemoteAddr().String()),
		zap.Int64("user_id", conn.UserId), zap.Int64("device_id", conn.DeviceId), zap.Error(err))

	DeleteConn(conn.DeviceId)

	if conn.UserId != 0 {
		_, _ = rpc.GetLogicIntClient().Offline(ctx, &pb.OfflineReq{
			UserId:     conn.UserId,
			DeviceId:   conn.DeviceId,
			ClientAddr: conn.TCP.RemoteAddr().String(),
		})
	}
}
