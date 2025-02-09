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

	Run: logout,
}

func logout(cmd *cobra.Command, args []string) {
	passed_registry, _ := cmd.Flags().GetString("registry")
	timeout, _ := cmd.Flags().GetInt("timeout")

	if passed_registry == "" {
		_ = cmd.Help()
		return
	}

	credentials := auth.RegistryCredits{
		Registry: passed_registry,
	}

	credentials.LogoutFromRegistry(timeout)
}

func init() {
	LogoutCmd.Flags().StringP("registry", "r", "", "Registry URL")
	LogoutCmd.Flags().IntP("timeout", "t", 10, "Timeout [s]")

}
