package main

import (
	"flag"
	"fmt"
	"gchat/client/ws_client"
)

func main() {
	userId := flag.Int64("user", 0, "User ID")
	deviceId := flag.Int64("device", 0, "Device ID")
	seq := flag.Int64("seq", 0, "Sequence number")
	flag.Parse()
	client := ws_client.WSClient{
		UserId:   *userId,
		DeviceId: *deviceId,
		Seq:      *seq,
	}
	fmt.Printf("UserId: %d, DeviceId: %d, Seq: %d\n", client.UserId, client.DeviceId, client.Seq)
	client.Login()
	client.Start()
	select {}
}
