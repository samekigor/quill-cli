package auths

import (
	"github.com/samekigor/quill-cli/internal/auths"
	"github.com/spf13/cobra"
)

var LoginCmd = &cobra.Command{
	Use:   `login`,
	Short: `Login to your chosen registry`,
	Long: `Login to your chosen registry using your credentials.

	Example:
	quill auths login --registry <registry> --username <username> --timeout <integerr>`,

	Run: login,
}

func login(cmd *cobra.Command, args []string) {
	passed_registry, _ := cmd.Flags().GetString("registry")
	passed_user, _ := cmd.Flags().GetString("user")
	timeout, _ := cmd.Flags().GetInt("timeout")

	if passed_user == "" || passed_registry == "" {
		_ = cmd.Help()
		return
	}

	credentials := auths.RegistryCredits{
		Registry: passed_registry,
		Username: passed_user,
	}
	credentials.GetPasswordFromUser()

	if credentials.Password == "" {
		cmd.PrintErrln("Failed to get password")
		return
	}
	msg, err := credentials.LoginToRegistry(timeout)

	if err != nil {
		cmd.PrintErrf("Error: %v, Message from daemon: %v\n", err, msg)
		return
	}
	cmd.Print(msg)
}

func init() {
	LoginCmd.Flags().StringP("registry", "r", "", "Registry URL")
	LoginCmd.Flags().StringP("user", "u", "", "Username")
	LoginCmd.Flags().IntP("timeout", "t", 10, "Timeout [s]")

}
