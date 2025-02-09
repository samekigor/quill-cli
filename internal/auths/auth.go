package auths

import (
	"context"
	"log"
	"time"

	"github.com/samekigor/quill-cli/cmd/client"
	"github.com/samekigor/quill-cli/internal/utils"
	"github.com/samekigor/quill-cli/proto/auth"

	"github.com/AlecAivazis/survey/v2"
)

type RegistryCredits struct {
	Registry   string
	Username   string
	Password   string
	Repository string
	Tag        string
}

func (rc *RegistryCredits) LogoutFromRegistry(timeout int) (msg string, err error) {
	defer func() { rc.Password = "" }()

	ctx, cancel := client.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	logoutStatus, err := client.GrpcClient.Auth.LogoutFromRegistry(ctx, &auth.LogoutRequest{
		Registry: rc.Registry,
		Username: rc.Username,
	})

	if err != nil || !logoutStatus.IsSuccess {
		return "Failure with logout", err
	} else {
		return logoutStatus.Message, err
	}
}

func (rc *RegistryCredits) LoginToRegistry(timeout int) (msg string, err error) {

	defer func() { rc.Password = "" }()
	ctx, cancel := client.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	re := utils.RegistryEntry{
		Registry: rc.Registry,
		User:     rc.Username,
		Password: rc.Password,
	}

	re.EncodePassword()

	if err = re.AddRegistryEntry(); err != nil {
		utils.ErrorLogger.Printf("Failed to add registry entry: %v", err)
		return "Failed to add registry entry.\n", err
	}

	loginStatus, err := client.GrpcClient.Auth.LoginToRegistry(ctx, &auth.LoginRequest{
		Registry: rc.Registry,
		Username: rc.Username,
	})

	utils.InfoLogger.Print(loginStatus)

	if err != nil || !loginStatus.IsSuccess {
		if re.RemoveRegistryEntry(); err != nil {
			utils.ErrorLogger.Printf("Failed to remove registry entry: %v", err)
		}
		return "Failure!", err
	} else {
		return loginStatus.Message, nil
	}

}

func (rc *RegistryCredits) GetPasswordFromUser() {
	var password string
	prompt := &survey.Password{
		Message: "Enter password:",
	}
	err := survey.AskOne(prompt, &password)
	if err != nil {
		log.Printf("Failed to read password: %v", err)
	}
	rc.Password = password
}
