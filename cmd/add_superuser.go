package cmd

import (
	"bufio"
	"context"
	"dewkit/config"
	"dewkit/internal/auth"
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"

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

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		db, err := config.GetDB(ctx)
		if err != nil {
			fmt.Println("Failed to connect to DB, error - ", err)
		}

		authService := auth.Service{DB: db}
		err = authService.CreateSuperuser(fullName, email, password)
		return err
	},
}

func init() {
	rootCmd.AddCommand(addSuperUserCmd)
}
