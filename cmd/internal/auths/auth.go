package auths

import (
	"context"
	"log"

	"github.com/zalando/go-keyring"
	// "golang.org/x/term"
	"github.com/AlecAivazis/survey/v2"
	"oras.land/oras-go/v2/registry/remote"
	orasAuth "oras.land/oras-go/v2/registry/remote/auth"
	// "golang.org/x/term"
	// ec "github.com/samekigor/quill-cli/cmd/internal/exitcodes"
)

var service = "quill-cli"

type RegistryCredits struct {
	RegistryUrl string
	Username    string
	Password    string
	Repository  string
	Tag         string
}

func (rc *RegistryCredits) deleteCredits() {
	// Delete the credentials in the keyring
	err := keyring.Delete(service, rc.Username)
	if err != nil {
		log.Fatalf("Failed to delete credentials: %v", err)
	}
}

func (rc *RegistryCredits) saveCredits() {
	// Store the credentials in the keyring
	err := keyring.Set(service+rc.RegistryUrl, rc.Username, rc.Password)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Credentials saved successfully.")
}

func (rc *RegistryCredits) GetCredentials() (string, string) {
	// Retrieve the credentials from the keyring
	password, err := keyring.Get(service+rc.RegistryUrl, rc.Username)
	if err != nil {
		if err == keyring.ErrNotFound {
			log.Println("Password not found.")
		} else {
			log.Fatalf("Failed to get password: %v", err)
		}
		return "", ""
	}
	log.Println("Credentials retrieved successfully.")
	return rc.Username, password
}

func (rc *RegistryCredits) GetPasswordFromUser() {
	var password string
	prompt := &survey.Password{
		Message: "Enter password:",
	}
	err := survey.AskOne(prompt, &password)
	if err != nil {
		log.Fatalf("Failed to read password: %v", err)
	}
	rc.Password = password
}

func (rc *RegistryCredits) LoginToRegistry() {
	ctx := context.Background()

	remoteRegistry, err := remote.NewRegistry(rc.RegistryUrl)
	if err != nil {
		log.Fatalf("Failed to create remote registry: %v", err)
	}

	remoteRegistry.Client = &orasAuth.Client{
		Credential: func(ctx context.Context, registry string) (orasAuth.Credential, error) {
			return orasAuth.Credential{
				Username: rc.Username,
				Password: rc.Password,
			}, nil
		},
		Cache: orasAuth.NewCache(),
	}

	err = remoteRegistry.Ping(ctx)
	if err != nil {
		log.Fatalf("Failed to ping registry %s: %v", rc.RegistryUrl, err)
	}

	rc.saveCredits()

	log.Println("Successfully logged into registry:", rc.RegistryUrl)
}

func (rc *RegistryCredits) LogoutFromRegistry() {
	rc.deleteCredits()
}
