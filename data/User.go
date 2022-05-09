package data

import (
	"context"
	"fmt"
)

var (
	CreateUser     = createUser
	GetUserDetails = getUserDetails
)

func createUser(ctx context.Context, userCreate *UserCreate) (user *User, err error) {

	userHash := asSha256(userCreate)
	user = &User{
		ID:        userHash,
		FirstName: userCreate.FirstName,
		LastName:  userCreate.LastName,
	}
	if err = setToRedis(ctx, user); err != nil {
		return nil, err
	}
	return user, nil

}

func getUserDetails(ctx context.Context, id string) (*User, error) {

	user := &User{
		ID: id,
	}
	if err := getFromRedis(ctx, user); err != nil {
		return nil, err
	}
	fmt.Println(user)
	return user, nil

}

func updateUser(ctx context.Context, user *User) error {
	return updateToRedis(ctx, user)
}
