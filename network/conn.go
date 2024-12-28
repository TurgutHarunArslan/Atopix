package network

import "net"

type Conn struct {
	Id   string
	Conn net.Conn
}
