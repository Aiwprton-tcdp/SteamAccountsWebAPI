package services

import (
	"time"

	steam_totp "github.com/fortis/go-steam-totp"
)

type AccountService interface {
	GetGuardCode(shared_secret string) string
}

func GetGuardCode(shared_secret string) string {
	code, err := steam_totp.GenerateAuthCode(shared_secret, time.Now())

	if err != nil {
		code = ""
	}

	return code
}
