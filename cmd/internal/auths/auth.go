package auths

import (
	"context"
	"log"
	"time"

	"github.com/samekigor/quill-cli/cmd/client"
	"github.com/samekigor/quill-cli/cmd/internal/utils"
	"github.com/samekigor/quill-cli/proto/auth"

	"github.com/AlecAivazis/survey/v2"
)

var service = "quill-cli"

type RegistryCredits struct {
	RegistryUrl string
	Username    string
	Password    string
	Repository  string
	Tag         string
}

// func (rc *RegistryCredits) GetCredentials() (string, string) {
// 	// Retrieve the credentials from the keyring
// 	password, err := keyring.Get(service+rc.RegistryUrl, rc.Username)
// 	if err != nil {
// 		if err == keyring.ErrNotFound {
// 			log.Println("Password not found.")
// 		} else {
// 			log.Fatalf("Failed to get password: %v", err)
// 		}
// 		return "", ""
// 	}
// 	log.Println("Credentials retrieved successfully.")
// 	return rc.Username, password
// }

func (rc *RegistryCredits) LogoutFromRegistry(timeout int) (msg string, err error) {
	ctx, cancel := client.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	logoutStatus, err := client.GrpcClient.Auth.LogoutFromRegistry(ctx, &auth.LogoutRequest{
		Registry: rc.RegistryUrl,
		Username: rc.Username,
	})
	if err != nil || !logoutStatus.IsSuccess {
		return "Failure with logout", err
	} else {
		log.Print("Removed credentials in keyring.")
		return logoutStatus.Message, err
	}
}

func (rc *RegistryCredits) LoginToRegistry(timeout int) (msg string, err error) {
	defer func() { rc.Password = "" }()
	ctx, cancel := client.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	err = utils.SaveCredits(service, rc.Username, rc.Password)
	if err != nil {
		return "Failure with saving password in keyring.", err
	}
	loginStatus, err := client.GrpcClient.Auth.LoginToRegistry(ctx, &auth.LoginRequest{
		Registry: rc.RegistryUrl,
		Username: rc.Username,
	})
	if err != nil || loginStatus.IsSuccess {
		return "Failed to send credentials to daemon.", err
	} else {
		return loginStatus.Message, err
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

// func (rc *RegistryCredits) LoginToRegistry() error {
// 	ctx := context.Background()

// 	remoteRegistry, err := remote.NewRegistry(rc.RegistryUrl)
// 	if err != nil {
// 		log.Fatalf("Failed to create remote registry: %v", err)
// 	}

// 	remoteRegistry.Client = &orasAuth.Client{
// 		Credential: func(ctx context.Context, registry string) (orasAuth.Credential, error) {
// 			return orasAuth.Credential{
// 				Username: rc.Username,
// 				Password: rc.Password,
// 			}, nil
// 		},
// 		Cache: orasAuth.NewCache(),
// 	}

// 	err = remoteRegistry.Ping(ctx)
// 	if err != nil {
// 		log.Fatalf("Failed to ping registry %s: %v", rc.RegistryUrl, err)
// 	}
// 	defer func() { rc.Password = "" }()
// 	rc.saveCredits()

// 	log.Println("Successfully logged into registry:", rc.RegistryUrl)
// }
