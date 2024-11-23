package connect

import (
	"container/list"
	"gchat/pkg/protocol/pb"
	"sync"
)

var RoomsManager sync.Map

func SubscribedRoom(conn *Conn, roomId int64) {
	if roomId == conn.RoomId {
		return
	}

	oldRoomId := conn.RoomId
	// cancel subscription
	if oldRoomId != 0 {
		value, ok := RoomsManager.Load(oldRoomId)
		if !ok {
			return
		}
		room := value.(*Room)
		room.Unsubscribe(conn)

		if room.Conns.Front() == nil {
			RoomsManager.Delete(oldRoomId)
		}
		return
	}

	// subscribe
	if roomId != 0 {
		value, ok := RoomsManager.Load(roomId)
		var room *Room
		if !ok {
			room = NewRoom(roomId)
			RoomsManager.Store(roomId, room)
		} else {
			room = value.(*Room)
		}
		room.Subscribe(conn)
		return
	}
}

// PushRoom push message to room
func PushRoom(roomId int64, message *pb.Message) {
	value, ok := RoomsManager.Load(roomId)
	if !ok {
		return
	}

	value.(*Room).Push(message)
}

type Room struct {
	RoomId int64      // room id
	Conns  *list.List // connections that have subscribed to the room
	lock   sync.RWMutex
}

func NewRoom(roomId int64) *Room {
	return &Room{
		RoomId: roomId,
		Conns:  list.New(),
	}
}

func (r *Room) Subscribe(conn *Conn) {
	r.lock.Lock()
	defer r.lock.Unlock()

	conn.Element = r.Conns.PushBack(conn)
	conn.RoomId = r.RoomId
}

func (r *Room) Unsubscribe(conn *Conn) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.Conns.Remove(conn.Element)
	conn.Element = nil
	conn.RoomId = 0
}

func (r *Room) Push(message *pb.Message) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	element := r.Conns.Front()
	for {
		conn := element.Value.(*Conn)
		conn.Send(pb.PackageType_PT_MESSAGE, 0, message, nil)

		element = element.Next()
		if element == nil {
			break
		}
	}
}
