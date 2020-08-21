package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"DJ-TPA-Backend/graph/model"
	"context"
	"errors"
	"fmt"
)

func (r *mutationResolver) CreateMembership(ctx context.Context, input model.NewMembership) (*model.Membership, error) {
	newMembership := model.Membership{
		UserID:          input.UserID,
		JoinPremiumDate: input.JoinPremiumDate,
		MemberType: input.MemberType,
	}
	_, err := r.DB.Model(&newMembership).Insert()
	if err != nil {
		return nil, errors.New("Insert user failed !")
	}
	return &newMembership, err
}

func (r *mutationResolver) UpdateMembership(ctx context.Context, id string, input model.NewMembership) (*model.Membership, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMembership(ctx context.Context, id string) (bool, error) {
	var memberships model.Membership
	err := r.DB.Model(&memberships).Where("id = ?", id).First()
	if err != nil {
		return false, errors.New("Failed to query!")
	}

	_, errDel := r.DB.Model(&memberships).Where("id = ?", id).Delete()

	if errDel != nil {
		return false, errors.New("Failed to query!")
	}

	return true, nil
}

func (r *queryResolver) Memberships(ctx context.Context) ([]*model.Membership, error) {
	var memberships []*model.Membership
	err := r.DB.Model(&memberships).Order("id").Select()
	if err != nil {
		return nil, errors.New("Failed to query!")
	}
	return memberships, nil
}
