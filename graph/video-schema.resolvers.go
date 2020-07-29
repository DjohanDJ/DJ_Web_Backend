package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
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

func (r *queryResolver) SearchKidVideos(ctx context.Context) ([]*model.Video, error) {
	var videos []*model.Video
	var filteredVideos []*model.Video
	err := r.DB.Model(&videos).Order("id").Select()
	if err != nil {
		return nil, errors.New("Failed to query!")
	}
	for i := range videos {
		if videos[i].Restriction == "kid" {
			filteredVideos = append(filteredVideos, videos[i])
		}
	}
	return filteredVideos, nil
}

func (r *queryResolver) SearchHomeVideosManager(ctx context.Context, isKid bool) ([]*model.Video, error) {
	if isKid {
		return r.SearchKidVideos(ctx)
	} else {
		return r.Videos(ctx)
	}
}
