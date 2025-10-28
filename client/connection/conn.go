package connection

import (
	"net"

	fsmidd "github.com/JoaooffZz/fire-simple-client/client/middleware"
)

type Connection struct {
	conn net.Conn
}

func Default(url string) (*Connection, error) {

	// firesp://user:test@firesimple:port/
	port, err := fsmidd.AuthUrl(url)
	if err != nil {
		return nil, err
	}

	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		return nil, err
	}

	return &Connection{conn: conn}, nil
}
