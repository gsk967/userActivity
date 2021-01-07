package service

import (
	"errors"
	"github.com/golang/protobuf/ptypes"
	"github.com/gsk967/userActivity/pb"
	"sync"
)

var ErrActivityAlreadyExists = errors.New("activity record already exists")

type UserActivity interface {
	Save(userActivity *pb.CreateUserActivityReq) error
	GetUserActivities(userEmail string) (*pb.UserActivity, error)
	FindActivity(userEmail string, day string, activity pb.ActivityType) (*pb.Activity, error)
	UpdateActivityStatus(userEmail string, day string, activity pb.ActivityType, status pb.Status) (*pb.Activity, error)
}

type UserActivityStore struct {
	mutex sync.RWMutex
	data  map[string][]*pb.Activity
}

func InMemoryUserActivityStore() *UserActivityStore {
	return &UserActivityStore{
		data: make(map[string][]*pb.Activity),
	}
}

func (activityStore *UserActivityStore) FindActivity(
	userEmail string, day string, activityType pb.ActivityType) (*pb.Activity, error) {
	activityStore.mutex.RLock()
	defer activityStore.mutex.RUnlock()
	userDailyActivities := activityStore.data[userEmail]
	for _, activity := range userDailyActivities {
		if activity.GetDay() == day && activity.GetActivity() == activityType {
			return activity, nil
		}
	}
	return nil, errors.New("no activity record found")
}

func (activityStore *UserActivityStore) UpdateActivityStatus(userEmail string, day string, activityType pb.ActivityType, status pb.Status) (*pb.Activity, error) {
	activityStore.mutex.Lock()
	defer activityStore.mutex.Unlock()
	userDailyActivities := activityStore.data[userEmail]
	for _, activity := range userDailyActivities {
		if activity.GetDay() == day && activity.GetActivity() == activityType {
			activity.Status = status
			activity.UpdatedAt = ptypes.TimestampNow()
			return activity, nil
		}
	}
	return nil, nil
}

func (activityStore *UserActivityStore) GetUserActivities(userEmail string) (*pb.UserActivity, error) {
	activityStore.mutex.RLock()
	defer activityStore.mutex.RUnlock()
	userDailyActivities := activityStore.data[userEmail]
	return &pb.UserActivity{UserEmail: userEmail, DailyActivities: userDailyActivities}, nil
}

func (activityStore *UserActivityStore) Save(userActivity *pb.CreateUserActivityReq) error {
	activityStore.mutex.Lock()
	defer activityStore.mutex.Unlock()

	userEmailAddr := userActivity.GetUserEmail()
	userReqActivity := userActivity.GetActivity()
	timeDuration := userReqActivity.GetTimeDuration()
	if err := validateActivity(userReqActivity.GetActivity(), timeDuration); err != nil {
		return err
	}

	userDailyActivities := activityStore.data[userEmailAddr]
	for _, activity := range userDailyActivities {
		if activity.GetDay() == userReqActivity.GetDay() && activity.GetActivity() == userReqActivity.GetActivity() {
			return ErrActivityAlreadyExists
		}
	}

	userReqActivity.Activity = userReqActivity.GetActivity()
	userReqActivity.Status = userReqActivity.GetStatus()
	userReqActivity.CreatedAt = ptypes.TimestampNow()
	userReqActivity.UpdatedAt = ptypes.TimestampNow()

	activityStore.data[userEmailAddr] = append(activityStore.data[userEmailAddr], userReqActivity)
	return nil
}

func validateActivity(activityType pb.ActivityType, timeDuration uint32) error {
	switch activityType {
	case 0:
		return errors.New("this is default ")
	case 1:
		if timeDuration >= 6 && timeDuration <= 8 {
			return nil
		}
		return errors.New("please make sure you sleep 6 to 8 hours daily ")
	case 2:
		if timeDuration >= 1 && timeDuration <= 2 {
			return nil
		}
		return errors.New("please make sure you eating duration from 1 hour to 2 hours")
	case 3:
		if timeDuration >= 1 && timeDuration <= 2 {
			return nil
		}
		return errors.New("please make sure you reading duration from 1 hour to 2 hours ")
	case 4:
		if timeDuration >= 1 && timeDuration <= 2 {
			return nil
		}
		return errors.New("please make sure you played around from 1 hour to 2 hours ")
	}
	return nil
}
