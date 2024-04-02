package cmd
import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "cobra",
		Short: "IPtracter CLI App,",
		Long: "IPtracker CLI App",
	}
)
func Execute() error {
	return rootCmd.Execute()
}