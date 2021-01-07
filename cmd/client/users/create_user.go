package users

import (
	"context"
	"fmt"
	"github.com/gsk967/userActivity/pb"
	"github.com/spf13/cobra"
)

func CreateUser() *cobra.Command {
	var createUserCmd = &cobra.Command{
		Use:   "create-user --username <UserName> --email <Email> --phone-no <PhoneNo>",
		Short: "create user with username, email and phone number",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmdFlags := cmd.Flags()
			serverAddr, _ := cmdFlags.GetString(ServerAddr)
			client, err := GrpcClient(serverAddr)
			if err != nil {
				return err
			}

			userName, _ := cmdFlags.GetString(UserName)
			email, _ := cmdFlags.GetString(Email)
			phoneNo, _ := cmdFlags.GetString(PhoneNo)

			createUserReq := &pb.UserInfo{
				UserName: userName,
				Email:    email,
				PhoneNo:  phoneNo,
			}
			_, err = client.CreateUsersReq(context.Background(), createUserReq)
			if err != nil {
				return err
			}
			fmt.Printf("user is created with username %s \n", userName)
			return nil
		},
	}
	createUserCmd.Flags().String(UserName, "", "--username <username>")
	createUserCmd.MarkFlagRequired(UserName)
	createUserCmd.Flags().String(Email, "", "--email <user@gmail.com>")
	createUserCmd.MarkFlagRequired(Email)
	createUserCmd.Flags().String(PhoneNo, "", "--phone-no <1231231232>")
	createUserCmd.MarkFlagRequired(PhoneNo)

	return createUserCmd
}
