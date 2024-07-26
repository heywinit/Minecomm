package minecomm

import (
	"net"
	"sync"
)

type Client struct {
	con *net.TCPConn
	readChannel sync.Mutex
	writeChannel sync.Mutex
}

func NewClient() *Client {
	return &Client{}
}

// Connect to a server
// addr format must be "ip:port" if port is specified, it uses 25565 
func (mc *Client) Connect(ip string, port string) error {
	//idk why i made it string but yes.
	if(port == "") {
		port = "25565"
	}

	conn, err := net.Dial("tcp", ip + ":" + string(port))
	if err != nil {
		return err
	}

	mc.con = conn.(*net.TCPConn)

	return nil
}


func (c *Client) GetAddr() string {
	return c.con.RemoteAddr().String()
}