package connection

import "encoding/json"

func (c *Connection) AuthClientIP(ip string) bool {
	msg := MessageData{
		typeEvent:   "ip.guarde",
		typeService: "auth.client",
		data:        []byte(ip),
	}

	msgJson, _ := json.Marshal(msg)

	c.conn.Write(msgJson)

	reply := make([]byte, 1024)
	c.conn.Read(reply)
	if string(reply) == "0" {
		return false
	}
	return true
}
