package auths

import (
	"github.com/samekigor/quill-cli/cmd/internal/auths"
	"github.com/spf13/cobra"
)

var LoginCmd = &cobra.Command{
	Use:   `login`,
	Short: `Login to your chosen registry`,
	Long: `Login to your chosen registry using your credentials.

	Example:
	quill auths login --registry <registry> --username <username>`,

	Run: loginToRegistry,
}

func loginToRegistry(cmd *cobra.Command, args []string) {
	passed_registry, _ := cmd.Flags().GetString("registry")
	passed_user, _ := cmd.Flags().GetString("user")

	if passed_user == "" || passed_registry == "" {
		_ = cmd.Help()
		return
	}

	credentials := auths.RegistryCredits{
		RegistryUrl: passed_registry,
		Username:    passed_user,
	}
	credentials.GetPasswordFromUser()
	credentials.LoginToRegistry()
}

func init() {
	LoginCmd.Flags().StringP("registry", "r", "", "Registry URL")
	LoginCmd.Flags().StringP("user", "u", "", "Username")

}
