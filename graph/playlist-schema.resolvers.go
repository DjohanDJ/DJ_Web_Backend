package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"DJ-TPA-Backend/graph/model"
	"context"
	"errors"
	"fmt"
)

func (r *mutationResolver) CreatePlaylist(ctx context.Context, input model.NewPlaylist) (*model.Playlist, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePlaylist(ctx context.Context, playlistID string, input model.NewPlaylist) ([]*model.Playlist, error) {
	var playlists []*model.Playlist

	err := r.DB.Model(&playlists).Where("playlist_id = ?", playlistID).Select()
	if err != nil {
		return nil, errors.New("Failed to query!")
	}

	for i := range playlists {
		playlists[i].Title = input.Title
		playlists[i].Description = input.Description
		playlists[i].Privacy = input.Privacy
		playlists[i].ViewCount = input.ViewCount
		_, errUpdate := r.DB.Model(playlists[i]).Where("playlist_id = ?", playlists[i].PlaylistID).Where("video_id = ?", playlists[i].VideoID).Update()
		if errUpdate != nil {
			return nil, errors.New("Failed to query!" + errUpdate.Error())
		}
	}

	return playlists, nil
}

func (r *mutationResolver) DeletePlaylist(ctx context.Context, playlistID string, videoID int) (bool, error) {
	var playlists []*model.Playlist
	err := r.DB.Model(&playlists).Where("playlist_id = ?", playlistID).Select()
	if err != nil {
		return false, errors.New(err.Error())
	}

	for i := range playlists {
		if playlists[i].VideoID == videoID {
			_, errDel := r.DB.Model(playlists[i]).Where("playlist_id = ?", playlists[i].PlaylistID).Where("video_id = ?", playlists[i].VideoID).Delete()
			if errDel != nil {
				return false, errors.New("Failed to query!")
			}
			return true, nil
		}
	}

	return false, nil
}

func (r *queryResolver) Playlists(ctx context.Context) ([]*model.Playlist, error) {
	var playlists []*model.Playlist
	err := r.DB.Model(&playlists).Order("playlist_id").Order("video_id").Select()
	if err != nil {
		return nil, errors.New("Failed to query!")
	}
	return playlists, nil
}
