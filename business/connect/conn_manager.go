package connect

import (
	"gchat/pkg/protocol/pb"
	"sync"
)

var ConnsManager = sync.Map{}

func SetConn(deviceId int64, conn *Conn) {
	ConnsManager.Store(deviceId, conn)
}

func GetConn(deviceId int64) *Conn {
	value, ok := ConnsManager.Load(deviceId)
	if ok {
		return value.(*Conn)
	}
	return nil
}

func DeleteConn(deviceId int64) {
	ConnsManager.Delete(deviceId)
}

func PushAll(message *pb.Message) {
	ConnsManager.Range(func(key, value interface{}) bool {
		conn := value.(*Conn)
		conn.Send(pb.PackageType_PT_MESSAGE, 0, message, nil)
		return true
	})
}
