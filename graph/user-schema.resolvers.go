package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"DJ-TPA-Backend/graph/model"
	"context"
	"errors"
	"fmt"
	"strconv"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	newUser := model.User{
		Username:     input.Username,
		Email:        input.Email,
		UserPassword: input.UserPassword,
		ChannelName:  input.ChannelName,
		UserImage:    input.UserImage,
		Restriction:  input.Restriction,
		Location:     input.Location,
		Membership:   "non-premium",
		ViewCount: 1,
		InstagramLink: "",
	}
	_, err := r.DB.Model(&newUser).Insert()
	if err != nil {
		return nil, errors.New("Insert user failed !")
	}
	return &newUser, err
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.NewUser) (*model.User, error) {
	var user model.User

	err := r.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, errors.New("Failed to query!")
	}

	user.Username = input.Username
	user.Email = input.Email
	user.UserPassword = input.UserPassword
	user.ChannelName = input.ChannelName
	user.UserImage = input.UserImage
	user.Restriction = input.Restriction
	user.Location = input.Location
	user.UserImage = input.UserImage
	user.ChannelBanner = input.ChannelBanner
	user.Membership = input.Membership
	user.ExpiredMember = input.ExpiredMember
	user.JoinDate = input.JoinDate

	_, errUpdate := r.DB.Model(&user).Where("id = ?", id).Update()
	if errUpdate != nil {
		return nil, errors.New("Failed to query!")
	}

	return &user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := r.DB.Model(&users).Order("id").Select()
	if err != nil {
		return nil, errors.New("Failed to query!")
	}
	return users, nil
}

func (r *queryResolver) SearchUserByEmail(ctx context.Context, searchEmail string) ([]*model.User, error) {
	var users []*model.User
	var filteredUser []*model.User
	err := r.DB.Model(&users).Order("id").Select()
	if err != nil {
		return nil, errors.New("Failed to query!")
	}
	for i := range users {
		if users[i].Email == searchEmail {
			filteredUser = append(filteredUser, users[i])
			break
		}
	}
	return filteredUser, nil
}

func (r *queryResolver) SearchUserByID(ctx context.Context, userID int) (*model.User, error) {
	var users []*model.User
	var filteredUser *model.User
	err := r.DB.Model(&users).Order("id").Select()
	if err != nil {
		return nil, errors.New("Failed to query!")
	}
	for i := range users {
		if users[i].ID == strconv.Itoa(userID) {
			filteredUser = users[i]
			break
		}
	}
	return filteredUser, nil
}
