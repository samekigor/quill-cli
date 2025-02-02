package utils

import (
	"fmt"

	"github.com/zalando/go-keyring"
)

func SaveCredits(service string, registry string, user string, pwd string) error {
	key := fmt.Sprintf("%s:%s", registry, user)
	err := keyring.Set(service, key, pwd)
	if err != nil {
		return err
	}
	return nil
}
