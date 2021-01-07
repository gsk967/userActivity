package service

import (
	"context"
	"github.com/gsk967/userActivity/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type UserService struct {
	userStore         UserStore
	userActivityStore UserActivity
}

func (userService *UserService) GetUserActivityServiceReq(ctx context.Context, req *pb.GetUserActivityReq) (*pb.UserActivity, error) {
	userEmailAddr := req.GetUserEmail()
	_, err := userService.userStore.FindUserByUserEmail(userEmailAddr)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user is not found %s", userEmailAddr)
	}
	return userService.userActivityStore.GetUserActivities(userEmailAddr)
}

func NewUserService(userStore UserStore, userActivityStore UserActivity) *UserService {
	return &UserService{
		userStore:         userStore,
		userActivityStore: userActivityStore,
	}
}

func (userService *UserService) CreateUsersReq(ctx context.Context, req *pb.UserInfo) (*pb.CreateUsersResponse, error) {
	userName := req.GetEmail()
	err := userService.userStore.Save(req)

	if err != nil {
		log.Printf("User already exists with name %s...", userName)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	resp := &pb.CreateUsersResponse{
		UserInfo: req,
	}
	return resp, nil
}

func (userService *UserService) AddUserActivityServiceReq(
	ctx context.Context,
	req *pb.CreateUserActivityReq) (*pb.CreateUserActivityResponse, error) {

	userEmailAddr := req.GetUserEmail()
	userActivity := req.GetActivity()
	_, err := userService.userStore.FindUserByUserEmail(userEmailAddr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "please check %s", err)
	}
	err = userService.userActivityStore.Save(req)
	if err != nil {
		log.Printf("error occured creating the activity %s", err)
		return nil, status.Errorf(codes.InvalidArgument, "please check %w", err)
	}
	return &pb.CreateUserActivityResponse{UserEmail: userEmailAddr, Activity: userActivity}, nil
}

func (userService *UserService) UpdateUserActivityServiceReq(ctx context.Context, req *pb.UpdateActivityStatusReq) (*pb.UpdateActivityStatusResponse, error) {
	userEmailAddr := req.GetUserEmail()
	_, err := userService.userStore.FindUserByUserEmail(userEmailAddr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "please check %s", err)
	}

	activity, err := userService.userActivityStore.FindActivity(userEmailAddr, req.GetDay(), req.GetActivity())
	if activity == nil {
		return nil, status.Errorf(codes.NotFound, "activity record not found")
	}

	updatedActivity, err := userService.userActivityStore.UpdateActivityStatus(userEmailAddr, req.GetDay(), req.GetActivity(), req.GetStatus())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "please check %s", err)
	}
	return &pb.UpdateActivityStatusResponse{UserEmail: userEmailAddr, Activity: updatedActivity}, nil
}
