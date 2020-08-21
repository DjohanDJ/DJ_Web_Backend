package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"DJ-TPA-Backend/graph/model"
	"context"
	"errors"
	"fmt"
)

func (r *mutationResolver) CreateSubscriber(ctx context.Context, input model.NewSubscriber) (*model.Subscriber, error) {
	newSubscriber := model.Subscriber{
		UserID:    input.UserID,
		ChannelID: input.ChannelID,
	}
	_, err := r.DB.Model(&newSubscriber).Insert()
	if err != nil {
		return nil, errors.New("Insert user failed !")
	}
	return &newSubscriber, err
}

func (r *mutationResolver) UpdateSubscriber(ctx context.Context, id string, input model.NewSubscriber) (*model.Subscriber, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteSubscriber(ctx context.Context, id string) (bool, error) {
	var subscribers model.Subscriber
	err := r.DB.Model(&subscribers).Where("id = ?", id).First()
	if err != nil {
		return false, errors.New("Failed to query!")
	}

	_, errDel := r.DB.Model(&subscribers).Where("id = ?", id).Delete()

	if errDel != nil {
		return false, errors.New("Failed to query!")
	}

	return true, nil
}

func (r *queryResolver) Subscribers(ctx context.Context) ([]*model.Subscriber, error) {
	var subscribers []*model.Subscriber
	err := r.DB.Model(&subscribers).Order("id").Select()
	if err != nil {
		return nil, errors.New("Failed to query!")
	}
	return subscribers, nil
}
