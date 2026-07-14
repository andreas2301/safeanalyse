package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/user/safeanalyze/pkg/install"
)

var installCmd = &cobra.Command{
	Use:   "install [tool|model|--all]",
	Short: "Install external scanners or the ML model",
	Long: `Install external security scanners from source or download the
protectai/deberta-v3-base-prompt-injection ONNX model.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		all, _ := cmd.Flags().GetBool("all")
		if all {
			if err := install.InstallAll(); err != nil {
				return fmt.Errorf("installing all scanners: %w", err)
			}
			if err := install.InstallModel(); err != nil {
				return fmt.Errorf("installing model: %w", err)
			}
			fmt.Println("All scanners and model installed.")
			return nil
		}

		if len(args) == 0 {
			return fmt.Errorf("specify a tool name, 'model', or --all")
		}
		target := args[0]
		if target == "model" {
			if err := install.InstallModel(); err != nil {
				return fmt.Errorf("installing model: %w", err)
			}
			fmt.Println("Model installed.")
			return nil
		}

		if err := install.Install(target); err != nil {
			return fmt.Errorf("installing %s: %w", target, err)
		}
		fmt.Printf("%s installed.\n", target)
		return nil
	},
}

func init() {
	installCmd.Flags().Bool("all", false, "install all known scanners and the model")
	rootCmd.AddCommand(installCmd)
}
