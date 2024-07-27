package minecomm

import (
	"net"
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
// addr format must be "ip:port" if port is specified, it uses 25565
func (mc *Client) Connect(ip string, port string) error {
	//IDK why I made it string but yes.
	if port == "" {
		port = "25565"
	}

	conn, err := net.Dial("tcp", ip+":"+string(port))
	if err != nil {
		return err
	}

	mc.con = conn.(*net.TCPConn)

	return nil
}

func (mc *Client) IsCompressionEnabled() bool {
	return mc.compressionThreshold > 0
}

func (mc *Client) GetAddr() string {
	return mc.con.RemoteAddr().String()
}
