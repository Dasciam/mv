package packet

import (
	v662packet "github.com/oomph-ac/mv/multiversion/mv662/packet"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func NewClientPool() packet.Pool {
	pool := v662packet.NewClientPool()

	pool[packet.IDPlayerAuthInput] = func() packet.Packet { return &PlayerAuthInput{} }
	pool[packet.IDLecternUpdate] = func() packet.Packet { return &LecternUpdate{} }

	return pool
}

func NewServerPool() packet.Pool {
	pool := v662packet.NewServerPool()

	pool[packet.IDMobEffect] = func() packet.Packet { return &MobEffect{} }
	pool[packet.IDResourcePacksInfo] = func() packet.Packet { return &ResourcePacksInfo{} }
	pool[packet.IDSetActorMotion] = func() packet.Packet { return &SetActorMotion{} }

	return pool
}
