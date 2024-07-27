package minecomm

import (
	"github.com/heywinit/minecomm/internal/models"
	"github.com/heywinit/minecomm/internal/models/packets"
	"net"
	"strconv"
	"sync"
)

type Client struct {
	con                  *net.TCPConn
	readChannel          sync.Mutex
	writeChannel         sync.Mutex
	compressionThreshold int32
}

func NewClient() *Client {
	return &Client{}
}

// Connect to a server
// addr format must be "ip:port"
func (mc *Client) Connect(ip string, port uint16, player models.Player, protocolVersion int32) error {
	conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(int(port)))
	if err != nil {
		return err
	}

	mc.con = conn.(*net.TCPConn)

	handshakePacket := packets.HandShakePacket{
		MinecraftPacket: packets.MinecraftPacket{PacketID: 0x00},
		ProtocolVersion: protocolVersion,
		ServerAddress:   ip,
		ServerPort:      port,
		NextState:       2,
	}

	err = mc.WritePacket(&handshakePacket)

	if err != nil {
		return err
	}

	loginStartPacket := packets.LoginStartPacket{
		MinecraftPacket: packets.MinecraftPacket{PacketID: 0x00},
		Name:            player.Name,
		UUID:            player.UUID,
	}

	err = mc.WritePacket(&loginStartPacket)
	if err != nil {
		return err
	}

	return nil
}

func (mc *Client) IsCompressionEnabled() bool {
	return mc.compressionThreshold > 0
}

func (mc *Client) GetAddr() string {
	return mc.con.RemoteAddr().String()
}
