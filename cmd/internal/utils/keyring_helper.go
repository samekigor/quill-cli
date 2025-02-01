package utils

import (
	"github.com/zalando/go-keyring"
)

func SaveCredits(service string, usr string, pwd string) error {
	err := keyring.Set(service, usr, pwd)
	if err != nil {
		return err
	}
	return nil
}
