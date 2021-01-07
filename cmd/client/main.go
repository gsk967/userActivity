package main

import (
	"context"
	"flag"
	"github.com/gsk967/userActivity/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

var (
	Email = "test@gmail.com"
)

func main() {
	serverAddress := flag.String("addr", "0.0.0.0:8080", "Server Addr : 0.0.0.0:8080")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)
	transportOption := grpc.WithInsecure()

	grpcClient, err := grpc.Dial(*serverAddress, transportOption)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	userServiceClient := pb.NewUsersClient(grpcClient)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// creating the users ....
	createUserReq := &pb.UserInfo{
		UserName: Email,
		Email:    "test@gmail.com",
		PhoneNo:  "1234567898",
	}
	_, err = userServiceClient.CreateUsersReq(ctx, createUserReq)
	if err != nil {
		log.Fatalln("error while creating user %w", err)
	}

	// add user activity
	createUserActivity := &pb.CreateUserActivityReq{UserEmail: Email, Activity: &pb.Activity{
		Activity:     pb.ActivityType_SLEEP,
		Status:       pb.Status_ACTIVE,
		Day:          "2020-10-10",
		TimeDuration: 7,
	}}

	resp, err := userServiceClient.AddUserActivityServiceReq(ctx, createUserActivity)
	if err != nil {
		log.Fatalln("error while adding user activity %w", err)
	}
	log.Printf("add activity resp  %v", resp)

	// get user activity
	getUserActivityReq := &pb.GetUserActivityReq{UserEmail: Email}
	activityResp, err := userServiceClient.GetUserActivityServiceReq(ctx, getUserActivityReq)
	if err != nil {
		log.Fatalln("error while adding user activity %w", err)
	}
	log.Printf("activity list %v", activityResp)

	// update user activity
	updateUserActivityReq := &pb.UpdateActivityStatusReq{
		UserEmail: Email,
		Activity:  pb.ActivityType_SLEEP,
		Status:    pb.Status_DONE,
		Day:       "2020-10-10",
	}
	updateResp, err := userServiceClient.UpdateUserActivityServiceReq(ctx, updateUserActivityReq)
	if err != nil {
		log.Fatalln("error while update user activity status %w", err)
	}
	log.Printf("updated activity info : %v", updateResp)

}
