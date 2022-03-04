package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sly",
	Short: "Sly enables you to store secrets in your Git repo, safely",
	Long: `Sly make it easy for you to Gnu Privacy Guard (GPG) 
encrypt specific files in a repo so they are "encrypted at 
rest" in your repository`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// add the 'version' or 'v' command
	rootCmd.AddCommand(versionCmd)

	// add the 'init' or 'i' command
	rootCmd.AddCommand(initCmd)

	// add the 'encrypt' or 'e' command
	rootCmd.AddCommand(encryptCmd)
}
