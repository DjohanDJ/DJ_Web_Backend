package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"DJ-TPA-Backend/graph/model"
	"context"
	"errors"
	"fmt"
)

func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	newComment := model.Comment{
		VideoID:         input.VideoID,
		CommentParentID: input.CommentParentID,
		CommentValue:    input.CommentValue,
		UserID:          input.UserID,
		CommentDate:     input.CommentDate,
		Likes:           input.Likes,
		Dislikes:        input.Dislikes,
	}
	_, err := r.DB.Model(&newComment).Insert()
	if err != nil {
		return nil, errors.New("Insert user failed !")
	}
	return &newComment, err
}

func (r *mutationResolver) UpdateComment(ctx context.Context, id string, input model.NewComment) (*model.Comment, error) {
	var comment model.Comment

	err := r.DB.Model(&comment).Where("id = ?", id).First()
	if err != nil {
		return nil, errors.New("Failed to query!")
	}

	comment.Likes = input.Likes
	comment.Dislikes = input.Dislikes

	_, errUpdate := r.DB.Model(&comment).Where("id = ?", id).Update()
	if errUpdate != nil {
		return nil, errors.New("Failed to query!")
	}

	return &comment, nil
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
	var comments []*model.Comment
	err := r.DB.Model(&comments).Order("id").Select()
	if err != nil {
		return nil, errors.New("Failed to query!")
	}
	return comments, nil
}
