package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"DJ-TPA-Backend/graph/generated"
	"DJ-TPA-Backend/graph/model"
	"context"
	"errors"
	"fmt"
	"strings"
)

func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateVideo(ctx context.Context, id string, input model.NewVideo) (*model.Video, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteVideo(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	var videos []*model.Video
	err := r.DB.Model(&videos).Order("id").Select()
	if err != nil {
		return nil, errors.New("Failed to query!")
	}
	return videos, nil
}

func (r *queryResolver) SearchVideos(ctx context.Context, searchQuery string) ([]*model.Video, error) {
	var videos []*model.Video
	var filteredVideos []*model.Video
	err := r.DB.Model(&videos).Order("id").Select()
	if err != nil {
		return nil, errors.New("Failed to query!")
	}
	for i := range videos {
		data, searchText := strings.ToLower(videos[i].Title), strings.ToLower(searchQuery)
		if strings.Contains(data, searchText) {
			filteredVideos = append(filteredVideos, videos[i])
		}
	}
	return filteredVideos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
