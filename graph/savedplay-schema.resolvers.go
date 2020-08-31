package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"DJ-TPA-Backend/graph/model"
	"context"
	"fmt"
	"errors"
)

func (r *mutationResolver) CreateSavedplay(ctx context.Context, input model.NewSavedplay) (*model.Savedplay, error) {
	newSavedplays := model.Savedplay{
		SavedplayID: input.SavedplayID,
		UserID:     input.UserID,
	}
	_, err := r.DB.Model(&newSavedplays).Insert()
	if err != nil {
		return nil, errors.New("Insert user failed !")
	}
	return &newSavedplays, err
}

func (r *mutationResolver) UpdateSavedplay(ctx context.Context, savedplayID string, input model.NewSavedplay) ([]*model.Savedplay, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteSavedplay(ctx context.Context, savedplayID string, userID int) (bool, error) {
	var savedplays []*model.Savedplay
	err := r.DB.Model(&savedplays).Where("savedplay_id = ?", savedplayID).Select()
	if err != nil {
		return false, errors.New(err.Error())
	}

	for i := range savedplays {
		if savedplays[i].UserID == userID {
			_, errDel := r.DB.Model(savedplays[i]).Where("savedplay_id = ?", savedplays[i].SavedplayID).Where("user_id = ?", savedplays[i].UserID).Delete()
			if errDel != nil {
				return false, errors.New("Failed to query!")
			}
			return true, nil
		}
	}

	return false, nil
}

func (r *queryResolver) Savedplays(ctx context.Context) ([]*model.Savedplay, error) {
	var savedplays []*model.Savedplay
	err := r.DB.Model(&savedplays).Select()
	if err != nil {
		return nil, errors.New("Failed to query!" + err.Error())
	}
	return savedplays, nil
}
