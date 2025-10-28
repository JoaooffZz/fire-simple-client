package middleware

import (
	"strings"
)

func AuthUri(url, user, password string) error {

	withoutPrefix := strings.TrimPrefix(url, "frsp://")

	parts := strings.SplitN(withoutPrefix, "@", 2)

	userPass := parts[0]

	credURL := strings.SplitN(userPass, ":", 2)

	if credURL[0] != user {
		return &NotAuthClient{
			Msg: "User different default",
		}
	}

	if credURL[1] != password {
		return &NotAuthClient{
			Msg: "Password different default",
		}
	}

	return nil
}
