package service_test

import (
	"context"
	"github.com/gsk967/userActivity/pb"
	"github.com/gsk967/userActivity/service"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewUserService(t *testing.T) {
	t.Parallel()

	userStore := service.NewInMemoryUserStore()
	userActivityStore := service.InMemoryUserActivityStore()

	userService := service.NewUserService(userStore, userActivityStore)

	ctx := context.Background()
	createUserReq := &pb.UserInfo{UserName: "test", Email: "test@gmail.com", PhoneNo: "1231231231"}
	createUserResp, err := userService.CreateUsersReq(ctx, createUserReq)
	require.Nil(t, err)
	require.NotNil(t, createUserResp)

	// duplicate user inserting throws error
	createUserResp, err = userService.CreateUsersReq(ctx, createUserReq)
	require.Nil(t, createUserResp)
	require.NotNil(t, err)

	// adding new user activity
	addUserActivityReq := &pb.CreateUserActivityReq{UserEmail: "test@gmail.com", Activity: &pb.Activity{
		Activity:     pb.ActivityType_SLEEP,
		Day:          "2020-10-10",
		Status:       pb.Status_ACTIVE,
		TimeDuration: 7,
	}}
	addUserActivityResp, err := userService.AddUserActivityServiceReq(ctx, addUserActivityReq)
	require.Nil(t, err)
	require.NotNil(t, addUserActivityResp)

	// adding user activity who is not in our platform
	addUserActivityReq.UserEmail = "asdasdasd@gmai..com"
	addUserActivityResp, err = userService.AddUserActivityServiceReq(ctx, addUserActivityReq)
	require.Nil(t, addUserActivityResp)
	require.NotNil(t, err)

	// Get user activity list
	getUserActivityList := &pb.GetUserActivityReq{UserEmail: "test@gmail.com"}
	userActivityList, err := userService.GetUserActivityServiceReq(ctx, getUserActivityList)
	require.Nil(t, err)
	require.NotNil(t, userActivityList)

	// Update user activity status
	updateUserActivityReq := &pb.UpdateActivityStatusReq{UserEmail: "test@gmail.com", Activity: pb.ActivityType_SLEEP, Status: pb.Status_DONE, Day: "2020-10-10"}
	updateActivityStatusResp, err := userService.UpdateUserActivityServiceReq(ctx, updateUserActivityReq)
	require.Nil(t, err)
	require.NotNil(t, updateActivityStatusResp)
	require.Equal(t, updateActivityStatusResp.GetActivity().GetStatus(), pb.Status_DONE)
}
