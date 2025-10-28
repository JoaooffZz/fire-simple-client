package connection

import "net"

type Connection struct {
}

func Default(url string) {
	// amqp://user:test@rabbitmq:5672/
	// firesp://user:test@firesimple:port/
	conn, err := net.Dial("tcp", "127.0.0.1:5000")

}
