package auths

import (
	auth "github.com/samekigor/quill-cli/internal/auths"
	"github.com/spf13/cobra"
)

var LogoutCmd = &cobra.Command{
	Use:   `logout`,
	Short: `Logout from your chosen registry`,
	Long: `Logout from your chosen registry.

	Example:
	quill auths logout --registry <registry> --username <username>`,

	Run: logoutFromRegistry,
}

func logoutFromRegistry(cmd *cobra.Command, args []string) {
	passed_registry, _ := cmd.Flags().GetString("registry")
	passed_user, _ := cmd.Flags().GetString("user")
	timeout, _ := cmd.Flags().GetInt("timeout")

	if passed_user == "" {
		_ = cmd.Help()
		return
	}

	credentials := auth.RegistryCredits{
		RegistryUrl: passed_registry,
		Username:    passed_user,
	}

	credentials.LogoutFromRegistry(timeout)
}

func init() {
	LogoutCmd.Flags().StringP("registry", "r", "", "Registry URL")
	LogoutCmd.Flags().StringP("user", "u", "", "Username")
	LogoutCmd.Flags().IntP("timeout", "t", 1, "Timeout [s]")

}
