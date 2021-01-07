package main

import (
	"fmt"
	"github.com/gsk967/userActivity/cmd/client/users"
	"github.com/spf13/cobra"
	"os"
)

const (
	ServerAddr = "server-addr"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "client --addr <ServerAddr>",
		Short: "client can query grpc server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	rootCmd.PersistentFlags().String(ServerAddr, "0.0.0.0:8080", "--server-addr <0.0.0.0:8080>")
	rootCmd.MarkPersistentFlagRequired(ServerAddr)

	// adding sub commands
	rootCmd.AddCommand(users.CreateUser())
	rootCmd.AddCommand(users.AddUserActivity())
	rootCmd.AddCommand(users.GetUserActivities())
	rootCmd.AddCommand(users.UpdateUserActivity())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}
}
