package packet

import (
	v649packet "github.com/oomph-ac/mv/multiversion/mv649/packet"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func NewClientPool() packet.Pool {
	pool := v649packet.NewClientPool()

	pool[packet.IDPlayerAuthInput] = func() packet.Packet { return &PlayerAuthInput{} }

	delete(pool, packet.IDSetHud)

	return pool
}

func NewServerPool() packet.Pool {
	pool := v649packet.NewServerPool()

	delete(pool, packet.IDSetHud)

	pool[packet.IDLevelChunk] = func() packet.Packet { return &LevelChunk{} }
	pool[packet.IDPlayerList] = func() packet.Packet { return &PlayerList{} }

	return pool
}
