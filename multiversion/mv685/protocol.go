package mv685

import (
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	gtpacket "github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type Protocol struct{}

func (Protocol) Encryption(key [32]byte) gtpacket.Encryption {
	return gtpacket.NewCTREncryption(key[:])
}

func (Protocol) ID() int32 {
	return 685
}

func (Protocol) Ver() string {
	return "1.21.0"
}

func (Protocol) NewReader(r minecraft.ByteReader, shieldID int32, enableLimits bool) protocol.IO {
	return protocol.NewReader(r, shieldID, enableLimits)
}

func (Protocol) NewWriter(r minecraft.ByteWriter, shieldID int32) protocol.IO {
	return protocol.NewWriter(r, shieldID)
}

func (Protocol) Packets(_ bool) gtpacket.Pool {
	return gtpacket.NewServerPool()
}

func (Protocol) ConvertToLatest(pk gtpacket.Packet, _ *minecraft.Conn) []gtpacket.Packet {
	return []gtpacket.Packet{pk}
}

func (Protocol) ConvertFromLatest(pk gtpacket.Packet, _ *minecraft.Conn) []gtpacket.Packet {
	if _, ok := pk.(*gtpacket.ClientBoundCloseForm); ok {
		return nil
	}
	return []gtpacket.Packet{pk}
}
