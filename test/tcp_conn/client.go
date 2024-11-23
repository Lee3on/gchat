package main

import (
	"fmt"
	"gchat/pkg/protocol/pb"
	"gchat/pkg/util"
	"log"
	"net"
	"sync"
	"time"

	jsoniter "github.com/json-iterator/go"
	"google.golang.org/protobuf/proto"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	client := TcpClient{}
	log.Println("Input UserId, DeviceId, SyncSeq")
	fmt.Scanf("%d %d %d", &client.UserId, &client.DeviceId, &client.Seq)
	client.Start()
	select {}
}

func Json(i interface{}) string {
	bytes, _ := jsoniter.Marshal(i)
	return string(bytes)
}

type TcpClient struct {
	UserId     int64
	DeviceId   int64
	Seq        int64
	Conn       net.Conn
	WriteMutex sync.Mutex // Mutex for writing to the connection
}

func (c *TcpClient) Output(pt pb.PackageType, requestId int64, message proto.Message) {
	var input = pb.Input{
		Type:      pt,
		RequestId: requestId,
	}

	if message != nil {
		bytes, err := proto.Marshal(message)
		if err != nil {
			log.Println(err)
			return
		}
		input.Data = bytes
	}

	inputBytes, err := proto.Marshal(&input)
	if err != nil {
		log.Println(err)
		return
	}

	c.Write(inputBytes)
}

func (c *TcpClient) Write(data []byte) {
	c.WriteMutex.Lock()
	defer c.WriteMutex.Unlock()

	length := len(data)
	header := make([]byte, 4) // Fixed 4-byte header for message length
	header[0] = byte(length >> 24)
	header[1] = byte(length >> 16)
	header[2] = byte(length >> 8)
	header[3] = byte(length)

	// Write the header and the message data
	_, err := c.Conn.Write(append(header, data...))
	if err != nil {
		log.Println("Error writing to connection:", err)
	}
}

func (c *TcpClient) Start() {
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		log.Println(err)
		return
	}

	c.Conn = conn

	c.SignIn()
	c.SyncTrigger()
	c.SubscribeRoom()
	go c.Heartbeat()
	go c.Receive()
}

func (c *TcpClient) SignIn() {
	signIn := pb.SignInInput{
		UserId:   c.UserId,
		DeviceId: c.DeviceId,
		Token:    "0",
	}
	c.Output(pb.PackageType_PT_SIGN_IN, time.Now().UnixNano(), &signIn)
}

func (c *TcpClient) SyncTrigger() {
	c.Output(pb.PackageType_PT_SYNC, time.Now().UnixNano(), &pb.SyncInput{Seq: c.Seq})
	log.Println("Start syncing")
}

func (c *TcpClient) Heartbeat() {
	ticker := time.NewTicker(time.Minute * 5)
	for range ticker.C {
		c.Output(pb.PackageType_PT_HEARTBEAT, time.Now().UnixNano(), nil)
	}
}

func (c *TcpClient) Receive() {
	buffer := make([]byte, 65536) // Buffer for reading messages

	for {
		// Read the header to get the message length
		_, err := c.Conn.Read(buffer[:4])
		if err != nil {
			log.Println("Error reading header:", err)
			return
		}
		length := int(buffer[0])<<24 | int(buffer[1])<<16 | int(buffer[2])<<8 | int(buffer[3])

		// Read the full message based on the length
		_, err = c.Conn.Read(buffer[:length])
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		c.HandlePackage(buffer[:length])
	}
}

func (c *TcpClient) SubscribeRoom() {
	c.Output(pb.PackageType_PT_SUBSCRIBE_ROOM, 0, &pb.SubscribeRoomInput{
		RoomId: 1,
		Seq:    0,
	})
}

func (c *TcpClient) HandlePackage(bytes []byte) {
	var output pb.Output
	err := proto.Unmarshal(bytes, &output)
	if err != nil {
		log.Println(err)
		return
	}

	switch output.Type {
	case pb.PackageType_PT_SIGN_IN:
		log.Println(Json(&output))
	case pb.PackageType_PT_HEARTBEAT:
		log.Println("Heartbeat response")
	case pb.PackageType_PT_SYNC:
		log.Println("Start offline message sync------")
		syncResp := pb.SyncOutput{}
		err := proto.Unmarshal(output.Data, &syncResp)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Offline message sync response: code", output.Code, "message:", output.Message)
		for _, msg := range syncResp.Messages {
			log.Println(util.MessageToString(msg))
			c.Seq = msg.Seq
		}

		ack := pb.MessageACK{
			DeviceAck:   c.Seq,
			ReceiveTime: util.UnixMilliTime(time.Now()),
		}
		c.Output(pb.PackageType_PT_MESSAGE, output.RequestId, &ack)
		log.Println("Offline message sync complete------")
	case pb.PackageType_PT_MESSAGE:
		msg := pb.Message{}
		err := proto.Unmarshal(output.Data, &msg)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(util.MessageToString(&msg))
		c.Seq = msg.Seq
		ack := pb.MessageACK{
			DeviceAck:   msg.Seq,
			ReceiveTime: util.UnixMilliTime(time.Now()),
		}
		c.Output(pb.PackageType_PT_MESSAGE, output.RequestId, &ack)
	default:
		log.Println("Unhandled package type", output, len(bytes))
	}
}
