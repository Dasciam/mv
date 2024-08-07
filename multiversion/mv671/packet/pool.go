package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol/packet"

func NewClientPool() packet.Pool {
	pool := packet.NewClientPool()

	pool[packet.IDContainerClose] = func() packet.Packet { return &ContainerClose{} }
	pool[packet.IDText] = func() packet.Packet { return &Text{} }
	pool[packet.IDCodeBuilderSource] = func() packet.Packet { return &CodeBuilderSource{} }

	return pool
}

func NewServerPool() packet.Pool {
	pool := packet.NewServerPool()

	pool[packet.IDCraftingData] = func() packet.Packet { return &CraftingData{} }
	pool[packet.IDContainerClose] = func() packet.Packet { return &ContainerClose{} }
	pool[packet.IDText] = func() packet.Packet { return &Text{} }
	pool[packet.IDStartGame] = func() packet.Packet { return &StartGame{} }

	return pool
}
