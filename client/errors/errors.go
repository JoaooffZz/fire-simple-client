package errors

import "fmt"

type UrlNotAuth struct {
	Msg string
}

func (u UrlNotAuth) Error() string {
	return fmt.Sprintf("Url malformed: %s", u.Msg)
}

type NotAuthClient struct {
	Msg string
}

func (n *NotAuthClient) Error() string {
	return fmt.Sprintf("CLIENT NOT AUTH CONNECTION. Message: %s", n.Msg)
}
