package middleware

import (
	"strings"

	fserr "github.com/JoaooffZz/fire-simple-client/client/errors"
)

// firesp://user:key@firesimple:port/
func AuthUrl(url string) (string, error) {

	withoutPrefix := strings.TrimPrefix(url, "frsp://")
	if withoutPrefix == "" {
		return "", fserr.UrlNotAuth{
			Msg: "prefix (firesp://) not found",
		}
	}

	parts := strings.SplitN(withoutPrefix, "@", 2)
	address := parts[1]

	domain := strings.SplitN(address, ":", 2)

	if domain[0] != "firesimple" {
		return "", fserr.UrlNotAuth{
			Msg: "prefix (firesimple) not found",
		}
	}

	return domain[1], nil
}
