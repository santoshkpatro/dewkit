package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install dewkit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing dewkit...")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
