package users

import (
	"context"
	"fmt"
	"github.com/gsk967/userActivity/pb"
	"github.com/spf13/cobra"
	"log"
)

func enumActivity(activity string) pb.ActivityType {
	switch activity {
	case "eat":
		return pb.ActivityType_EAT
	case "sleep":
		return pb.ActivityType_SLEEP
	case "play":
		return pb.ActivityType_PLAY
	case "read":
		return pb.ActivityType_READ
	default:
		return pb.ActivityType_UNKNOWN
	}
}

func enumActivityStatus(activityStatus string) pb.Status {
	switch activityStatus {
	case "active":
		return pb.Status_ACTIVE
	case "done":
		return pb.Status_DONE
	default:
		return pb.Status_UN_KNOWN
	}
}

func AddUserActivity() *cobra.Command {
	var AddUserActivityCmd = &cobra.Command{
		Use:   "add-activity --username <UserName> --email <Email> --phone-no <PhoneNo>",
		Short: "add-activity with username, email and phone number",
		RunE: func(cmd *cobra.Command, args []string) error {
			// add user activity
			cmdFlags := cmd.Flags()
			serverAddr, _ := cmdFlags.GetString(ServerAddr)
			client, err := GrpcClient(serverAddr)
			if err != nil {
				return err
			}

			email, _ := cmdFlags.GetString(Email)
			activityDay, _ := cmdFlags.GetString(ActivityDay)
			activityType, _ := cmdFlags.GetString(ActivityType)
			activityStatus, _ := cmdFlags.GetString(ActivityStatus)
			activityTimeDuration, _ := cmdFlags.GetUint32(ActivityTimeDuration)

			createUserActivityReq := &pb.CreateUserActivityReq{
				UserEmail: email,
				Activity: &pb.Activity{
					Activity:     enumActivity(activityType),
					Status:       enumActivityStatus(activityStatus),
					Day:          activityDay,
					TimeDuration: activityTimeDuration,
				},
			}
			_, err = client.AddUserActivityServiceReq(context.Background(), createUserActivityReq)
			if err != nil {
				return err
			}
			fmt.Printf("user activity is created %s \n", email)
			return nil
		},
	}
	AddUserActivityCmd.Flags().String(Email, "", "--email <user@gmail.com>")
	AddUserActivityCmd.MarkFlagRequired(Email)
	AddUserActivityCmd.Flags().String(ActivityDay, "", "--activity-day <2020-10-10>")
	AddUserActivityCmd.MarkFlagRequired(ActivityDay)
	AddUserActivityCmd.Flags().String(ActivityStatus, "", "--activity-status <active/done>")
	AddUserActivityCmd.MarkFlagRequired(ActivityStatus)
	AddUserActivityCmd.Flags().Uint32(ActivityTimeDuration, 0, "--activity-time-duration <6>")
	AddUserActivityCmd.MarkFlagRequired(ActivityTimeDuration)
	AddUserActivityCmd.Flags().String(ActivityType, "", "--activity-type <eat/sleep/play/>")
	AddUserActivityCmd.MarkFlagRequired(ActivityType)

	return AddUserActivityCmd
}

func GetUserActivities() *cobra.Command {
	var GetUserActivitiesCmd = &cobra.Command{
		Use:   "get-activities --email <email>>",
		Short: "Get activities list for user",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmdFlags := cmd.Flags()
			// grpc client connection
			serverAddr, _ := cmdFlags.GetString(ServerAddr)
			client, err := GrpcClient(serverAddr)
			if err != nil {
				return err
			}

			email, _ := cmdFlags.GetString(Email)

			getUserActivityReq := &pb.GetUserActivityReq{
				UserEmail: email,
			}
			userActivities, err := client.GetUserActivityServiceReq(context.Background(), getUserActivityReq)
			if err != nil {
				return err
			}
			fmt.Printf("user %s activities %v\n", email, userActivities)
			return nil
		},
	}
	GetUserActivitiesCmd.Flags().String(Email, "", "--email <user@gmail.com>")
	GetUserActivitiesCmd.MarkFlagRequired(Email)

	return GetUserActivitiesCmd
}

func UpdateUserActivity() *cobra.Command {
	var UpdateUserActivityCmd = &cobra.Command{
		Use:   "update-activity --email <email>>",
		Short: "update user activity status  for user",
		RunE: func(cmd *cobra.Command, args []string) error {
			// update user activity
			cmdFlags := cmd.Flags()

			// grpc client connection
			serverAddr, _ := cmdFlags.GetString(ServerAddr)
			client, err := GrpcClient(serverAddr)
			if err != nil {
				return err
			}

			email, _ := cmdFlags.GetString(Email)
			activityDay, _ := cmdFlags.GetString(ActivityDay)
			activityType, _ := cmdFlags.GetString(ActivityType)
			activityStatus, _ := cmdFlags.GetString(ActivityStatus)

			updateUserActivityStatusReq := &pb.UpdateActivityStatusReq{
				UserEmail: email,
				Day:       activityDay,
				Status:    enumActivityStatus(activityStatus),
				Activity:  enumActivity(activityType),
			}
			updateResp, err := client.UpdateUserActivityServiceReq(context.Background(), updateUserActivityStatusReq)
			log.Printf("updated activity info : %v", updateResp)
			if err != nil {
				return err
			}
			return nil
		},
	}

	UpdateUserActivityCmd.Flags().String(Email, "", "--email <user@gmail.com>")
	UpdateUserActivityCmd.MarkFlagRequired(Email)
	UpdateUserActivityCmd.Flags().String(ActivityDay, "", "--activity-day <2020-10-10>")
	UpdateUserActivityCmd.MarkFlagRequired(ActivityDay)
	UpdateUserActivityCmd.Flags().String(ActivityStatus, "", "--activity-status <active/done>")
	UpdateUserActivityCmd.MarkFlagRequired(ActivityStatus)
	UpdateUserActivityCmd.Flags().String(ActivityType, "", "--activity-type <eat/sleep/play/>")
	UpdateUserActivityCmd.MarkFlagRequired(ActivityType)

	return UpdateUserActivityCmd
}
