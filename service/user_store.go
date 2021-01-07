package service

import (
	"errors"
	"fmt"
	"github.com/gsk967/userActivity/pb"
	"github.com/jinzhu/copier"
	"sync"
)

// ErrAlreadyExists is returned when a record with the same ID already exists in the store
var ErrAlreadyExists = errors.New("record already exists")
var NotFound = errors.New("not found user ")

type UserStore interface {
	Save(user *pb.UserInfo) error
	FindUserByUserEmail(userName string) (*pb.UserInfo, error)
}

// InMemoryUserStore stores users in memory
type InMemoryUserStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.UserInfo
}

func (store *InMemoryUserStore) FindUserByUserEmail(email string) (*pb.UserInfo, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	if store.data[email] != nil {
		return store.data[email], nil
	}
	return nil, NotFound
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		data: make(map[string]*pb.UserInfo),
	}
}

func (store *InMemoryUserStore) Save(user *pb.UserInfo) error {

	// validations.... for user
	if len(user.GetEmail()) == 0 {
		return errors.New("email should not be empty")
	}

	if len(user.GetUserName()) == 0 {
		return errors.New("username should not be empty")
	}

	if len(user.GetPhoneNo()) == 0 {
		return errors.New("phone number should not be empty")
	}
	if len(user.GetPhoneNo()) < 10 || len(user.GetPhoneNo()) > 12 {
		return errors.New("phone number length should be between 10 to 12")
	}

	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[user.Email] != nil {
		return ErrAlreadyExists
	}

	copyUserInfo, err := deepCopy(user)
	if err != nil {
		return err
	}

	store.data[copyUserInfo.Email] = copyUserInfo
	return nil
}

func deepCopy(user *pb.UserInfo) (*pb.UserInfo, error) {
	other := &pb.UserInfo{}

	err := copier.Copy(other, user)
	if err != nil {
		return nil, fmt.Errorf("cannot copy users data: %w", err)
	}

	return other, nil
}
