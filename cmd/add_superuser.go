package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var addSuperUserCmd = &cobra.Command{
	Use:   "add_superuser",
	Short: "Add superuser",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Full Name: ")
		fullName, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		fullName = strings.TrimSpace(fullName)

		fmt.Print("Email: ")
		email, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		email = strings.TrimSpace(email)

		fmt.Print("Password: ")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		if err != nil {
			return err
		}
		password := strings.TrimSpace(string(bytePassword))

		fmt.Println("\nUser details:")
		fmt.Println("Full Name:", fullName)
		fmt.Println("Email:", email)
		fmt.Println("Password:", password)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addSuperUserCmd)
}
