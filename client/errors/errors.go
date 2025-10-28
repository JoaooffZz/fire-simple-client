package errors

import "fmt"

type NotAuthClient struct {
	Msg string
}

func (n *NotAuthClient) Error() string {
	return fmt.Sprintf("CLIENT NOT AUTH CONNECTION. Message: %s", n.Msg)
}
